package test

import (
	"testing"

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
}
