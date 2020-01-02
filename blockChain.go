package main

import (
	"MyBlockChain/bolt"
	"log"
)

//4. 引入区块链
type BlockChain struct {
	//定义一个区块链数组
	//blocks []*Block
	db *bolt.DB
	tail []byte //存储最后一个区块的哈希
}


const (
	blockChainDB  = "blockChain.db"
	blockBucket="blockBucket"
	)
//5. 定义一个区块链
func NewBlockChain()*BlockChain {
	//创建一个创世块,作为第一个区块添加到区块链

	//return &BlockChain{
	//	blocks:[]*Block{genesisBlock},
	//}
	lastHash:=[]byte{}
	//1.打开数据库
	db,err:=bolt.Open(blockChainDB,0600,nil)

	if err!=nil{
		panic("打开数据库失败")
	}
	defer db.Close()
	//2.找到bucket(如果没有就创建)
	db.Update(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte(blockBucket))
		if bucket==nil{
			//没有抽屉,创建抽屉
			bucket,err=tx.CreateBucket([]byte(blockBucket))
			if err!=nil{
				log.Panic("创建blockBucket失败")
			}
			genesisBlock:=GenesisBlock()
			//3.写数据 存key value 存该区块哈希
			//hash作为key,block字节流作为value
			bucket.Put(genesisBlock.Hash,genesisBlock.toByte())
			bucket.Put([]byte("LastHashKey"),genesisBlock.Hash)
			lastHash=genesisBlock.Hash
		}else{
			lastHash=bucket.Get([]byte("LastHashKey"))
		}

		return nil
	})
	return &BlockChain{db,lastHash}
}
//创世快
func GenesisBlock() *Block {
	return NewBlock("创世区块",[]byte{})
}
//6. 添加区块
func (bc *BlockChain)AddBlock(data string)  {
	//获取最后一个区块
	/*lastBlock:=bc.blocks[len(bc.blocks)-1]
	prevHash:=lastBlock.Hash
	//a. 创建新的区块
	block:=NewBlock(data,prevHash)
	//b. 添加到区块链数组中
	bc.blocks=append(bc.blocks, block)*/
}