package main

import (
	"MyBlockChain/bolt"
	"fmt"
	"log"
)

func main(){
	//1.打开数据库
	db,err:=bolt.Open("test.db",0600,nil)

	if err!=nil{
		panic("打开数据库失败")
	}
	defer db.Close()
	//2.找到bucket(如果没有就创建)
	db.Update(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte("b1"))
		if bucket==nil{
			//没有抽屉,创建抽屉
			bucket,err=tx.CreateBucket([]byte("b1"))
			if err!=nil{
				log.Panic("创建bucket(b1)失败")
			}
		}
		//3.写数据
		bucket.Put([]byte("11111"),[]byte("Hello"))
		bucket.Put([]byte("22222"),[]byte("World"))
		return nil
	})
	//4.读数据(验证)
	db.View(func(tx *bolt.Tx) error {
		//1. 找到bucket,没有报错退出
		bucket:=tx.Bucket([]byte("b1"))
		if bucket==nil{
			log.Panic("bucket b1 不存在")
		}
		//2. 读取数据
		v1:=bucket.Get([]byte("11111"))
		v2:=bucket.Get([]byte("22222"))
		fmt.Printf("%s\n",v1)
		fmt.Printf("%s\n",v2)
		return nil
	})

}