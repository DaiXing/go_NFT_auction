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
	nowSecond := time.Now().Second()

	// 创建 拍卖
	req3 := bean.CreateAuctionReq{
		NftContract: tokenA.NftContract,
		TokenId:     tokenA.TokenId,
		MinPrice:    int64(100),
		BeginTime:   int64(nowSecond),
		PeriodTime:  int64(60),
	}
	_, err3 := util.HttpPostJson[bean.CreateAuctionReq, bean.BaseResp](URL_MOCK_CREATE_AUCTION, &req3)
	if err3 != nil {
		tt.Fatal("创建 拍卖 error", err3)
	}

}
