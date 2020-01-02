package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

//定义一个工作量证明的结构ProofOfWork
type ProofOfWork struct {
	//a. block
	block *Block
	//b. 目标值
	target *big.Int
}
//2. 提供创建POW的函数
//NewProofOfWork(参数)
func NewProofOfWork(block *Block) *ProofOfWork {
	pow:=ProofOfWork{
		block:  block,
	}
	//指定难度值 string类型需转换
	targetStr:="0000f00000000000000000000000000000000000000000000000000000000000"
	//辅助变量,string难度值转big.Int
	tmpInt:=big.Int{}
	//指定16进制
	tmpInt.SetString(targetStr,16)
	pow.target=&tmpInt
	return &pow
}
//3. 提供计算不断计算hash的哈数
//Run()
func (pow *ProofOfWork)Run()([]byte,uint64)  {

	var nonce uint64
	block:=pow.block
	hash:=[32]byte{}
	fmt.Println("开始挖矿...")
	for {
		//1. 拼装数据(区块数据,nonce)
		tmp:=[][]byte{
			UintToByte(block.Version),
			block.PrevHash,
			block.MerkleRoot,
			UintToByte(block.TimeStamp),
			UintToByte(block.Difficulty),
			UintToByte(nonce),
			block.Data,
		}
		blockInfo:=bytes.Join(tmp,[]byte{})

		//2. 做哈希运算
		hash=sha256.Sum256(blockInfo)
		//3. 比对目标值
		tmpInt:=big.Int{}
		//将hash数组转换未big.Int
		tmpInt.SetBytes(hash[:])

		//比较当前哈希值与目标哈希值,小于则停止,否则继续
		//   -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y
		//func (x *Int) Cmp(y *Int) (r int) {
		if tmpInt.Cmp(pow.target)==-1{
			//a. 找到退出返回
			fmt.Printf("挖矿成功! hash : %x, nonce : %d \n",hash,nonce)
			return hash[:],nonce
		}else{
			//b. 没找到,继续找,随机数+1
			nonce++
		}
	}
	//return []byte("HelloWorld"),0

}
//4. 提供一个校验函数
//IsValid()