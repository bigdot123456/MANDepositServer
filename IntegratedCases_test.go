package ManDepositServer

import (
	"github.com/MatrixAINetwork/go-AIMan/manager"
	"github.com/MatrixAINetwork/go-AIMan/waiting"
	"math/big"
	"testing"
	"time"
)

//集成测试用例01
func TestIntegrated01(t *testing.T) {
	Man := new(big.Int).SetUint64(1e18)

	//1.活期抵押矿工：20000MAN
	amount := new(big.Int).SetUint64(20000)
	amount = amount.Mul(amount, Man)
	depositType := big.NewInt(0) //0-活期,1-定期1个月,3-定期3个月,6-定期6个月
	depositRole := big.NewInt(1) // 1是矿工，2是验证者
	err := SendDepositTrans_v2(manager.Self_Manager, "MAN.kMhYY36XgyGA6dikKGkdpJDcVJdG",
		amount, depositType, depositRole, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	//2.追加活期抵押：10000MAN
	amount1 := new(big.Int).SetUint64(10000)
	amount1 = amount.Mul(amount1, Man)
	err = SendDepositTrans_v2(manager.Self_Manager, "MAN.kMhYY36XgyGA6dikKGkdpJDcVJdG",
		amount, depositType, depositRole, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	//3.活期转定期1个月：  20000MAN
	amount2 := new(big.Int).SetUint64(20000)
	amount2 = amount.Mul(amount2, Man)
	err = SendModifyDepositTrans(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(1), amount)
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}

	//4.追加活期抵押：15000MAN
	amount3 := new(big.Int).SetUint64(15000)
	amount3 = amount.Mul(amount3, Man)
	err = SendDepositTrans_v2(manager.Self_Manager, "MAN.kMhYY36XgyGA6dikKGkdpJDcVJdG",
		amount, depositType, depositRole, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	//5.退选、退款定期1个月抵押
	amount4 := new(big.Int).SetUint64(99)
	amount4 = amount.Mul(amount4, Man)
	err = SendWithDrawTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(1), amount)
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}

	wait3 := waiting.NewMultiWaiting(waiting.NewWaitTime(150 * time.Second))
	wait3.Waiting()

	err = SendRefundTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(1))
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}
}

//集成测试用例02
func TestIntegrated02(t *testing.T) {
	Man := new(big.Int).SetUint64(1e18)

	//1.定期抵押1个月矿工：20000MAN
	amount := new(big.Int).SetUint64(20000)
	amount = amount.Mul(amount, Man)
	depositType := big.NewInt(1) //0-活期,1-定期1个月,3-定期3个月,6-定期6个月
	depositRole := big.NewInt(1) // 1是矿工，2是验证者
	err := SendDepositTrans_v2(manager.Self_Manager, "MAN.kMhYY36XgyGA6dikKGkdpJDcVJdG",
		amount, depositType, depositRole, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	//2.追加活期抵押：10000MAN
	amount1 := new(big.Int).SetUint64(10000)
	amount1 = amount.Mul(amount1, Man)
	depositType = big.NewInt(0) //0-活期,1-定期1个月,3-定期3个月,6-定期6个月
	err = SendDepositTrans_v2(manager.Self_Manager, "MAN.kMhYY36XgyGA6dikKGkdpJDcVJdG",
		amount, depositType, depositRole, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	//3.退选、退款定期1个月抵押
	amount4 := new(big.Int).SetUint64(99)
	amount4 = amount.Mul(amount4, Man)
	err = SendWithDrawTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(1), amount)
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}

	wait3 := waiting.NewMultiWaiting(waiting.NewWaitTime(150 * time.Second))
	wait3.Waiting()

	err = SendRefundTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(1))
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}
}

