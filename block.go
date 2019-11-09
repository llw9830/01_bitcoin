package main

import (
	"crypto/sha256"
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
