package main

import (
	"fmt"
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
func NewBlcokChain(address string) *BlockChain {
	var lastHashKey []byte
	// 打开数据库
	db, err := bolt.Open(blockChainDB, 0600, nil)
	if err != nil {
		log.Panic("打开数据库失败")
	}
	//defer db.Close()

	// 操作数据库
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				log.Panic("创建bucket（b1）失败！")
			}
			// 创世快
			block := GensisBlock(address)
			fmt.Printf("创世块：%s\n", block)

			// 写数据
			bucket.Put(block.Hash, block.Serialize())
			bucket.Put([]byte("LastHashKey"), block.Hash)
			lastHashKey = block.Hash

			// 测试序列化反序列化
			//fmt.Printf( "block info : %s\n", Deserialize(bucket.Get(block.Hash)) )
		} else {
			lastHashKey = bucket.Get([]byte("LastHashKey"))
		}
		return nil
	})
	return &BlockChain{db, lastHashKey}
}

// 定义一个创世快
func GensisBlock(adderss string) *Block {
	coinbase := NewCoinbaseTx(adderss, "这是一个创世快")
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

// 5. 添加区块
func (bc *BlockChain) AddBlock (txs []*Transaction) {
	// 前区快hash
	db := bc.db
	lastHash := bc.tail // 最后一个区块的hash

	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("bucket 不应该为空！")
		}

		// a.创建新区块
		block := NewBlock(txs, lastHash)
		// b.添加到区块链db中
		bucket.Put(block.Hash, block.Serialize())
		bucket.Put([]byte("LastHashKey"), block.Hash)

		// 更新内存中的区块链
		bc.tail = block.Hash

		return  nil
	})
}

// 找到指定地址的所有utxo
func (bc *BlockChain) FindUTXOs (address string) []TXOutput {
	var UTXO []TXOutput
	// TODO

	return UTXO
}
