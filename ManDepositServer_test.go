package ManDepositServer

import (
	"github.com/MatrixAINetwork/go-AIMan/manager"
	"math/big"
	"testing"
)

//10w抵押6个月定期矿工
func TestSendDeposit01_v2(t *testing.T) {
	Man1 := new(big.Int).SetUint64(1e18)
	amount := new(big.Int).SetUint64(100000)
	amount = amount.Mul(amount, Man1)
	depositType := big.NewInt(6) //0-活期,1-定期1个月,3-定期3个月,6-定期6个月
	depositRole := big.NewInt(1) // 1是矿工，2是验证者
	err := SendDepositTrans_v2(manager.Jerry_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J",
		amount, depositType, depositRole, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

//1w抵押活期矿工
func TestSendDeposit02_v2(t *testing.T) {
	Man1 := new(big.Int).SetUint64(1e18)
	amount := new(big.Int).SetUint64(10000)
	amount = amount.Mul(amount, Man1)
	//amount,_:= new(big.Int).SetString("10000453574939769761612",10)
	depositType := big.NewInt(0) //0-活期,1-定期1个月,3-定期3个月,6-定期6个月
	depositRole := big.NewInt(1) // 1是矿工，2是验证者
	err := SendDepositTrans_v2(manager.Jerry_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J",
		amount, depositType, depositRole, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

//0活期抵押验证者
func TestSendDeposit03_v2(t *testing.T) {

	amount := new(big.Int).SetUint64(0)
	depositType := big.NewInt(0) //0-活期,1-定期1个月,3-定期3个月,6-定期6个月
	depositRole := big.NewInt(2) // 1是矿工，2是验证者
	err := SendDepositTrans_v2(manager.Self_Manager, "MAN.46nMLcK7y24NL47hPLt1xDjbxrfM5",
		amount, depositType, depositRole, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

//0元1月定期抵押矿工
func TestSendDeposit04_v2(t *testing.T) {
	amount := new(big.Int).SetUint64(0)
	depositType := big.NewInt(0) //0-活期,1-定期1个月,3-定期3个月,6-定期6个月
	depositRole := big.NewInt(1) // 1是矿工，2是验证者
	err := SendDepositTrans_v2(manager.Self_Manager, "MAN.46nMLcK7y24NL47hPLt1xDjbxrfM5",
		amount, depositType, depositRole, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

//连续抵押矿工，1w 定期1个月，1w，定期3个月，1w活期
func TestSendDeposit05_v2(t *testing.T) {
	Man1 := new(big.Int).SetUint64(1e18)
	amount := new(big.Int).SetUint64(10000)
	amount = amount.Mul(amount, Man1)
	depositType := big.NewInt(1) //0-活期,1-定期1个月,3-定期3个月,6-定期6个月
	depositRole := big.NewInt(1) // 1是矿工，2是验证者

	err := SendDepositTrans_v2(manager.Self_Manager, "MAN.46nMLcK7y24NL47hPLt1xDjbxrfM5",
		amount, depositType, depositRole, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	depositType = big.NewInt(3)
	err = SendDepositTrans_v2(manager.Self_Manager, "MAN.46nMLcK7y24NL47hPLt1xDjbxrfM5",
		amount, depositType, depositRole, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	depositType = big.NewInt(0)
	err = SendDepositTrans_v2(manager.Self_Manager, "MAN.46nMLcK7y24NL47hPLt1xDjbxrfM5",
		amount, depositType, depositRole, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

//连续抵押，1w 定期1个月矿工，1w 3个月矿工，8w活期验证者
func TestSendDeposit06_v2(t *testing.T) {
	Man1 := new(big.Int).SetUint64(1e18)
	amount := new(big.Int).SetUint64(10000)
	amount = amount.Mul(amount, Man1)
	depositType := big.NewInt(1) //0-活期,1-定期1个月,3-定期3个月,6-定期6个月
	depositRole := big.NewInt(1) // 1是矿工，2是验证者

	err := SendDepositTrans_v2(manager.Self_Manager, "MAN.46nMLcK7y24NL47hPLt1xDjbxrfM5",
		amount, depositType, depositRole, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	depositType = big.NewInt(3)
	err = SendDepositTrans_v2(manager.Self_Manager, "MAN.46nMLcK7y24NL47hPLt1xDjbxrfM5",
		amount, depositType, depositRole, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	amount2 := new(big.Int).SetUint64(80000)
	amount = amount.Mul(amount2, Man1)
	depositType = big.NewInt(0)
	depositRole = big.NewInt(2)
	err = SendDepositTrans_v2(manager.Self_Manager, "MAN.46nMLcK7y24NL47hPLt1xDjbxrfM5",
		amount, depositType, depositRole, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

//1w定期9个月/5个月抵押
func TestSendDeposit07_v2(t *testing.T) {
	Man1 := new(big.Int).SetUint64(1e18)
	amount := new(big.Int).SetUint64(10000)
	amount = amount.Mul(amount, Man1)
	depositType := big.NewInt(9) //0-活期,1-定期1个月,3-定期3个月,6-定期6个月
	depositRole := big.NewInt(1) // 1是矿工，2是验证者

	err := SendDepositTrans_v2(manager.Self_Manager, "MAN.46nMLcK7y24NL47hPLt1xDjbxrfM5",
		amount, depositType, depositRole, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	depositType = big.NewInt(5)
	err = SendDepositTrans_v2(manager.Self_Manager, "MAN.46nMLcK7y24NL47hPLt1xDjbxrfM5",
		amount, depositType, depositRole, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
