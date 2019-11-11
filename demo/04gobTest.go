package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

// 用gob编码，解码

type Person struct {
	Name	string
	age		uint
}

func main()  {
	var xiaoming Person
	xiaoming.Name = "xiaoming"
	xiaoming.age = 10
	// ============编码
	// 编码数据放到buffer中
	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&xiaoming)
	if err != nil {
		log.Panic("编码出错！", err)
	}
	fmt.Printf("编码后小明：%v\n", buffer.Bytes())


	// ============解码
	decoder := gob.NewDecoder(bytes.NewReader(buffer.Bytes()))
	var tmp Person
	err = decoder.Decode(&tmp)
	if err != nil {
		log.Panic("解码出错！", err)
	}
	fmt.Printf("解码的数据：%v\n", tmp)
}
