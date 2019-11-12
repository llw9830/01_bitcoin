package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// 定义⼀个⼯作量证明的结构ProofOfWork
type ProofOfWork struct {
	block 		*Block //a.	block
	target		big.Int // b.	⽬标值
}

//2.	提供创建POW的函数
func NewProofOfWork(b *Block) *ProofOfWork {
	pow := &ProofOfWork{block: b,}
	// 指定难度值
	targetStr := "0000100000000000000000000000000000000000000000000000000000000000"
	targetInt := big.Int{}
	targetInt.SetString(targetStr, 16) // 将str转为16进制int
	pow.target = targetInt

	return pow
}

//3.	提供计算不断计算hash的哈数
// Run()
func (pow *ProofOfWork)Run () ([]byte, uint64) {
	var nonce uint64
	b := pow.block

	fmt.Println("开始挖矿！")
	for {
		tmp := [][]byte{
			b.PrevHash,
			Uint64oByte(b.Version),
			b.MerkelRoot,
			Uint64oByte(b.TimeStamp),
			Uint64oByte(b.Difficulty),
			Uint64oByte(nonce),
			// 只对区块头做hash，区块头通过MerkelRoot产生影响
			//b.Data,

		}
		// 拼接数据
		blockInfo := bytes.Join(tmp, []byte{})
		// hash运算
		hash := sha256.Sum256(blockInfo)
		tmpInt := big.Int{}
		tmpInt.SetBytes(hash[:])
		//   -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y
		//
		//func (x *Int) Cmp(y *Int) (r int) {
		if tmpInt.Cmp(&pow.target) == -1 {
			fmt.Printf("挖矿成功！ hash: %x, nonce: %d\n", hash, nonce)
			return hash[:], nonce
		} else {
			nonce++
		}
	}
}


// 4.	提供⼀个校验函数
// IsValid()