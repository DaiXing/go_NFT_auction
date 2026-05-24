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
var UserJack *eth.UserInfo
var UserTom *eth.UserInfo
var UserBobo *eth.UserInfo
var UserMap = map[string]*eth.UserInfo{}

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
	UserJack = eth.UserFromPrivateKey(mock.JackPrivateKey)
	UserTom = eth.UserFromPrivateKey(mock.TomPrivateKey)
	UserBobo = eth.UserFromPrivateKey(mock.BoboPrivateKey)
	UserJack.Username = "jack"
	UserTom.Username = "tom"
	UserBobo.Username = "bobo"
	UserMap[UserJack.Username] = UserJack
	UserMap[UserTom.Username] = UserTom
	UserMap[UserBobo.Username] = UserBobo

	util.Logger.Info("mock ", "userJack", util.ToJson(UserJack))
	util.Logger.Info("mock ", "userTom", util.ToJson(UserTom))
	util.Logger.Info("mock ", "userBobo", util.ToJson(UserBobo))

	// 事件。
	NftSubscribeEvent()

	// 授权
	InitNftApprove()

	util.Logger.Info("mock 完成初始化")
}
