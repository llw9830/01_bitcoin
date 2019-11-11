package main

import (
	"github.com/boltdb/bolt"
	"log"
)

const blockChainDB = "blockChain.db"
const blockBucket  = "blockBucket"
// 4.引入区块链
type BlockChain struct {
	//blocks []*Block
	db 		*bolt.DB
	tail	[]byte // 存储最后一个区块的hash
}

// 定义一个区块链
func NewBlcokChain() *BlockChain {
	var lastHashKey []byte
	// 打开数据库
	db, err := bolt.Open(blockChainDB, 0600, nil)
	if err != nil {
		log.Panic("打开数据库失败")
	}
	defer db.Close()

	// 操作数据库
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				log.Panic("创建bucket（b1）失败！")
			}
			// 创世快
			block := GensisBlock()
			// 写数据
			bucket.Put(block.Hash, block.toByte())
			bucket.Put([]byte("LastHashKey"), block.Hash)
			lastHashKey = block.Hash
		} else {
			lastHashKey = bucket.Get([]byte("LastHashKey"))
		}
		return nil
	})
	return &BlockChain{db, lastHashKey}
}

// 定义一个创世快
func GensisBlock() *Block {
	return NewBlock("这是一个创世快！", []byte{})
}

// 5. 添加区块
func (bc *BlockChain) AddBlock (data string) {
	/*// 前区快hash
	prevHash := bc.blocks[len(bc.blocks)-1].Hash
	// 添加到链中
	bc.blocks = append(bc.blocks, NewBlock(data, prevHash))*/
}
