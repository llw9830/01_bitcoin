package main

import "fmt"

// 添加区块
func (cli *CLI) AddBlock (data string)  {
	cli.bc.AddBlock(data)
} 


// 打印区块
func (cli *CLI) PrintBlockChain()  {
	// 调用迭代器
	it := cli.bc.NewIterator()
	for  {
		// 翻回去看左移
		block := it.Next()

		fmt.Printf("===========================\n\n")
		fmt.Printf("版本号: %d\n", block.Version)
		fmt.Printf("前区块哈希值: %x\n", block.PrevHash)
		fmt.Printf("梅克尔根: %x\n", block.MerkelRoot)
		fmt.Printf("时间戳: %d\n", block.TimeStamp)
		fmt.Printf("难度值(随便写的）: %d\n", block.Difficulty)
		fmt.Printf("随机数 : %d\n", block.Nonce)
		fmt.Printf("当前区块哈希值: %x\n", block.Hash)
		fmt.Printf("区块数据 :%s\n", block.Data)


		if block.PrevHash == nil {
			fmt.Println("区块链遍历结束！")
			break
		}
	}
}
