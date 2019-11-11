package main

import (
	"github.com/boltdb/bolt"
	"log"
)

func main()  {
	// 打开数据库
	db, err := bolt.Open("bolttest.db", 0600, nil)
	if err != nil {
		log.Panic("打开数据库失败")
	}

	// 操作数据库
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("b1"))
		if bucket == nil{
			bucket, err = tx.CreateBucket([]byte("b1"))
			if err != nil {
				log.Panic("创建bucket（b1）失败！")
			}
		}
		err = bucket.Put([]byte("1"), []byte("hello"))
		err = bucket.Put([]byte("2"), []byte("world!"))
		return nil
	})


}
