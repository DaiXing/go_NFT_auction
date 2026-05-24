package test

import (
	"testing"

	"my.nft.auction/src/bean"
	"my.nft.auction/src/util"
)

func TestHealth(tt *testing.T) {
	tt.Log("测试：健康检测")
	resp, err := util.HttpGetJson[bean.BaseResp](URL_HEALTH)
	if err != nil {
		tt.Fatal("URL_HEALTH error", err)
	}

	tt.Log("健康检测 ", resp.Message)
}

func TestQueryToken(tt *testing.T) {
	tt.Log("测试：查询token列表")
	req := bean.GetTokenListReq{
		Seller: "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
		BaseReq: bean.BaseReq{
			PageSize: 2,
		},
	}
	resp, err := util.HttpPostJson[bean.GetTokenListReq, bean.GetTokenListResp](URL_GET_TOKEN_LIST, &req)
	if err != nil {
		tt.Fatal("URL_GET_TOKEN_LIST error", err)
	}
	tt.Log(" token数量 = ", len(resp.TokenList))
}

func TestSTATISTIC(tt *testing.T) {
	tt.Log("测试： 查询统计")
	resp, err := util.HttpGetJson[bean.StatisticResp](URL_GLOBAL_STATISTIC)
	if err != nil {
		tt.Fatal("查询统计 error", err)
	}

	tt.Log("查询统计 ", util.ToJson(resp))
}
