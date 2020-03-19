package ManDepositServer

import (
	"github.com/MatrixAINetwork/go-AIMan/manager"
	"math/big"
	"testing"
)

//直接退款活期
func TestSendDepositRefund01_v2(t *testing.T) {
	err := SendRefundTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(0))
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}
}

//直接退款定期1号仓位
func TestSendDepositRefund02_v2(t *testing.T) {
	err := SendRefundTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(1))
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}
}

//直接退款定期100号仓位
func TestSendDepositRefund03_v2(t *testing.T) {
	err := SendRefundTrans_v2(manager.Self_Manager, "MAN.wipBPMF5AP6R3etpL2LCRFMecL2J", "xxx", big.NewInt(100))
	if err != nil && err != errDeposit {
		t.Error(err)
		t.FailNow()
	}
}
