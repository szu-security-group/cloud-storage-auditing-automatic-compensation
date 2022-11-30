package main

import (
	"autoPayAudit/core"
	"autoPayAudit/data"
	"fmt"
	"os"
	"time"
)

func main() {
	// 数据文件大小
	filename := "16MB"

	// 生成数据文件
	//data.GenDataAs1Kb(1024*1024, filename)

	d, err := os.ReadFile("data/" + filename)
	check(err)

	var timeList []int64

	// 将文件分块
	fileSet := data.DivideInto4kb(d)

	// 设置密钥长度
	sk, pk := core.KeySetup(1024)

	// 计算HVT标签生成时间
	start := time.Now().UnixMicro()
	core.TagGenerate(&fileSet, sk, pk, 16)
	end := time.Now().UnixMicro()

	interval := end - start
	timeList = append(timeList, interval)

	Type := 0 // 0 代表单标签算法流程，1 代表聚合标签算法流程
	if Type == 1 {
		s, gs := core.GsSet(pk.N, sk.G)
		fmt.Println("GsSet SUCCESS")

		chalSuite := core.ChalSet(&fileSet, 460, s, pk.N)
		fmt.Println("ChallengeTagSet SUCCESS")

		proof := core.ProofGen(&fileSet, chalSuite.Index, gs, pk.N)
		fmt.Println("ProofGen Success SUCCESS")

		result := core.Verify(chalSuite, proof, pk.E, pk.N)
		fmt.Println("Verify Success SUCCESS")
		fmt.Println(result)

	} else {
		// 计算聚合标签
		start = time.Now().UnixMicro()
		aggTs := core.AggregateTag(&fileSet, pk, 20, 20)
		end = time.Now().UnixMicro()
		interval = end - start
		timeList = append(timeList, interval)
		fmt.Println("Aggregate Tags SUCCESS")

		start = time.Now().UnixMicro()
		s, gs := core.GsSet(pk.N, sk.G)
		chalSuitea := core.ChalASet(aggTs, 460, s, pk.N)
		end = time.Now().UnixMicro()
		interval = end - start
		timeList = append(timeList, interval)
		fmt.Println("Aggregate ChallengeTags SUCCESS")

		// 服务端生成Proof
		start = time.Now().UnixMicro()
		proofa := core.ProofGenA(fileSet, chalSuitea.Index, gs, pk.N)
		end = time.Now().UnixMicro()
		interval = end - start
		timeList = append(timeList, interval)
		fmt.Println("Aggregate ProofGen SUCCESS")

		// 对结果进行验证，这一部分会用区块链实现
		resulta := core.VerifyA(chalSuitea, proofa, pk.E, pk.N)
		fmt.Println("Aggregate Verify Success SUCCESS")
		fmt.Println(resulta)
	}

	fmt.Println(timeList)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
