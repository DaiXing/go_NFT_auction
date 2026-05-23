package mock

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"my.nft.auction/src/eth"
	"my.nft.auction/src/util"
)

// NFT
var nftABI *abi.ABI
var nftContract common.Address

// 用户
var userJack *eth.UserInfo
var userTom *eth.UserInfo

func InitMock() {
	util.Logger.Info("mock 开始初始化")
	mock := &util.Params.Mock

	// 合约。
	nftContract = common.HexToAddress(mock.NftContractAddr)

	// ABI
	abix, err := eth.AbiFromFile("ABI_my_NFT.json")
	util.CheckError(err)
	nftABI = abix
	util.Logger.Info("mock ", "nftABI", nftABI)

	// 用户
	userJack = eth.UserFromPrivateKey(mock.JackPrivateKey)
	userTom = eth.UserFromPrivateKey(mock.TomPrivateKey)

	util.Logger.Info("mock ", "userJack", util.ToJson(userJack))
	util.Logger.Info("mock ", "userTom", util.ToJson(userTom))
	util.Logger.Info("mock 完成初始化")
}
