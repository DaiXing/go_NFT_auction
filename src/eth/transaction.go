package eth

import (
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"my.nft.auction/src/util"
)

// 交易数据。
func NewTxData(from common.Address, toContract common.Address, funcData []byte) *types.DynamicFeeTx {
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
		From: from,
		To:   &toContract, // 合约地址
		Data: funcData,    // 合约函数
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
		Value:     nil,
		To:        &toContract,
	}
	return &txData
}
