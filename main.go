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

// 1.创建区块
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

// 4.引入区块链
type BlockChain struct {
	blocks []*Block
}

// 5.定义一个区块链
func NewBlcokChain() *BlockChain {
	block := GensisBlock()
	return &BlockChain{ blocks: []*Block{block, }}
}

// 创世快
func GensisBlock() *Block {
	return NewBlock("这是一个创世快！", []byte{})
}

func main()  {
	blockChains := NewBlcokChain()
	for i, block := range blockChains.blocks {
		fmt.Printf("=============当前区块高度%d=============\n", i)
		fmt.Printf("前区块哈希： %x\n", block.PrevHash)
		fmt.Printf("当前区块哈希： %x\n", block.Hash)
		fmt.Printf("区块数据： %s\n", block.Data)
	}
}