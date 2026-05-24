package mock

import (
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"my.nft.auction/src/bean"
	"my.nft.auction/src/eth"
	"my.nft.auction/src/util"
)

// 创建拍卖
func TxCreateAuction(caller *eth.UserInfo, req *bean.CreateAuctionReq) {
	logMaker := util.LogMaker{}
	defer logMaker.LogString()
	logMaker.AddLine(">> path 创建拍卖")

	// 合约
	nftAddr := common.HexToAddress(req.NftContract)

	// 转数字。
	tokenIdInt, err1 := strconv.ParseInt(req.TokenId, 10, 64)
	util.CheckError(err1)
	tokenIdInt2 := big.NewInt(tokenIdInt)

	minPrice := big.NewInt(int64(req.MinPrice))
	beginTime := big.NewInt(int64(req.BeginTime))
	periodTime := big.NewInt(int64(req.PeriodTime))

	logMaker.AddKV(" nftAddr ", nftAddr)
	logMaker.AddKV(" tokenId ", tokenIdInt2)
	logMaker.AddKV(" minPrice ", minPrice)
	logMaker.AddKV(" beginTime ", beginTime)
	logMaker.AddKV(" periodTime ", periodTime)

	// 函数+入参
	funcData, err := eth.AuctionABI.Pack("createAuction", nftAddr, tokenIdInt2,
		minPrice, beginTime, periodTime)
	logMaker.AddKV(" 创建 funcData error ", err)
	util.CheckError(err)

	// 交易。
	eth.CallTx(eth.AuctionContractAddr, funcData, caller)
	logMaker.AddLine(" 创建 交易 ")
}

// 出价
func TxBidAuction(caller *eth.UserInfo, req *bean.BidAuctionReq) {
	logMaker := util.LogMaker{}
	defer logMaker.LogString()
	logMaker.AddLine(">> path 拍卖出价")
	logMaker.AddKV(" caller ", caller.Username)

	auctionId, _ := big.NewInt(0).SetString(req.AuctionId, 10)
	bidPrice := big.NewInt(0).SetUint64(uint64(req.BidPrice))
	// bidPrice.SetUint64(70000000)

	logMaker.AddKV(" auctionId ", auctionId)
	logMaker.AddKV(" bidPrice ", bidPrice)

	// 函数+入参
	funcData, err := eth.AuctionABI.Pack("bidAuction", auctionId)
	logMaker.AddKV(" 创建 funcData error ", err)
	util.CheckError(err)

	// 交易。
	eth.CallTx2(eth.AuctionContractAddr, funcData, caller, bidPrice)
	logMaker.AddLine(" 创建 交易 ")
}
