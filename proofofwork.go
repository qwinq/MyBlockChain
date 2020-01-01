package main

import "math/big"

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
	//TODO
	return []byte("HelloWorld"),0
}
//4. 提供一个校验函数
//IsValid()