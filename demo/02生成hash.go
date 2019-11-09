package main

import (
	"crypto/sha256"
	"fmt"
)

func main()  {
	// 交易数据
	var data = "silas"

	for i := 1; i < 1000; i++ {
		hash := sha256.Sum256([]byte(data + string(i)))
		fmt.Printf("hash: %x, %d\n", hash, i)
	}
}
