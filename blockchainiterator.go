package main

import (
	"github.com/boltdb/bolt"
	"log"
)

// 区块链迭代器
type BlockChainIterator struct {
	db						*bolt.DB // 数据库
	currentHashPointer		[]byte // 游标，
}

func (bc *BlockChain) NewIterator () *BlockChainIterator {
	return &BlockChainIterator{
		db:						bc.db,
		// 指向最后一个区块，
		currentHashPointer:		bc.tail,
	}
}

// 1. 返回当前区块，
// 2. 指针前移
func (it *BlockChainIterator) Next () *Block {
	var block Block
	it.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("blockchainiterator Next() 的 bucket 不应该为空！")
		}
		// 拿到当前区块
		blocktmp := bucket.Get(it.currentHashPointer)
		// 解码
		block = Deserialize(blocktmp)
		// 前区块hash,指针前移
		it.currentHashPointer = block.PrevHash
		return nil
	})
	return &block
}