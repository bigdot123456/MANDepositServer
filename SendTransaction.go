package ManDepositServer

import (
	"errors"
	"fmt"
	"github.com/MatrixAINetwork/go-AIMan/manager"
	"github.com/MatrixAINetwork/go-AIMan/transactions"
	"github.com/MatrixAINetwork/go-AIMan/waiting"
	"math/big"
	"time"
	//"testing"
)
var (
	errTimeOut        = errors.New("Time Out")
	errContractRun    = errors.New("Contract Run Error")
	errDeposit        = errors.New("No deposit")
)

func SendDepositTrans_v2(connect *manager.Manager, depAddr string, amount *big.Int, depositType *big.Int, depositRole *big.Int, from string, passphrase string) error {

	var depositFunc func(address string, depositType *big.Int) ([]byte, error)
	err := connect.Unlock(from, passphrase)
	if err != nil {
		return err
	}
	blockNumber, err := connect.Man.GetBlockNumber()
	if err != nil {
		return err
	}
	nonce, err := connect.Man.GetTransactionCount(from, "latest")
	if err != nil {
		return err
	}

	if depositRole.Cmp(big.NewInt(1)) == 0 {
		depositFunc = transactions.MinerDeposit_v2
	} else {
		depositFunc = transactions.ValiDeposit_v2
	}

	data, err := depositFunc(depAddr, depositType)
	if err != nil {
		return err
	}
	trans := transactions.NewTransaction(nonce.Uint64(), transactions.DepositAddr, amount, 100000, big.NewInt(18e9),
		data, 0, 0, 0)

	raw, err := connect.SignTx(trans, from)
	if err != nil {
		return err
	}
	//test1 := common.SendTxArgs1{}
	//err = json.Unmarshal(raw,&test1)
	txID, err := connect.Man.SendRawTransaction(raw)
	if err != nil {
		return err
	}
	fmt.Println("deposit tx hash", txID)
	ReceiptWait := waiting.NewWaitTxReceipt(connect, txID)
	wait3 := waiting.NewMultiWaiting(waiting.NewWaitBlockHeight(connect, blockNumber.Uint64()+10),
		waiting.NewWaitTime(60*time.Second),
		ReceiptWait)
	index := wait3.Waiting()
	if index != 2 {
		return errTimeOut
	}
	if !ReceiptWait.Receipt.Status {
		return errContractRun
	}
	return nil
}

func SendWithDrawTrans_v2(connect *manager.Manager, from string, passphrase string, depositPosition *big.Int, amount *big.Int) error {

	err := connect.Unlock(from, passphrase)
	if err != nil {
		return err
	}
	blockNumber, err := connect.Man.GetBlockNumber()

	if err != nil {
		return err
	}
	nonce, err := connect.Man.GetTransactionCount(from, "latest")
	if err != nil {
		return err
	}
	data, err := transactions.Withdraw_v2(depositPosition, amount)
	if err != nil {
		return err
	}
	trans := transactions.NewTransaction(nonce.Uint64(), transactions.DepositAddr, new(big.Int).SetUint64(0), uint64(400000), big.NewInt(18e9),
		data, 0, 0, 0)
	raw, err := connect.SignTx(trans, from)
	if err != nil {
		return err
	}
	txID, err := connect.Man.SendRawTransaction(raw)
	if err != nil {
		return err
	}
	fmt.Println("withdraw tx hash", txID, "仓位", depositPosition, "金额", amount)
	ReceiptWait := waiting.NewWaitTxReceipt(connect, txID)
	wait3 := waiting.NewMultiWaiting(waiting.NewWaitBlockHeight(connect, blockNumber.Uint64()+10),
		waiting.NewWaitTime(90*time.Second),
		ReceiptWait)
	index := wait3.Waiting()
	if index != 2 {
		return errTimeOut
	}
	if !ReceiptWait.Receipt.Status {
		return errContractRun
	}
	return nil
}

func SendRefundTrans_v2(connect *manager.Manager, from string, passphrase string, depositPosition *big.Int) error {

	err := connect.Unlock(from, passphrase)
	if err != nil {
		return err
	}
	blockNumber, err := connect.Man.GetBlockNumber()

	if err != nil {
		return err
	}

	fmt.Println("blockNumber", blockNumber)
	nonce, err := connect.Man.GetTransactionCount(from, "latest")
	if err != nil {
		return err
	}
	data, err := transactions.Refund_v2(depositPosition)
	if err != nil {
		return err
	}
	trans := transactions.NewTransaction(nonce.Uint64(), transactions.DepositAddr, new(big.Int).SetUint64(0), uint64(400000), big.NewInt(18e9),
		data, 0, 0, 0)
	raw, err := connect.SignTx(trans, from)
	if err != nil {
		return err
	}
	txID, err := connect.Man.SendRawTransaction(raw)
	if err != nil {
		return err
	}
	fmt.Println("refund tx hash", txID, "仓位号", depositPosition)
	ReceiptWait := waiting.NewWaitTxReceipt(connect, txID)
	wait3 := waiting.NewMultiWaiting(waiting.NewWaitBlockHeight(connect, blockNumber.Uint64()+10),
		waiting.NewWaitTime(60*time.Second),
		ReceiptWait)
	index := wait3.Waiting()
	if index != 2 {
		return errTimeOut
	}
	if !ReceiptWait.Receipt.Status {
		return errContractRun
	}
	return nil
}

func SendModifyDepositTrans(connect *manager.Manager, from string, passphrase string, depositType *big.Int, amount *big.Int) error {
	err := connect.Unlock(from, passphrase)
	if err != nil {
		return err
	}
	blockNumber, err := connect.Man.GetBlockNumber()

	if err != nil {
		return err
	}
	nonce, err := connect.Man.GetTransactionCount(from, "latest")
	if err != nil {
		return err
	}
	data, err := transactions.ModifyDepositType_v2(depositType, amount)
	if err != nil {
		return err
	}
	trans := transactions.NewTransaction(nonce.Uint64(), transactions.DepositAddr, new(big.Int).SetUint64(0), uint64(400000), big.NewInt(18e9),
		data, 0, 0, 0)
	raw, err := connect.SignTx(trans, from)
	if err != nil {
		return err
	}
	txID, err := connect.Man.SendRawTransaction(raw)
	if err != nil {
		return err
	}
	fmt.Println("ModifyDepositType tx hash", txID, "抵押类型-活期/定期", depositType, "金额", amount)
	ReceiptWait := waiting.NewWaitTxReceipt(connect, txID)
	wait3 := waiting.NewMultiWaiting(waiting.NewWaitBlockHeight(connect, blockNumber.Uint64()+10),
		waiting.NewWaitTime(60*time.Second),
		ReceiptWait)
	index := wait3.Waiting()
	if index != 2 {
		return errTimeOut
	}
	if !ReceiptWait.Receipt.Status {
		return errContractRun
	}
	return nil
}
