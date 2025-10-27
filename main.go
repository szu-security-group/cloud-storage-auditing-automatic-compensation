/*

The smart contract was deployed on the latest version of Ethereum.
We used solidity (version 0.60) to write the smart contract.
The protocol is implemented using Golang 1.17 (blockchain offline) and solidity (blockchain online).
The modulus $N$ is 1024 bits and the size of data block is 4KB. The experimental data ranges from 100MB to 1GB.


*/

package main

import (
	"autoPayAudit/core"
	"autoPayAudit/data"
	"fmt"
)

var SecurityLevel = 1024
var DataSize = "16MB"
var GoNumber = 16
var ChallengeNumber = 10
var windowSize = 20

func main() {
	timeCollect := core.TimeController{}

	// GetDataAsSize generates dataset
	f := data.GetDataAsSize(DataSize)

	// divide data block
	fileSet := data.DivideInto4kb(f)

	timeCollect.Start("BuildMHT")
	core.BuildMHT(fileSet)
	timeCollect.End()

	sk, pk := core.KeySetup(SecurityLevel)

	timeCollect.Start("TagGenerate")
	core.TagGenerate(&fileSet, sk, pk, GoNumber)
	timeCollect.End()

	timeCollect.Start("AggregateTag")
	aggTs := core.AggregateTag(&fileSet, pk, windowSize, windowSize)
	timeCollect.End()

	timeCollect.Start("Challenge")
	s, gs := core.GetGs(pk.N, sk.G)

	chalSuitea := core.SetChallengeTagWithAggregation(aggTs, ChallengeNumber, s, pk.N)
	timeCollect.End()
	fmt.Println("Aggregate ChallengeTags SUCCESS")

	timeCollect.Start("GenerateProof")
	proofa := core.ProofGenWithAggregation(fileSet, chalSuitea.Index, gs, pk.N)
	timeCollect.End()
	fmt.Println("Aggregate ProofGen SUCCESS")

	// implement it on blockchain
	_ = core.VerifyWithAggregation(chalSuitea, proofa, pk.E, pk.N)
	fmt.Printf("Aggregate Verify SUCCESS\n")

	timeCollect.Output()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
