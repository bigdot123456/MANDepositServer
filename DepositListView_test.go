package ManDepositServer

import (
	"fmt"
	"github.com/MatrixAINetwork/go-AIMan/manager"
	"github.com/MatrixAINetwork/go-AIMan/transactions"
	"math/big"
	"strconv"
	"testing"
)

func TestViewDepositList_v2(t *testing.T) {
	ManAddr, err := transactions.GetDepositList_v2(manager.Tom_Manager)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	totalAmount := big.NewInt(0)

	blocknum, _ := manager.Tom_Manager.Man.GetBlockNumber()

	fmt.Println("区块高度", blocknum)

	if ManAddr != nil {
		for _, item := range ManAddr.([]string) {
			info, err := manager.Tom_Manager.Man.GetDepositInfo(item, Convert(int(blocknum.Int64())))
			if err != nil {
				t.Error(err)
				t.FailNow()
			}

			bignum, _ := new(big.Int).SetString("1000000000000000000", 0)
			//info..Quo(info.Amount,bignum)
			//info.Dpstmsg[0].DepositAmount.Quo(info.Dpstmsg[0].DepositAmount,bignum)

			fmt.Println("------------------start 账户地址", info.AddressA0, "角色", info.Role)

			itemAmount := big.NewInt(0)
			for i, item := range info.Dpstmsg {
				//itemAmount = itemAmount.Sub(itemAmount.Add(item.DepositAmount,item.Interest), item.Slash)
				fmt.Println("抵押值", item.DepositAmount.Quo(item.DepositAmount, bignum), "利息", item.Interest.Quo(item.Interest, bignum), "惩罚", item.Slash, "抵押类型", item.DepositType)
				itemAmount = itemAmount.Add(itemAmount, item.DepositAmount)
				for _, item2 := range item.WithDrawInfolist {
					if i == 0 {
						itemAmount = itemAmount.Add(itemAmount, item2.WithDrawAmount)
					}

				}
			}

			totalAmount = totalAmount.Add(totalAmount, itemAmount)
			fmt.Println("------------------end 账户地址", info.AddressA0, "总金额", itemAmount)
		}

		DAmount, _ := manager.Tom_Manager.Man.GetBalance("MAN.1111111111111111111B8", Convert(int(blocknum.Int64())))
		fmt.Println("抵押总金额", totalAmount, "0x000...a账户余额", DAmount[0].Balance.ToInt())
	}
}
