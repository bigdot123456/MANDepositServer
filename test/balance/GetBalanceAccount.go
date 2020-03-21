package main

import (
	"ManDepositServer"
	"math/big"

	//"ManDepositServer"
	"fmt"
	"github.com/MatrixAINetwork/go-AIMan/manager"
)

func main() {

	var start_num big.Int
	RewardAddr := "MAN.2nRsUetjWAaYUizRkgBxGETimfUTz"
	bignum, _ := new(big.Int).SetString("1000000000000000000", 0)
	nowblocknum, _ := manager.Tom_Manager.Man.GetBlockNumber()
	fmt.Println("当前区块高度", nowblocknum)
	//blocknum=1
	var s int64
	MANRewardList := make(map[string]string)

	const monthValue = 300 * 24 * 25
	s = 100000
	//e=s+monthValue*12
	for i := int64(0); i < 12; i++ {

		start_num.SetInt64(s + i*monthValue)
		blocknum := start_num
		//fmt.Println("获得账号区块高度", blocknum.Int64())
		if blocknum.Cmp(nowblocknum) > 0 {
			panic("error in block setting, exit")
		}

		RewardMount, err := manager.Tom_Manager.Man.GetBalance(RewardAddr, ManDepositServer.Convert(int(blocknum.Int64())))
		MANReward := RewardMount[0].Balance.ToInt()
		MANReward.Quo(MANReward, bignum)
		MANRewardList[blocknum.String()] = MANReward.String()

		//a, _ :=manager.Tom_Manager.Man.GetHashRate()
		if err != nil {
			fmt.Print("error in get mount!")
		} else {
			fmt.Println("地址:", RewardAddr, ":", i+1, "月", "账户余额:", MANReward, "区块高度：", blocknum.Int64())
		}

	}
	for i, l := range MANRewardList {
		fmt.Printf("block %s is %s \n", i, l)
	}

}
