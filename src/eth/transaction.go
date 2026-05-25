package eth

import (
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"my.nft.auction/src/util"
)

// 交易数据。
func NewTxData(from common.Address, toContract common.Address,
	funcData []byte, value *big.Int) *types.DynamicFeeTx {
	ctx, cancel := util.NewContext(5)
	defer cancel()

	chainId, err0 := EthClient.ChainID(ctx)
	util.CheckError(err0)

	// blockNum, err1 := EthClient.BlockNumber(ctx)
	// util.CheckError(err1)

	// gas价格
	gasPrice, err2 := EthClient.SuggestGasPrice(ctx)
	util.CheckError(err2)

	// gas小费
	gasTip, err3 := EthClient.SuggestGasTipCap(ctx)
	util.CheckError(err3)

	// gas数量
	callMsg := ethereum.CallMsg{
		From:  from,        // 调用方。
		To:    &toContract, // 合约地址
		Data:  funcData,    // 合约函数
		Value: value,       // 金额。 也是必须的。
	}
	gasLimit, err4 := EthClient.EstimateGas(ctx, callMsg)
	util.CheckError(err4)
	gasLimit2 := big.NewInt(int64(gasLimit))

	// gas 总费用。
	gasFee := big.NewInt(0).Mul(big.NewInt(0).Add(gasPrice, gasTip), gasLimit2)

	// 序号。
	nonce, err5 := EthClient.PendingNonceAt(ctx, from)
	util.CheckError(err5)

	// 交易数据。
	txData := types.DynamicFeeTx{
		ChainID:   chainId,
		Nonce:     nonce,
		GasTipCap: gasTip,
		Gas:       gasLimit,
		GasFeeCap: gasFee,
		Data:      funcData,
		Value:     value,
		To:        &toContract,
	}
	return &txData
}

// 触发交易。
func CallTx(title string, contract common.Address, funcData []byte, caller *UserInfo) {
	CallTx2(title, contract, funcData, caller, nil)
}

// 触发交易。
func CallTx2(title string, contract common.Address, funcData []byte, caller *UserInfo, value *big.Int) {
	logMaker := util.LogMaker{}
	defer logMaker.LogString()
	logMaker.AddLine(">> 发送一个交易: " + title)
	logMaker.AddKV(" caller", caller.Username)
	logMaker.AddKV(" contract", contract)

	// 交易。
	txData := NewTxData(caller.Addr, contract, funcData, value)
	logMaker.AddKV(" txData.Value", txData.Value)

	// 签名
	signer := types.NewLondonSigner(txData.ChainID)
	tx, err2 := types.SignNewTx(caller.PrivateKey, signer, txData)
	util.CheckError(err2)
	txHash := tx.Hash().Hex()
	logMaker.AddKV(" txHash", txHash)

	ctx, cancel := util.NewContext(5)
	defer cancel()

	// 发送交易。
	err5 := EthClient.SendTransaction(ctx, tx)
	util.CheckError(err5)
	logMaker.AddKV(" 提示", "发生完成")
}
