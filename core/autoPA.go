/*

The smart contract was deployed on the latest version of Ethereum.
We used solidity (version 0.60) to write the smart contract.
The protocol is implemented using Golang 1.17 (blockchain offline) and solidity (blockchain online).
The modulus $N$ is 1024 bits and the size of data block is 4KB. The experimental data ranges from 100MB to 1GB.

*/

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
	ChalTs *big.Int
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
			t := new(big.Int).Add((*fileSet)[i].Tag, big.NewInt(0))
			l1 := make([]int, 0)
			l1 = append(l1, i)

			if i+w > l {
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

// SetChallengeTagWithAggregation use BAT as base tag, boosting effective of protocol
func SetChallengeTagWithAggregation(tags *[]ATag, rSize int, s *big.Int, N *big.Int) ACh {
	length := len(*tags)
	if rSize > length {
		rSize = length
	}
	randIndex := make([][]int, rSize)
	finalChalT := big.NewInt(1)

	for i := 0; i < rSize; i++ {
		ri, _ := rand.Int(rand.Reader, big.NewInt(int64(length)))
		j := int(ri.Int64())

		randIndex[i] = (*tags)[j].Index
		chalT := new(big.Int).Exp((*tags)[j].Tag, s, N)

		finalChalT.Mul(finalChalT, chalT).Mod(finalChalT, N)
	}

	c := ACh{
		Index:  randIndex,
		ChalTs: finalChalT,
	}
	return c
}

// ProofGenWithAggregation return a proof for challenge (BAT)
func ProofGenWithAggregation(fileSet []data.Block, ranIndex [][]int, gs *big.Int, N *big.Int) *big.Int {
	finalProof := big.NewInt(1)

	for _, indexs := range ranIndex {
		windowProof := big.NewInt(1)
		for _, i := range indexs {
			m := new(big.Int).SetBytes(fileSet[i].Content)
			p := new(big.Int).Exp(gs, m, N)
			windowProof.Mul(windowProof, p).Mod(windowProof, N)
		}

		finalProof.Mul(finalProof, windowProof).Mod(finalProof, N)
	}

	return finalProof
}

func VerifyWithAggregation(chs ACh, proof *big.Int, E *big.Int, N *big.Int) bool {
	right := new(big.Int).Exp(chs.ChalTs, E, N)

	if proof.Cmp(right) == 0 {
		fmt.Println("All challenge windows are verified successfully")
		return true
	} else {
		fmt.Println("Verification failed, data is tampered")
		return false
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
