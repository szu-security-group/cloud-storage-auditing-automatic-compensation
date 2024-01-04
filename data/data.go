/*

The smart contract was deployed on the latest version of Ethereum.
We used solidity (version 0.60) to write the smart contract.
The protocol is implemented using Golang 1.17 (blockchain offline) and solidity (blockchain online).
The modulus $N$ is 1024 bits and the size of data block is 4KB. The experimental data ranges from 100MB to 1GB.


*/

package data

import (
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Block struct {
	Content []byte   // 4096 bytes
	Tag     *big.Int // 128 bytes
}

func GetDataAsSize(size string) []byte {
	size = strings.ToLower(size)
	ans := 1
	if strings.Contains(size, "mb") {
		ans = 1024
	} else if strings.Contains(size, "gb") {
		ans = 1024 * 1024
	}

	num, _ := strconv.Atoi(size[:len(size)-2])

	_, err := os.Stat(size)
	if err != nil && os.IsNotExist(err) {
		GenDataAs1Kb(num*ans, size)
	}
	d, err := os.ReadFile("data/" + size)
	return d
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
