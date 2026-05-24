package mock

import (
	"math/big"

	"my.nft.auction/src/bean"
	"my.nft.auction/src/eth"
	"my.nft.auction/src/util"
)

// 创建拍卖
func TxCreateAuction(caller *eth.UserInfo, req *bean.CreateAuctionReq) {
	// 函数+入参
	funcData, err := eth.AuctionABI.Pack("createAuction", req.NftContract, req.TokenId,
		big.NewInt(int64(req.MinPrice)), big.NewInt(int64(req.BeginTime)),
		big.NewInt(int64(req.PeriodTime)))
	util.CheckError(err)

	eth.CallTx(eth.AuctionContractAddr, funcData, caller)
}
