package mock

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"my.nft.auction/src/bean"
	"my.nft.auction/src/database"
	"my.nft.auction/src/util"
)

// 设置路径。
func SetupMockPath(webServer *gin.Engine) {
	group1 := webServer.Group("/mock")
	group1.POST("/nft-mint", pathMockNftMint)
	group1.POST("/get-token-list", pathMockGetTokenList)
	group1.POST("/create-auction", pathMockTxCreateAuction)
	group1.POST("/bid-auction", pathMockBidAuction)
}

// NFT 创建token
func pathMockNftMint(ctx *gin.Context) {
	// 创建token
	caller := UserJack
	TxNftMint(caller)

	ctx.JSON(http.StatusOK, bean.BaseResp{
		Message: "OK",
	})
}

// 查询 token列表。
func pathMockGetTokenList(ctx *gin.Context) {
	// 参数。
	var req bean.GetTokenListReq
	err := ctx.ShouldBindJSON(&req)
	util.CheckError(err)

	// 查 NFT
	var tokenList []database.MockTokenPo
	err2 := database.Db.Where("owner = ?", req.Seller).Find(&tokenList).Error
	util.CheckError(err2)

	// 返回。
	var resp bean.GetTokenListResp

	// token
	for _, token := range tokenList {
		tokenInfo := bean.TokenInfo{
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

	ctx.JSON(http.StatusOK, &resp)
}

// 创建拍卖
func pathMockTxCreateAuction(ctx *gin.Context) {
	// 参数。
	var req bean.CreateAuctionReq
	err := ctx.ShouldBindJSON(&req)
	util.CheckError(err)

	// 创建token
	caller := UserJack
	TxCreateAuction(caller, &req)

	ctx.JSON(http.StatusOK, bean.BaseResp{
		Message: "OK",
	})
}

// 出价
func pathMockBidAuction(ctx *gin.Context) {
	// 参数。
	var req bean.BidAuctionReq
	err := ctx.ShouldBindJSON(&req)
	util.CheckError(err)

	// 创建token
	caller := UserMap[req.CallerName]
	TxBidAuction(caller, &req)

	ctx.JSON(http.StatusOK, bean.BaseResp{
		Message: "OK",
	})
}
