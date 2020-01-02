package main

import (
	"bytes"
	"encoding/binary"
	"time"
)

//1. 定义结构
type Block struct {
	//+1 版本号
	Version uint64
	//1. 前区块哈希
	PrevHash []byte
	//+2 MerKle根 梅克尔根,一个哈希值
	MerkleRoot []byte
	//+3 时间戳
	TimeStamp uint64
	//+4 难度值
	Difficulty uint64
	//+5 随机数(旷工要找的数据)
	Nonce uint64

	//a. 当前区块哈希 正常比特币区块中无当前区块哈希
	Hash []byte
	//b. 数据
	Data []byte
}

//1. 补充区块字段
//2. 更新计算哈希函数
//3. 优化代码

//辅助函数:实现uint 转[]byte
func UintToByte(num uint64) []byte {
	buffer := bytes.Buffer{}
	err:=binary.Write(&buffer, binary.BigEndian, num)
	if err!=nil{
		panic(err)
	}
	return buffer.Bytes()
}

//2. 创建区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		Version:    0,
		PrevHash:   prevBlockHash,
		MerkleRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 0,
		Nonce:      0,
		Hash:       []byte{},
		Data:       []byte(data),
	}

	//block.SetHash()
	//创建pow对象
	pow:=NewProofOfWork(&block)
	//查找随机数,不停进行哈希运算
	hash,nonce:=pow.Run()
	//根据挖矿结果对区块数据进行更新(补充)
	block.Hash=hash
	block.Nonce=nonce

	return &block
}
func (block *Block)toByte()[]byte  {
	//TODO
	return []byte{}
}
/*//3. 生成哈希
func (block *Block) SetHash() {

	//1. 拼接数据
	blockInfo := []byte{}
	blockInfo = append(blockInfo, UintToByte(block.Version)...)
	blockInfo = append(blockInfo, block.PrevHash...)
	blockInfo = append(blockInfo, block.MerkleRoot...)
	blockInfo = append(blockInfo, UintToByte(block.TimeStamp)...)
	blockInfo = append(blockInfo, UintToByte(block.Difficulty)...)
	blockInfo = append(blockInfo, UintToByte(block.Nonce)...)
	blockInfo = append(blockInfo, block.Data...)

	//tmp:=[][]byte{
		UintToByte(block.Version),
		block.PrevHash,
		block.MerkleRoot,
		UintToByte(block.TimeStamp),
		UintToByte(block.Difficulty),
		UintToByte(block.Nonce),
		block.Data,
	}
	blockInfo:=bytes.Join(tmp,[]byte{})

	//2. sha256
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}*/