//集成测试用例03
func TestIntegrated03(t *testing.T) {
	Man := new(big.Int).SetUint64(1e18)

	//1.定期抵押1个月矿工：20000MAN
	amount := new(big.Int).SetUint64(20000)
	amount = amount.Mul(amount, Man)
	depositType := big.NewInt(1) //0-活期,1-定期1个月,3-定期3个月,6-定期6个月
	depositRole := big.NewInt(1) // 1是矿工，2是验证者
	err := SendDepositTrans_v2(manager.Self_Manager, "MAN.kMhYY36XgyGA6dikKGkdpJDcVJdG",
		amount, depositType, depositRole, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	//2.追加定期抵押3个月：10000MAN
	amount1 := new(big.Int).SetUint64(10000)
	amount1 = amount.Mul(amount1, Man)
	depositType = big.NewInt(3) //0-活期,1-定期1个月,3-定期3个月,6-定期6个月
	err = SendDepositTrans_v2(manager.Self_Manager, "MAN.kMhYY36XgyGA6dikKGkdpJDcVJdG",
		amount, depositType, depositRole, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	//3.追加定期抵押6个月：10000MAN
	amount2 := new(big.Int).SetUint64(10000)
	amount2 = amount.Mul(amount2, Man)
	depositType = big.NewInt(6) //0-活期,1-定期1个月,3-定期3个月,6-定期6个月
	err = SendDepositTrans_v2(manager.Self_Manager, "MAN.kMhYY36XgyGA6dikKGkdpJDcVJdG",
		amount, depositType, depositRole, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	//4.追加定期抵押12个月：10000MAN
	amount3 := new(big.Int).SetUint64(10000)
	amount3 = amount.Mul(amount3, Man)
	depositType = big.NewInt(12) //0-活期,1-定期1个月,3-定期3个月,6-定期6个月
	err = SendDepositTrans_v2(manager.Self_Manager, "MAN.kMhYY36XgyGA6dikKGkdpJDcVJdG",
		amount, depositType, depositRole, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	//5.追加活期抵押：10000MAN
	amount4 := new(big.Int).SetUint64(10000)
	amount4 = amount.Mul(amount4, Man)
	depositType = big.NewInt(0) //0-活期,1-定期1个月,3-定期3个月,6-定期6个月
	err = SendDepositTrans_v2(manager.Self_Manager, "MAN.kMhYY36XgyGA6dikKGkdpJDcVJdG",
		amount, depositType, depositRole, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	//6.退选、退款定期1、3、6、12个月抵押和活期抵押
	err = SendWithDrawTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(1), amount)
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}

	err = SendWithDrawTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(2), amount)
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}

	err = SendWithDrawTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(3), amount)
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}

	err = SendWithDrawTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(4), amount)
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}

	err = SendWithDrawTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(0), amount4)
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}

	wait3 := waiting.NewMultiWaiting(waiting.NewWaitTime(1300 * time.Second))
	wait3.Waiting()

	err = SendRefundTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(1))
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}
	err = SendRefundTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(2))
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}
	err = SendRefundTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(3))
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}
	err = SendRefundTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(4))
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}
	err = SendRefundTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(0))
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}
}

//集成测试用例04
func TestIntegrated04(t *testing.T) {
	Man := new(big.Int).SetUint64(1e18)

	//1.活期抵押验证者：20000MAN
	amount := new(big.Int).SetUint64(20000)
	amount = amount.Mul(amount, Man)
	depositType := big.NewInt(0) //0-活期,1-定期1个月,3-定期3个月,6-定期6个月
	depositRole := big.NewInt(2) // 1是矿工，2是验证者
	err := SendDepositTrans_v2(manager.Self_Manager, "MAN.kMhYY36XgyGA6dikKGkdpJDcVJdG",
		amount, depositType, depositRole, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	//2.追加活期抵押：10000MAN
	amount1 := new(big.Int).SetUint64(10000)
	amount1 = amount.Mul(amount1, Man)
	depositType = big.NewInt(0) //0-活期,1-定期1个月,3-定期3个月,6-定期6个月
	err = SendDepositTrans_v2(manager.Self_Manager, "MAN.kMhYY36XgyGA6dikKGkdpJDcVJdG",
		amount, depositType, depositRole, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	//3.活期转定期1个月：  20000MAN
	amount2 := new(big.Int).SetUint64(20000)
	amount2 = amount.Mul(amount2, Man)
	err = SendModifyDepositTrans(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(1), amount)
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}

	//4.追加活期抵押：15000MAN
	amount3 := new(big.Int).SetUint64(15000)
	amount3 = amount.Mul(amount3, Man)
	depositType = big.NewInt(0) //0-活期,1-定期1个月,3-定期3个月,6-定期6个月
	err = SendDepositTrans_v2(manager.Self_Manager, "MAN.kMhYY36XgyGA6dikKGkdpJDcVJdG",
		amount, depositType, depositRole, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	//5.退选、退款定期1个月抵押和活期抵押
	err = SendWithDrawTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(1), amount)
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}

	amount4 := new(big.Int).SetUint64(25000)
	amount4 = amount.Mul(amount4, Man)
	err = SendWithDrawTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(0), amount4)
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}

	wait3 := waiting.NewMultiWaiting(waiting.NewWaitTime(150 * time.Second))
	wait3.Waiting()

	err = SendRefundTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(1))
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}

	err = SendRefundTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(0))
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}
}

