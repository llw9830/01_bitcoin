package main

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	PrevHash	[]byte
	Hash 		[]byte
	Data 		[]byte
}

// 创建区块
func NewBlock(data string, prevBloclHash []byte) *Block {
	block := Block{
		PrevHash: prevBloclHash,
		Hash:     []byte{},
		Data:     []byte(data),
	}
	block.SetHash()
	return &block
}

// 3. 生成hash
func (b *Block)SetHash () {
	// 拼装数据
	blockInfo := append(b.PrevHash, b.Data...)
	// sha256
	hash := sha256.Sum256(blockInfo)
	b.Hash =  hash[:]
}

func main()  {
	block := NewBlock("a给b比特币！", []byte{})

	fmt.Printf("前区块哈希： %x\n", block.PrevHash)
	fmt.Printf("当前区块哈希： %x\n", block.Hash)
	fmt.Printf("区块数据： %s\n", block.Data)

}