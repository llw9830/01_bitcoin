package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)

const reward  = 12.5
// 交易
// 1.定义交易结构
type Transaction struct {
	TXID			[]byte 		// 交易id
	TXInputs		[]TXInput	// 交易输入数组
	TXOutputs		[]TXOutput	// 交易输出数组
}

// 定义交易输入
type TXInput struct {
	// 引用交易id
	TXid		[]byte
	// 引用output索引值
	Index		int64
	// 解锁脚本，用地址来模拟
	Sig 		string
}

// 定义交易输出
type TXOutput struct {
	// 转账金额
	Value		float64
	// 锁定脚本，用地址模拟
	PutKeyHash 	string
}

// 设置交易ID
func (tx *Transaction) SetHash ()  {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	data := buffer.Bytes()
	hash := sha256.Sum256(data)
	tx.TXID = hash[:]
}

// 判断当前交易是否为挖矿交易
func (tx *Transaction)IsCoinbase () bool {
	// 1.交易input只有一个
	// 2.交易id为空
	// 3.交易index为-1
	if len(tx.TXInputs) == 1 && len(tx.TXInputs[0].TXid) == 0 && tx.TXInputs[0].Index == -1 {
		return true
	}
	return false
}

// 2.提供创建交易的方法

// 3.创建挖矿交易
func NewCoinbaseTx(address string, data string) *Transaction {
	// 挖矿交易特点：1.只有一个input；2.无需引用交易id；3.无需引用index
	// 矿工由于挖矿时无需指定签名，所以这个sig字段可以由矿工自行填写数据，一般填写矿池名字
	input := TXInput{
		TXid:   []byte{},
		Index:  -1,
		Sig:    data,
	}
	output := TXOutput{
		Value:      reward,
		PutKeyHash: address,
	}
	// 挖矿交易只有一个input和output
	tx := &Transaction{
		TXID:      	[]byte{},
		TXInputs:  	[]TXInput{input},
		TXOutputs: 	[]TXOutput{output},
	}
	tx.SetHash()
	return tx
}

// 创建普通的转账方法
// 3.创建outputs
// 4.如果有零钱要找零
func NewTransaction(from, to string, amount float64, bc *BlockChain) *Transaction { // amount转账金额
	// 1.找到合理的UTXO集合 map[string]uint64
	utxos, resValue := bc.FindNeedUTXOs(from , amount)

	if resValue < amount {
		fmt.Printf("余额不足，交易失败。")
		return nil
	}

	var inputs 		[]TXInput
	var outputs 	[]TXOutput

	// 2.时间交易输入,将这些UTXO转为input
	for id, indexArray := range utxos {
		for _, i := range indexArray {
			input := TXInput{[]byte(id), int64(i), from}
			inputs = append(inputs, input)
		}
	}

	// 创建能交易输入
	ouput := TXOutput{Value: amount, PutKeyHash: to}
	outputs = append(outputs, ouput)

	// 找零
	if resValue > amount {
		outputs = append(outputs, TXOutput{resValue - amount, from})
	}

	tx := &Transaction{TXID: []byte{},	TXInputs: inputs, TXOutputs: outputs}
	tx.SetHash()
	return tx
}



// 4.根据交易调整程序