//集成测试用例05
func TestIntegrated05(t *testing.T) {
	Man := new(big.Int).SetUint64(1e18)

	//1.定期抵押1个月验证者：20000MAN
	amount := new(big.Int).SetUint64(20000)
	amount = amount.Mul(amount, Man)
	depositType := big.NewInt(1) //0-活期,1-定期1个月,3-定期3个月,6-定期6个月
	depositRole := big.NewInt(2) // 1是矿工，2是验证者
	err := SendDepositTrans_v2(manager.Self_Manager, "MAN.kMhYY36XgyGA6dikKGkdpJDcVJdG",
		amount, depositType, depositRole, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	//2.追加活期抵押：10000MAN
	amount1 := new(big.Int).SetUint64(10000)
	amount1 = amount.Mul(amount1, Man)
	depositType = big.NewInt(0) //0-活期,1-定期1个月,3-定期3个月,6-定期6个月
	err = SendDepositTrans_v2(manager.Self_Manager, "MAN.kMhYY36XgyGA6dikKGkdpJDcVJdG",
		amount, depositType, depositRole, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	//3.退选、退款定期1个月抵押和活期抵押
	err = SendWithDrawTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(1), amount)
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}

	amount4 := new(big.Int).SetUint64(10000)
	amount4 = amount.Mul(amount4, Man)
	err = SendWithDrawTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(0), amount4)
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}

	wait3 := waiting.NewMultiWaiting(waiting.NewWaitTime(150 * time.Second))
	wait3.Waiting()

	err = SendRefundTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(1))
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}

	err = SendRefundTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(0))
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}
}

//集成测试用例06
func TestIntegrated06(t *testing.T) {
	Man := new(big.Int).SetUint64(1e18)

	//1.定期抵押1个月矿工：20000MAN
	amount := new(big.Int).SetUint64(20000)
	amount = amount.Mul(amount, Man)
	depositType := big.NewInt(1) //0-活期,1-定期1个月,3-定期3个月,6-定期6个月
	depositRole := big.NewInt(2) // 1是矿工，2是验证者
	err := SendDepositTrans_v2(manager.Self_Manager, "MAN.kMhYY36XgyGA6dikKGkdpJDcVJdG",
		amount, depositType, depositRole, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	//2.追加定期抵押3个月：10000MAN
	amount1 := new(big.Int).SetUint64(10000)
	amount1 = amount.Mul(amount1, Man)
	depositType = big.NewInt(3) //0-活期,1-定期1个月,3-定期3个月,6-定期6个月
	err = SendDepositTrans_v2(manager.Self_Manager, "MAN.kMhYY36XgyGA6dikKGkdpJDcVJdG",
		amount, depositType, depositRole, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	//3.追加定期抵押6个月：10000MAN
	amount2 := new(big.Int).SetUint64(10000)
	amount2 = amount.Mul(amount2, Man)
	depositType = big.NewInt(6) //0-活期,1-定期1个月,3-定期3个月,6-定期6个月
	err = SendDepositTrans_v2(manager.Self_Manager, "MAN.kMhYY36XgyGA6dikKGkdpJDcVJdG",
		amount, depositType, depositRole, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	//4.追加定期抵押12个月：10000MAN
	amount3 := new(big.Int).SetUint64(10000)
	amount3 = amount.Mul(amount3, Man)
	depositType = big.NewInt(12) //0-活期,1-定期1个月,3-定期3个月,6-定期6个月
	err = SendDepositTrans_v2(manager.Self_Manager, "MAN.kMhYY36XgyGA6dikKGkdpJDcVJdG",
		amount, depositType, depositRole, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	//5.追加活期抵押：10000MAN
	amount4 := new(big.Int).SetUint64(10000)
	amount4 = amount.Mul(amount4, Man)
	depositType = big.NewInt(0) //0-活期,1-定期1个月,3-定期3个月,6-定期6个月
	err = SendDepositTrans_v2(manager.Self_Manager, "MAN.kMhYY36XgyGA6dikKGkdpJDcVJdG",
		amount, depositType, depositRole, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	//6.退选、退款定期1、3、6、12个月抵押和活期抵押
	err = SendWithDrawTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(1), amount)
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}

	err = SendWithDrawTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(2), amount)
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}

	err = SendWithDrawTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(3), amount)
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}

	err = SendWithDrawTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(4), amount)
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}

	err = SendWithDrawTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(0), amount4)
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}

	wait3 := waiting.NewMultiWaiting(waiting.NewWaitTime(1300 * time.Second))
	wait3.Waiting()

	err = SendRefundTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(1))
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}
	err = SendRefundTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(2))
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}
	err = SendRefundTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(3))
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}
	err = SendRefundTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(4))
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}
	err = SendRefundTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(0))
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}
}
