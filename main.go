package main

import "fmt"

type Block struct {
	PrevHash	[]byte
	Hash 		[]byte
	Data 		[]byte
}

func NewBlock(data string, prevBloclHash []byte) *Block {
	hash := Block{
		PrevHash: prevBloclHash,
		Hash:     []byte{}, // TODO
		Data:     []byte(data),
	}
	return &hash
}

func main()  {
	block := NewBlock("a给b比特币！", []byte{})

	fmt.Printf("前区块哈希： %x\n", block.PrevHash)
	fmt.Printf("当前区块哈希： %x\n", block.Hash)
	fmt.Printf("区块数据： %s\n", block.Data)

}