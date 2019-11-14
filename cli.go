package main

import (
	"fmt"
	"os"
	"strconv"
)

type CLI struct {
	bc *BlockChain
}

const Usage  = `
	addBlock --data DATA			"add data to BlockChain."
	printChain						"print all BlockChain data."
	getBalance --address ADDRESS	"获取指定地址余额"
	send FROM TO AMOUNT MINER DATA	"由FROM转AMOUNT给TO，由MINER挖矿，同时写入DATA"
`

func (cli *CLI) Run ()  {
	arge := os.Args
	if len(arge) < 2 { // 参数不够，打印帮助
		fmt.Printf(Usage)
		return
	}

	cmd := arge[1]
	switch cmd {
	case "addBlock": // 添加数据
		fmt.Println("添加区块")

		if len(arge) == 4 &&  arge[2] == "--data" {
			// a.获取数据 b.使用bc添加区块Addblock
			data := arge[3]

			cli.AddBlock(data)
		} else  {
			fmt.Printf("添加区块参数使用不当，请检查.\n")
			fmt.Printf(Usage)
		}
	case "printChain": // 打印数据
		fmt.Println("打印区块")
		cli.PrintBlockChain()
	case "getBalance": // 打印数据
		if len(arge) == 4 &&  arge[2] == "--address" {
			// a.获取数据 b.使用bc添加区块Addblock
			address := arge[3]
			cli.GetBalance(address)
		}
	case "send":
		fmt.Println("转账开始...")
		if len(arge) != 7 {
			fmt.Printf("参数个数错误，请检查！\n")
			fmt.Println(Usage)
			return
		}
		//.\block.exe send FROM TO AMOUNT MINER DATA "由FROM转AMOUNT给TO，由MINER挖矿，同时写入DATA"
		from := arge[2]
		to := arge[3]
		amount, _ := strconv.ParseFloat(arge[4], 64)
		miner := arge[5]
		data := arge[6]
		cli.send(from, to, amount, miner, data)
	default:
		fmt.Println("无效的命令, 请检查！")
		fmt.Printf(Usage)
	}

}