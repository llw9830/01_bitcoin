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
	coinbase := NewCoinbaseTx(adderss, "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks")
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

	// map[交易id][]int64
	spentOutput := make(map[string][]int64)

	it := bc.NewIterator()
	for {
		// 1.遍历区块
		block := it.Next()
		// 2.遍历交易

		for _, tx := range block.Transaction {
			fmt.Printf("当前交易id：%x\n", tx.TXID)

			OUTPUT:
			// 3.遍历output,找到和自己相关的utxo
			for i, output := range tx.TXOutputs {
				fmt.Printf("current index：%d\n", i)

				// map[2222] = []int64{0}
				// map[3333] = []int64{0, 1}
				if spentOutput[string(tx.TXID)] != nil {
					for _,  j := range spentOutput[string(tx.TXID)] {
						if int64(i) == j {
							// 当前output已经消耗过不用添加
							continue OUTPUT
						}
					}
				}

				// output和目标地址相同，添加到utxo中
				if output.PutKeyHash == address {
					UTXO = append(UTXO, output)
				}
			}

			// 判断是否为挖矿交易，直接跳过
			if !tx.IsCoinbase() {
				// 4.遍历input， 找到自己花费过的utxo
				for _, input := range tx.TXInputs {
					// 判断当前input和目标（silas）是否一致，相同说明是silas消耗过的output，就加进来
					if input.Sig == address {
						//indeArray := spentOutput[string(input.TXid)]
						//indeArray = append(indeArray, input.Index)
						spentOutput[string(input.TXid)] = append(spentOutput[string(input.TXid)], input.Index)
					}
				}
			} else {
				fmt.Println("这是coinbase,不做input遍历！")
			}
		}

		if len(block.PrevHash) == 0 {
			fmt.Println("区块遍历结束退出！")
			break
		}
	}
	return UTXO
}

// 根据需求找到合理的的utxo
func (bc *BlockChain)FindNeedUTXOs(from string, amount float64) (map[string][]uint64, float64) {
	// 找到的utxo集合
	utxos := make(map[string][]uint64 )
	// 找到的utxo中钱的集合
	var calc float64
	// 标识已经消耗过的utxo
	spentOutput := make(map[string][]int64)
	// 1111111111111111111111111

	it := bc.NewIterator()
	for {
		// 1.遍历区块
		block := it.Next()
		// 2.遍历交易

		for _, tx := range block.Transaction {
			fmt.Printf("当前交易id：%x\n", tx.TXID)

		OUTPUT:
			// 3.遍历output,找到和自己相关的utxo
			for i, output := range tx.TXOutputs {
				fmt.Printf("current index：%d\n", i)

				// map[2222] = []int64{0}
				// map[3333] = []int64{0, 1}
				if spentOutput[string(tx.TXID)] != nil {
					for _,  j := range spentOutput[string(tx.TXID)] {
						if int64(i) == j {
							// 当前output已经消耗过不用添加
							continue OUTPUT
						}
					}
				}

				// output和目标地址相同，添加到utxo中
				if output.PutKeyHash == from {
					//UTXO = append(UTXO, output)
					// 找到自己需要的最少utxo
					// TODO
					if calc >= amount {

						// 1.把utxo加进来
						array := utxos[string(tx.TXID)]
						array = append(array, uint64(i))
						// 2.统计utxo总额
						calc += output.Value
						// 3.比较是否满足转账需求（满足返回，不满足继续）

						if calc >= amount {
							fmt.Printf("找到了满足的金额：%f\n", calc)
							return utxos, calc
						}
					} else {
						fmt.Printf("不满足转账金额，当前总额：%f, 目标金额：%f\n", calc, amount)
					}

				}
			}

			// 判断是否为挖矿交易，直接跳过
			if !tx.IsCoinbase() {
				// 4.遍历input， 找到自己花费过的utxo
				for _, input := range tx.TXInputs {
					// 判断当前input和目标（silas）是否一致，相同说明是silas消耗过的output，就加进来
					if input.Sig == from {
						//indeArray := spentOutput[string(input.TXid)]
						//indeArray = append(indeArray, input.Index)
						spentOutput[string(input.TXid)] = append(spentOutput[string(input.TXid)], input.Index)
					}
				}
			} else {
				fmt.Println("这是coinbase,不做input遍历！")
			}
		}

		if len(block.PrevHash) == 0 {
			fmt.Println("区块遍历结束退出！")
			break
		}
	}
	// 222222222222222222222222222222

	return utxos, calc
}