package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"log"
	"time"
)
// 0.定义结构体
type Block struct {
	// 版本号
	Version		uint64
	// 前区块hash
	PrevHash	[]byte
	// Merker根
	MerkelRoot	[]byte
	// 时间戳
	TimeStamp	uint64
	// 难度值
	Difficulty	uint64
	// 随机数
	Nonce		uint64

	// 当前区块hash值
	Hash 		[]byte
	// 数据
	//Data 		[]byte
	Transaction			[]*Transaction
}

func Uint64oByte(num uint64) []byte {
	var b bytes.Buffer
	err := binary.Write(&b, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return b.Bytes()
}

// 1.创建区块
func NewBlock(txs []*Transaction, prevBlockHash []byte) *Block {
	block := Block{
		Version:    00,
		PrevHash:   prevBlockHash,
		MerkelRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 0,
		Nonce:      0,
		Hash:       []byte{},
		//Data:       []byte(data),
		Transaction:			txs,
	}
	block.MerkelRoot = block.MakeMerkelRoot()

	//block.SetHash()
	pow := NewProofOfWork(&block)
	// 查找随机数，不停继续hash运算
	hash, nonce := pow.Run()
	// 工具挖矿结果对block进行更新补充
	block.Hash = hash
	block.Nonce = nonce
	
	return &block
}

// 序列化
func (block *Block)Serialize() []byte  {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(block)
	if err != nil {
		log.Panic("编码出错！", err)
	}
	return buffer.Bytes()
}

// 反序列化
func Deserialize(data []byte) Block  {
	// ============解码
	decoder := gob.NewDecoder(bytes.NewReader(data))
	var block Block
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic("解码出错！", err)
	}
	return block
}

//
//// 3. 生成hash
//func (b *Block)SetHash () {
//	/*// 拼装数据
//	blockInfo := append(b.PrevHash, b.Data...)
//	blockInfo = append(blockInfo, Uint64oByte(b.Version)...)
//	blockInfo = append(blockInfo, b.PrevHash...)
//	blockInfo = append(blockInfo, b.MerkelRoot...)
//	blockInfo = append(blockInfo, Uint64oByte(b.TimeStamp)...)
//	blockInfo = append(blockInfo, Uint64oByte(b.Difficulty)...)
//	blockInfo = append(blockInfo, Uint64oByte(b.Nonce)...)*/
//	tmp := [][]byte{
//		b.PrevHash,
//		b.Data,
//		Uint64oByte(b.Version),
//		b.MerkelRoot,
//		Uint64oByte(b.TimeStamp),
//		Uint64oByte(b.Difficulty),
//		Uint64oByte(b.Nonce),
//	}
//	// 将二维的数组转为一维的
//	blockInfo := bytes.Join(tmp, []byte{})
//	// sha256
//	hash := sha256.Sum256(blockInfo)
//	b.Hash =  hash[:]
//}

// 模拟梅克尔根，只做简单拼接，不做二叉树处理
func (b *Block) MakeMerkelRoot () []byte {
	// TODO
	return []byte{}
}