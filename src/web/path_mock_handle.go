package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"my.nft.auction/src/database"
	"my.nft.auction/src/mock"
	"my.nft.auction/src/util"
)

// NFT 创建token
func pathMockNftMint(ctx *gin.Context) {
	// 创建token
	caller := mock.UserJack
	mock.TxNftMint(caller)

	ctx.JSON(http.StatusOK, BaseResp{
		Message: "OK",
	})
}

// 查询 token列表。
func pathMockGetTokenList(ctx *gin.Context) {
	// 参数。
	var req GetTokenListReq
	err := ctx.ShouldBindJSON(&req)
	util.CheckError(err)

	// 查 NFT
	var tokenList []database.MockTokenPo
	err2 := database.Db.Where("owner = ?", req.Seller).Find(&tokenList).Error
	util.CheckError(err2)

	// 返回。
	var resp GetTokenListResp

	// token
	for _, token := range tokenList {
		tokenInfo := TokenInfo{
			NftContract: token.NftContract,
			TokenId:     token.TokenId,
			TokenUri:    "no_uri",
			TokenType:   "ERC721",
			FloorPrice:  3.2,
			Description: "no_desc",
			Image1:      "no_image",
			Image2:      "no_image",
			Image3:      "no_image",
		}

		// 查拍卖。
		auctionInfo, _ := database.QueryActionInfoBySeller(req.Seller, tokenInfo.NftContract, tokenInfo.TokenId)
		tokenInfo.AuctionInfo = auctionInfo

		resp.TokenList = append(resp.TokenList, &tokenInfo)
	}

	webReturnOKJson(ctx, &resp)
}
