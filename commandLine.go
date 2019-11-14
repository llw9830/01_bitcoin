package main

import "fmt"

// 添加区块
func (cli *CLI) AddBlock (data string)  {
	//cli.bc.AddBlock(data)
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
		fmt.Printf("区块数据 :%s\n", block.Transaction[0].TXInputs[0].Sig)


		if block.PrevHash == nil {
			fmt.Println("区块链遍历结束！")
			break
		}
	}
}

func (cli *CLI) GetBalance (address string) {
	utxos := cli.bc.FindUTXOs(address)
	total := 0.0
	for _, utxo := range utxos {
		total += utxo.Value
	}
	fmt.Printf("\"%s\" 的账户余额为： %f\n", address, total)
}


// 转账
func (cli *CLI) send (from, to string, amount float64, miner, data string) {
	fmt.Printf("from: %s, to: %s, amount: %f, miner: %s, data: %s.\n", from, to, amount, miner, data)
	// TODO
}