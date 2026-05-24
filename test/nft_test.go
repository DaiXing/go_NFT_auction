package test

import (
	"testing"
	"time"

	"my.nft.auction/src/util"
	"my.nft.auction/src/web"
)

func TestMockNftMint(tt *testing.T) {
	tt.Log("测试： mock NFT 发币")

	resp, err := util.HttpPostJson[web.BaseReq, web.BaseResp](URL_MOCK_NFT_MINT, nil)
	if err != nil {
		tt.Fatal("URL_MOCK_NFT_MINT error", err)
	}
	tt.Log("返回 ", resp.Message)

	// 等交易执行完成。
	time.Sleep(time.Second * 2)

	// 查token
	req2 := web.GetTokenListReq{
		Seller: jackAddr,
		BaseReq: web.BaseReq{
			PageSize: 3,
		},
	}
	resp2, err2 := util.HttpPostJson[web.GetTokenListReq, web.GetTokenListResp](URL_MOCK_GET_TOKEN_LIST, &req2)
	if err2 != nil {
		tt.Fatal("URL_MOCK_GET_TOKEN_LIST error", err2)
	}
	tokenLen := len(resp2.TokenList)
	tt.Log("token数量 ", tokenLen)
	if tokenLen == 0 {
		tt.Errorf("没有查到token")
	}

}
