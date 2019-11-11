package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	// 打开数据库
	db, err := bolt.Open("bolttest.db", 0600, nil)
	if err != nil {
		log.Panic("打开数据库失败")
	}
	defer db.Close()

	// 操作数据库
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("b1"))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte("b1"))
			if err != nil {
				log.Panic("创建bucket（b1）失败！")
			}
		}
		// 写
		err = bucket.Put([]byte("1"), []byte("hello"))
		err = bucket.Put([]byte("2"), []byte("world!"))
		return nil
	})

	// 读
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("b1"))
		if err != nil {
			log.Panic("bucket b1 不应该为空，请检查！")
		}
		s1 := bucket.Get([]byte("1"))
		s2 := bucket.Get([]byte("2"))
		fmt.Printf("s1: %s\n", s1)
		fmt.Printf("s2: %s\n", s2)
		return nil
	})

}
