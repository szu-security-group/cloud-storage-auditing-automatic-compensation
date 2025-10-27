package core

import (
	"autoPayAudit/data"

	"golang.org/x/crypto/sha3"
)

type MHT struct {
	LeafHashes [][]byte
	TreeNodes  [][][]byte
	RootHash   []byte
	LeafCount  int
	BlockSize  int
}

func Keccak256(data ...[]byte) []byte {
	h := sha3.NewLegacyKeccak256()
	for _, d := range data {
		h.Write(d)
	}
	return h.Sum(nil)
}

func BuildMHT(blocks []data.Block) *MHT {
	leafCount := len(blocks)
	blockSize := 4096
	n := 1
	for n < leafCount {
		n *= 2
	}

	leaves := make([][]byte, n)
	for i := 0; i < n; i++ {
		var leafData []byte
		if i < leafCount {
			leafData = blocks[i].Content
		} else {
			leafData = make([]byte, blockSize)
		}

		prefix := []byte{0x00}
		idx := intToBytes(i)
		hash := Keccak256(prefix, idx, leafData)
		leaves[i] = hash
	}
	tree := [][]byte{}
	tree = append(tree, leaves...)
	levels := [][][]byte{leaves}
	cur := leaves
	for len(cur) > 1 {
		next := make([][]byte, 0, len(cur)/2)
		for i := 0; i < len(cur); i += 2 {
			left := cur[i]
			var right []byte
			if i+1 < len(cur) {
				right = cur[i+1]
			} else {
				right = left
			}
			hash := Keccak256([]byte{0x01}, left, right)
			next = append(next, hash)
		}
		levels = append(levels, next)
		cur = next
	}
	root := levels[len(levels)-1][0]
	return &MHT{
		LeafHashes: leaves,
		TreeNodes:  levels,
		RootHash:   root,
		LeafCount:  n,
		BlockSize:  blockSize,
	}
}

func (m *MHT) Root() []byte {
	return m.RootHash
}

func (m *MHT) GenProof(idx int) [][]byte {
	proof := [][]byte{}
	pos := idx

	for level := 0; level < len(m.TreeNodes)-1; level++ {
		layer := m.TreeNodes[level]
		var sibling int
		if pos%2 == 0 {
			sibling = pos + 1
		} else {
			sibling = pos - 1
		}
		if sibling >= len(layer) {
			proof = append(proof, layer[pos])
		} else {
			proof = append(proof, layer[sibling])
		}
		pos /= 2
	}
	return proof
}

func VerifyProof(leafData []byte, idx int, proof [][]byte, root []byte, totalLeaves int) bool {
	hash := Keccak256([]byte{0x00}, intToBytes(idx), leafData)
	pos := idx
	for _, sib := range proof {
		if pos%2 == 0 {
			hash = Keccak256([]byte{0x01}, hash, sib)
		} else {
			hash = Keccak256([]byte{0x01}, sib, hash)
		}
		pos /= 2
	}
	return equal(hash, root)
}

func intToBytes(i int) []byte {
	b := make([]byte, 4)
	b[0] = byte(i >> 24)
	b[1] = byte(i >> 16)
	b[2] = byte(i >> 8)
	b[3] = byte(i)
	return b
}

func equal(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
