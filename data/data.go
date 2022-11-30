package data

import (
	"fmt"
	"math/big"
	"math/rand"
	"os"
)

type Block struct {
	Content []byte   // 4096 bytes
	Tag     *big.Int // 128 bytes
}

func GenDataAs1Kb(size int, name string) {
	wholeData := make([]byte, 0)
	for size > 0 {
		uintBlock := make([]byte, 1024, 1024)
		rand.Read(uintBlock)
		wholeData = append(wholeData, uintBlock...)
		size--
	}
	err := os.WriteFile(fmt.Sprintf("data/%s", name), wholeData, 0666)
	check(err)
}

func DivideInto4kb(ba []byte) []Block {
	nums := len(ba) / 4096
	last := len(ba) % 4096
	if last != 0 {
		nums += 1
	}

	var temp []byte
	fileSet := make([]Block, 0)

	for i := 0; i < nums; i++ {
		if i == nums-1 {
			temp = ba[i*4096:]
		} else {
			temp = ba[i*4096 : (i+1)*4096]
		}
		block := Block{
			Content: temp,
		}

		fileSet = append(fileSet, block)
	}
	return fileSet
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
