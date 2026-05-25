package mock

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"my.nft.auction/src/eth"
	"my.nft.auction/src/util"
)

// NFT铸造一个token
func TxNftMint(caller *eth.UserInfo) {
	logMaker := util.LogMaker{}
	defer logMaker.LogString()
	logMaker.AddLine("NFT铸造一个token")

	tokenUri := fmt.Sprintf("http://aa.com/%d.json", time.Now().UnixMilli())
	logMaker.AddKV("caller", caller.Username)
	logMaker.AddKV("tokenUri", tokenUri)

	// 函数+入参
	funcData, err := nftABI.Pack("mintToken", tokenUri)
	util.CheckError(err)

	eth.CallTx("NFT铸造", nftContract, funcData, caller)
}

// NFT授权。简化。
func TxNftApproveAll(caller *eth.UserInfo, toAddr common.Address) {
	// 函数+入参
	funcData, err := nftABI.Pack("setApprovalForAll", toAddr, true)
	util.CheckError(err)

	eth.CallTx("NFT授权", nftContract, funcData, caller)
}
