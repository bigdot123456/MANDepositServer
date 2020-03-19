package ManDepositServer

import (
	"github.com/MatrixAINetwork/go-AIMan/manager"
	"math/big"
	"testing"
)

//活期转定期1w 3个月
func TestSendModifyDepositType01_v2(t *testing.T) {
	Man1 := new(big.Int).SetUint64(1e18)
	amount := new(big.Int).SetUint64(10000)
	amount = amount.Mul(amount, Man1)
	err := SendModifyDepositTrans(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(3), amount)
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}
}

//活期转定期1w 6个月
func TestSendModifyDepositType02_v2(t *testing.T) {
	Man1 := new(big.Int).SetUint64(1e18)
	amount := new(big.Int).SetUint64(10000)
	amount = amount.Mul(amount, Man1)
	err := SendModifyDepositTrans(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(6), amount)
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}
}

//活期转定期2w 3个月
func TestSendModifyDepositType03_v2(t *testing.T) {
	Man1 := new(big.Int).SetUint64(1e18)
	amount := new(big.Int).SetUint64(20000)
	amount = amount.Mul(amount, Man1)
	err := SendModifyDepositTrans(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(3), amount)
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}
}

//活期转1个月定期9999
func TestSendModifyDepositType04_v2(t *testing.T) {
	Man1 := new(big.Int).SetUint64(1e18)
	amount := new(big.Int).SetUint64(9999)
	amount = amount.Mul(amount, Man1)
	err := SendModifyDepositTrans(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(1), amount)
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}
}

//活期转活期1w
func TestSendModifyDepositType05_v2(t *testing.T) {
	Man1 := new(big.Int).SetUint64(1e18)
	amount := new(big.Int).SetUint64(10000)
	amount = amount.Mul(amount, Man1)
	err := SendModifyDepositTrans(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(0), amount)
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}
}
