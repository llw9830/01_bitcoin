package main

import "fmt"

func main()  {
	total	:=	0.0			//	⽐特币总数
	rewardCount	:=	50.0	//	奖励	BTC	的数量
	blockInterval	:=	21	//	区块间隔，单位万
	reduceCount	:=	0		//减半次数

	for rewardCount > 0 {
		total += rewardCount * float64(blockInterval)
		reduceCount++
		// 每挖到21w	个矿，奖励减半
		rewardCount *= 0.5
	}
	fmt.Printf("⽐特币总数：%f	万,衰减次数：%d\n",	total,	reduceCount)
}
