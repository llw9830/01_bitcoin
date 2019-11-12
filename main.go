package main

import "fmt"

func main()  {
	blockChains := NewBlcokChain()
	// 加区块
	blockChains.AddBlock("a向b转10比特币")
	blockChains.AddBlock("a向b转30比特币")

	it := blockChains.NewIterator()
	// 调用迭代器
	for  {
		// 翻回去看左移
		block := it.Next()
		fmt.Printf("============================\n\n")
		fmt.Printf("前区块哈希： %x\n", block.PrevHash)
		fmt.Printf("当前区块哈希： %x\n", block.Hash)
		fmt.Printf("区块数据： %s\n", block.Data)

		if block.PrevHash == nil {
			break
		}
	}

	/*for i, block := range blockChains.blocks {
		fmt.Printf("=============当前区块高度%d=============\n", i)
		fmt.Printf("前区块哈希： %x\n", block.PrevHash)
		fmt.Printf("当前区块哈希： %x\n", block.Hash)
		fmt.Printf("区块数据： %s\n", block.Data)
	}*/
}