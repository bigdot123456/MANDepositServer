package ManDepositServer

import (
	"github.com/MatrixAINetwork/go-AIMan/manager"
	"math/big"
	"testing"
)

//活期退选100
func TestSendDepositWithdraw01_v2(t *testing.T) {
	//Man1 := new(big.Int).SetUint64(1e18)
	//amount := new(big.Int).SetUint64(100)
	//amount = amount.Mul(amount,Man1)
	//amount,_:= new(big.Int).SetString("10000453574939770000000",10)
	amount, _ := new(big.Int).SetString("10000400000000000000000", 10)
	err := SendWithDrawTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(0), amount)
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}
}

//活期退选99
func TestSendDepositWithdraw02_v2(t *testing.T) {
	Man1 := new(big.Int).SetUint64(1e18)
	amount := new(big.Int).SetUint64(99)
	amount = amount.Mul(amount, Man1)
	err := SendWithDrawTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(0), amount)
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}
}

//定期仓位1退选99
func TestSendDepositWithdraw03_v2(t *testing.T) {
	Man1 := new(big.Int).SetUint64(1e18)
	amount := new(big.Int).SetUint64(99)
	amount = amount.Mul(amount, Man1)
	err := SendWithDrawTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(1), amount)
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}
}

//定期仓位10退选99
func TestSendDepositWithdraw04_v2(t *testing.T) {
	Man1 := new(big.Int).SetUint64(1e18)
	amount := new(big.Int).SetUint64(99)
	amount = amount.Mul(amount, Man1)
	err := SendWithDrawTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(10), amount)
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}
}

//连续100次活期退选100
func TestSendDepositWithdraw05_v2(t *testing.T) {
	Man1 := new(big.Int).SetUint64(1e18)
	amount := new(big.Int).SetUint64(100)
	amount = amount.Mul(amount, Man1)
	for i := 0; i < 5; i++ {
		err := SendWithDrawTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(0), amount)
		if err != nil && err != errDeposit {
			t.Error(err)
			t.FailNow()
		}
	}
}
