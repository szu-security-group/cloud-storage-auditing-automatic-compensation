package core

import (
	"autoPayAudit/data"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"math/big"
	"sync"
)

// KeySuite Define the structure of the public key and secret key required by the protocol
type KeySuite struct {
	D *big.Int // D is a secret key
	G int64
	N *big.Int // N and E are public keys
	E int
}

// follow above.
type sk struct {
	D *big.Int
	G *big.Int
}

type Pk struct {
	E *big.Int
	N *big.Int
}

// ATag is aggregations base Tag (BAT)
type ATag struct {
	Index []int    // block indexes included in a window
	Tag   *big.Int // value
}

// Ch is Challenge information
type Ch struct {
	Index         []int
	ChallengeTags []*big.Int
}

// ACh is aggregations challenge Tag
type ACh struct {
	Index  [][]int
	ChalTs []*big.Int
}

var wg sync.WaitGroup

// KeySetup generate keys and public keys according to the securityLevel
func KeySetup(securityLevel int) (Sk sk, pk Pk) {
	keyPair, err := rsa.GenerateKey(rand.Reader, securityLevel)

	g, err := rand.Int(rand.Reader, keyPair.N)
	check(err)

	Sk = sk{
		D: keyPair.D,
		G: g,
	}
	pk = Pk{
		N: keyPair.N,
		E: big.NewInt(int64(keyPair.E)),
	}
	return
}

// TagGenerate generates HVT for each data block called by client
// Open goCount goroutines for concurrent computing
func TagGenerate(fileSet *[]data.Block, sk2 sk, pk2 Pk, goCount int) {
	d := sk2.D
	g := sk2.G
	N := pk2.N

	length := len(*fileSet)

	wg = sync.WaitGroup{}
	for i := 0; i < goCount; i++ {
		s := i * (length / goCount)
		end := (i + 1) * (length / goCount)
		wg.Add(1)
		go func(start int, end int) {
			defer wg.Done()
			for j := start; j < end; j++ {
				// convert bytes to bigNum
				m := new(big.Int).SetBytes((*fileSet)[j].Content)
				temp := new(big.Int).Exp(g, m, N)
				// g^m^d mod N
				temp.Exp(temp, d, N)
				// c <- temp
				(*fileSet)[j].Tag = temp
			}
		}(s, end)
	}
	wg.Wait()
}

// AggregateTag aggregates HVTs into windows
func AggregateTag(fileSet *[]data.Block, pk2 Pk, w, c int) *[]ATag {
	l := len(*fileSet)
	ts := make([]ATag, 0)

	for i := 0; i < l; i++ {
		if i%c == 0 && i != l-1 {
			t := new(big.Int).Add((*fileSet)[i].Tag, big.NewInt(0)) // 这里是浅拷贝而不是深拷贝，主要复制值
			l1 := make([]int, 0)                                    // 存储聚合标签的索引
			l1 = append(l1, i)

			if i+w > l { // 最后聚集标签不满的情况，取剩余标签聚合
				for b := i + 1; b <= l-1; b++ {
					t.Mul(t, (*fileSet)[b].Tag).Mod(t, pk2.N)
					l1 = append(l1, b)
				}
			} else {
				for j := 1; j < w; j++ {
					t.Mul(t, (*fileSet)[i+j].Tag).Mod(t, pk2.N)
					l1 = append(l1, i+j)
				}
			}
			ts = append(ts, ATag{
				Tag:   t,
				Index: l1,
			})
		}
	}
	return &ts
}

// GetGs generates a random number s for each challenge
func GetGs(N *big.Int, g *big.Int) (s *big.Int, gs *big.Int) {
	s, _ = rand.Int(rand.Reader, big.NewInt(1000000))

	// g^s mod N
	gs = new(big.Int).Exp(g, s, N)
	return

}

