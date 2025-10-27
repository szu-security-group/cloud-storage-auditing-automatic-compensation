/*

The smart contract was deployed on the latest version of Ethereum.
We used solidity (version 0.60) to write the smart contract.
The protocol is implemented using Golang 1.17 (blockchain offline) and solidity (blockchain online).
The modulus $N$ is 1024 bits and the size of data block is 4KB. The experimental data ranges from 100MB to 1GB.


*/

package core

import (
	"fmt"
	"time"
)

type TimeController struct {
	timeQueue []string
	phase     string
	startTime int64
	endTime   int64
}

func (t *TimeController) Start(phase string) {

	if t.timeQueue == nil {
		t.timeQueue = []string{}
	}
	t.phase = phase
	t.startTime = time.Now().UnixMicro()
}

func (t *TimeController) End() {
	t.endTime = time.Now().UnixMicro()
	t.timeQueue = append(t.timeQueue, fmt.Sprintf("%v:%v µs", t.phase, t.endTime-t.startTime))
}

func (t *TimeController) Output() {
	s := ""
	for i := range t.timeQueue {
		s += fmt.Sprintf("%v|", t.timeQueue[i])
	}
	fmt.Println(s)
}
