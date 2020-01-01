package demo

import "fmt"

func totalBtc()  {
	//1.每21万个块减半
	//2. 最初奖励50个比特币
	//3. 用一个循环判断,累加
	total:=0.0
	blockInterval:=21.0//单位 万
	currentReward:=50.0
	for currentReward>0{
		//每个区间内总量
		amount1:=blockInterval*currentReward
		//currentReward/=2 除运算效率低
		currentReward*=0.5
		total+=amount1
	}
	fmt.Println("比特币总量:",total,"万")
}
