package main

import (
	"fmt"
	"os"
)

type CLI struct {
	bc *BlockChain
}

const Usage  = `
	addBlock --data DATA		"add data to BlockChain."
	printChain					"print all BlockChain data."
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
	default:
		fmt.Println("无效的命令, 请检查！")
		fmt.Printf(Usage)
	}

}