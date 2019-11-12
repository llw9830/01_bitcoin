package main

func main() {
	blockChains := NewBlcokChain()
	cli := CLI{blockChains}
	cli.Run()
}
	/*// 加区块
	blockChains.AddBlock("a向b转10比特币")
	blockChains.AddBlock("a向b转30比特币")
*/



