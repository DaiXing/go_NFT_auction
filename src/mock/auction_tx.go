package mock

import (
	"my.nft.auction/src/eth"
	"my.nft.auction/src/util"
)

// 创建拍卖
func TxCreateAuction(caller *eth.UserInfo) {
	tokenUri := "http://aa.com/token.json"

	// 函数+入参
	funcData, err := eth.AuctionABI.Pack("createAuction", tokenUri)
	util.CheckError(err)

	eth.CallTx(eth.AuctionContractAddr, funcData, caller)
}
