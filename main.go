package main


func main()  {
	blockChains := NewBlcokChain()
	// 加区块
	blockChains.AddBlock("a向b转10比特币")
	blockChains.AddBlock("a向b转30比特币")

	/*for i, block := range blockChains.blocks {
		fmt.Printf("=============当前区块高度%d=============\n", i)
		fmt.Printf("前区块哈希： %x\n", block.PrevHash)
		fmt.Printf("当前区块哈希： %x\n", block.Hash)
		fmt.Printf("区块数据： %s\n", block.Data)
	}*/
}