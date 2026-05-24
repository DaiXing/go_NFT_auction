package mock

import (
	"my.nft.auction/src/eth"
	"my.nft.auction/src/util"
)

// 铸造一个token
func TxNftMint(caller *eth.UserInfo) {
	tokenUri := "http://aa.com/token.json"

	// 函数+入参
	funcData, err := nftABI.Pack("mintToken", tokenUri)
	util.CheckError(err)

	eth.CallTx(nftContract, funcData, caller)
}