// SetChallengeWithHVT use HVT as base tag, not BAT
func SetChallengeWithHVT(fileSet *[]data.Block, rSize int, s *big.Int, N *big.Int) Ch {
	length := len(*fileSet)
	if rSize > length {
		rSize = length
	}
	randIndex := make([]int, rSize)
	for i := 0; i < rSize; i++ {
		ri, _ := rand.Int(rand.Reader, big.NewInt(int64(length)))
		randIndex[i] = int(ri.Int64())
	}

	ChallengeSet := make([]*big.Int, 0)
	for _, index := range randIndex {
		// challengeTag = T^s mod N
		chalT := new(big.Int).Exp((*fileSet)[index].Tag, s, N)
		ChallengeSet = append(ChallengeSet, chalT)
	}

	Chs := Ch{
		Index:         randIndex,
		ChallengeTags: ChallengeSet,
	}
	return Chs

}

// SetChallengeTagWithAggregation use BAT as base tag, boosting effective of protocol
func SetChallengeTagWithAggregation(tags *[]ATag, rSize int, s *big.Int, N *big.Int) ACh {
	length := len(*tags)
	w := len((*tags)[0].Index)
	allBlock := length * w

	if rSize > allBlock {
		rSize = allBlock
	}
	rSize = rSize/w + 1

	randIndex := make([][]int, rSize)
	chalS := make([]*big.Int, 0)

	for i := 0; i < rSize; i++ {
		ri, _ := rand.Int(rand.Reader, big.NewInt(int64(length)))
		j := int(ri.Int64())

		randIndex[i] = (*tags)[j].Index
		chalT := new(big.Int).Exp((*tags)[j].Tag, s, N)
		chalS = append(chalS, chalT)

	}

	c := ACh{
		Index:  randIndex,
		ChalTs: chalS,
	}
	return c
}

// ProofGen return a proof for challenge (HVT)
func ProofGen(fileSet *[]data.Block, ranIndex []int, gs *big.Int, N *big.Int) []*big.Int {
	proof := make([]*big.Int, 0)
	for _, index := range ranIndex {
		m := new(big.Int).SetBytes((*fileSet)[index].Content)

		// proof = gs^m mod N
		p := new(big.Int).Exp(gs, m, N)
		proof = append(proof, p)
	}
	return proof
}

// ProofGenWithAggregation return a proof for challenge (BAT)
func ProofGenWithAggregation(fileSet []data.Block, ranIndex [][]int, gs *big.Int, N *big.Int) []*big.Int {
	proof := make([]*big.Int, 0)

	for _, indexs := range ranIndex {
		t := big.NewInt(1)
		for _, i := range indexs {
			m := new(big.Int).SetBytes(fileSet[i].Content)
			p := new(big.Int).Exp(gs, m, N)
			t.Mul(t, p).Mod(t, N)
		}
		proof = append(proof, t)
	}
	return proof
}

func Verify(challSet Ch, proof []*big.Int, E *big.Int, N *big.Int) (wrongFlag []int) {
	var Proof *big.Int
	var right *big.Int

	length := len(challSet.Index)
	wrongFlag = make([]int, 0)

	for i := 0; i < length; i++ {
		Proof = proof[i]
		right = new(big.Int).Exp(challSet.ChallengeTags[i], E, N)
		if Proof.Cmp(right) == 0 {
			fmt.Printf("Number %d Block is intact \n", challSet.Index[i])
		} else {
			fmt.Printf("Number %d Block is wrong  \n", challSet.Index[i])
			wrongFlag = append(wrongFlag, i)
		}
	}
	return

}

func VerifyWithAggregation(chs ACh, proof []*big.Int, E *big.Int, N *big.Int) (wFlag [][]int) {
	var Proof *big.Int
	var right *big.Int

	wFlag = make([][]int, 0)
	for i := 0; i < len(chs.Index); i++ {
		Proof = proof[i]
		right = new(big.Int).Exp(chs.ChalTs[i], E, N)
		if Proof.Cmp(right) == 0 {
			continue
			fmt.Printf("Number %d Block is intact \n", chs.Index[i])
		} else {
			fmt.Printf("Number %d Block is wrong  \n", chs.Index[i])
			wFlag = append(wFlag, chs.Index[i])
		}
	}
	return
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
