package test

import (
	"testing"
	"time"

	"my.nft.auction/src/bean"
	"my.nft.auction/src/util"
)

func TestMock(tt *testing.T) {
	tt.Log("测试： mock NFT发币，拍卖，出价")

	// NFT发币
	resp, err := util.HttpPostJson[bean.BaseReq, bean.BaseResp](URL_MOCK_NFT_MINT, nil)
	if err != nil {
		tt.Fatal("NFT发币 error", err)
	}
	tt.Log("返回 ", resp.Message)

	// 等交易执行完成。
	time.Sleep(time.Second * 2)

	// 查token
	req2 := bean.GetTokenListReq{
		Seller: jackAddr,
		BaseReq: bean.BaseReq{
			PageSize: 3,
		},
	}
	resp2, err2 := util.HttpPostJson[bean.GetTokenListReq, bean.GetTokenListResp](URL_MOCK_GET_TOKEN_LIST, &req2)
	if err2 != nil {
		tt.Fatal("查token error", err2)
	}
	tokenLen := len(resp2.TokenList)
	tt.Log("token数量 ", tokenLen)
	if tokenLen == 0 {
		tt.Errorf("没有查到token")
	}

	// 用1个token。
	tokenA := resp2.TokenList[0]
	nowSecond := time.Now().Unix()

	// 创建 拍卖
	req3 := bean.CreateAuctionReq{
		NftContract: tokenA.NftContract,
		TokenId:     tokenA.TokenId,
		MinPrice:    int64(100),
		BeginTime:   nowSecond,
		PeriodTime:  int64(60 * 6),
	}
	_, err3 := util.HttpPostJson[bean.CreateAuctionReq, bean.BaseResp](URL_MOCK_CREATE_AUCTION, &req3)
	if err3 != nil {
		tt.Fatal("创建 拍卖 error", err3)
	}

	// 查询最新的拍卖。
	req4 := bean.GetAuctionListReq{
		Seller:      jackAddr,
		NftContract: tokenA.NftContract,
		BaseReq: bean.BaseReq{
			PageSize: 100,
		},
	}
	resp4, err4 := util.HttpPostJson[bean.GetAuctionListReq, bean.GetAuctionListResp](URL_GET_AUCTION_LIST, &req4)
	if err4 != nil {
		tt.Fatal("查询 拍卖 error", err4)
	}

	// 一个拍卖。
	auction := resp4.AuctionList[0]

	// 出价。
	req5 := bean.BidAuctionReq{
		AuctionId:  auction.AuctionId,
		BidderName: "bobo",
		BidPrice:   105,
	}
	_, err5 := util.HttpPostJson[bean.BidAuctionReq, bean.BaseResp](URL_MOCK_BID_AUCTION, &req5)
	if err5 != nil {
		tt.Fatal("拍卖 出价 error", err5)
	}

	// 出价。
	req6 := bean.BidAuctionReq{
		AuctionId:  auction.AuctionId,
		BidderName: "tom",
		BidPrice:   108,
	}
	_, err6 := util.HttpPostJson[bean.BidAuctionReq, bean.BaseResp](URL_MOCK_BID_AUCTION, &req6)
	if err6 != nil {
		tt.Fatal("拍卖 出价 error", err6)
	}
	// 出价。
	req7 := bean.BidAuctionReq{
		AuctionId:  auction.AuctionId,
		BidderName: "bobo",
		BidPrice:   125,
	}
	_, err7 := util.HttpPostJson[bean.BidAuctionReq, bean.BaseResp](URL_MOCK_BID_AUCTION, &req7)
	if err7 != nil {
		tt.Fatal("拍卖 出价 error", err7)
	}

	time.Sleep(time.Second * 2)

	// 查询出价列表。
	req8 := bean.GetBidListReq{
		AuctionId: auction.AuctionId,
	}
	_, err8 := util.HttpPostJson[bean.GetBidListReq, bean.GetBidListResp](URL_GET_BID_LIST, &req8)
	if err8 != nil {
		tt.Fatal("查询出价列表 error", err8)
	}

}
