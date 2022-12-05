package main

import (
	"autoPayAudit/core"
	"autoPayAudit/data"
	"fmt"
)

var SecurityLevel = 1024
var DataSize = "16MB"
var GoNumber = 16
var ChallengeNumber = 460
var windowSize = 20

func main() {
	timeCollect := core.TimeController{}

	// GetDataAsSize generates dataset
	f := data.GetDataAsSize(DataSize)

	// divide data block
	fileSet := data.DivideInto4kb(f)

	sk, pk := core.KeySetup(SecurityLevel)

	timeCollect.Start("TagGenerate")
	core.TagGenerate(&fileSet, sk, pk, GoNumber)
	timeCollect.End()

	Type := 0 // 0 represent HVT based protocol，and 1 represent BAT based protocol
	if Type == 1 {
		s, gs := core.GetGs(pk.N, sk.G)
		fmt.Println("GsSet SUCCESS")

		chalSuite := core.SetChallengeWithHVT(&fileSet, ChallengeNumber, s, pk.N)
		fmt.Println("ChallengeTagSet SUCCESS")

		proof := core.ProofGen(&fileSet, chalSuite.Index, gs, pk.N)
		fmt.Println("ProofGen Success SUCCESS")

		result := core.Verify(chalSuite, proof, pk.E, pk.N)
		fmt.Println("Verify Success SUCCESS")
		fmt.Println(result)

	} else {
		timeCollect.Start("AggregateTag")
		aggTs := core.AggregateTag(&fileSet, pk, windowSize, windowSize)
		timeCollect.End()

		timeCollect.Start("Challenge")
		s, gs := core.GetGs(pk.N, sk.G)
		chalSuitea := core.SetChallengeTagWithAggregation(aggTs, 460, s, pk.N)
		timeCollect.End()
		fmt.Println("Aggregate ChallengeTags SUCCESS")

		timeCollect.Start("GenerateProof")
		proofa := core.ProofGenWithAggregation(fileSet, chalSuitea.Index, gs, pk.N)
		timeCollect.End()
		fmt.Println("Aggregate ProofGen SUCCESS")

		// implement it on blockchain
		_ = core.VerifyWithAggregation(chalSuitea, proofa, pk.E, pk.N)
		fmt.Printf("Aggregate Verify SUCCESS")
	}
	timeCollect.Output()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
