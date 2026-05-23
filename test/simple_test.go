package test

import (
	"testing"

	"my.nft.auction/src/util"
	"my.nft.auction/src/web"
)

func TestHealth(tt *testing.T) {
	tt.Log("测试：健康检测")
	resp, err := util.HttpGetJson[web.BaseResp](URL_HEALTH)
	if err != nil {
		tt.Fatal("URL_HEALTH error", err)
	}

	tt.Log("健康检测 ", resp.Message)
}

func TestQueryToken(tt *testing.T) {
	tt.Log("测试：查询token列表")
	req := web.GetTokenListReq{
		Seller: "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
		BaseReq: web.BaseReq{
			PageSize: 2,
		},
	}
	resp, err := util.HttpPostJson[web.GetTokenListReq, web.GetTokenListResp](URL_GET_TOKEN_LIST, &req)
	if err != nil {
		tt.Fatal("URL_GET_TOKEN_LIST error", err)
	}
	tt.Log(" token数量 = ", len(resp.TokenList))
}
