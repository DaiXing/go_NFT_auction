package web

import (
	"time"

	"github.com/gin-gonic/gin"
	"my.nft.auction/src/database"
	"my.nft.auction/src/util"
)

// 健康检测。
func handleHealth(ctx *gin.Context) {
	resp := BaseResp{
		Message: "NFT auction : " + time.Now().Format(time.DateTime),
	}
	webReturnOKJson(ctx, resp)
}

// 查询 token列表。
func handleGetTokenList(ctx *gin.Context) {
	// 参数。
	var req GetTokenListReq
	err := ctx.ShouldBindJSON(&req)
	util.CheckError(err)

	// 返回。
	var resp GetTokenListResp
	webReturnOKJson(ctx, &resp)
}

// 查询 拍卖列表。
func handleGetAuctionList(ctx *gin.Context) {
	// 参数。
	var req GetAuctionListReq
	err := ctx.ShouldBindJSON(&req)
	util.CheckError(err)

	// 查询。
	tx := database.Db.Model(&database.AuctionInfoPo{})
	if len(req.NftContract) > 0 {
		tx = tx.Where("nft_contract = ?", req.NftContract)
	}
	if req.TokenId > 0 {
		tx = tx.Where("token_id = ?", req.TokenId)
	}
	if len(req.TokenOwner) > 0 {
		tx = tx.Where("token_owner = ?", req.TokenOwner)
	}
	if req.AuctionId > 0 {
		tx = tx.Where("auction_id = ?", req.AuctionId)
	}

	// 查数量。
	var count int64
	err2 := tx.Count(&count).Error
	util.CheckError(err2)

	// 分页。
	offset := (req.PageNo - 1) * req.PageSize
	limit := req.PageSize
	tx.Offset(offset).Limit(limit)

	// 查列表。
	var auctions []database.AuctionInfoPo
	err3 := tx.Find(&auctions).Error
	util.CheckError(err3)

	// 返回。
	var resp GetAuctionListResp
	resp.TotalSize = count
	resp.AuctionList = auctions
	webReturnOKJson(ctx, &resp)
}

// 查询 出价列表。
func handleGetBidList(ctx *gin.Context) {
	// 参数。
	var req GetBidListReq
	err := ctx.ShouldBindJSON(&req)
	util.CheckError(err)

	if req.AuctionId == 0 {
		webAbortBadReq(ctx, "AuctionId is required")
	}

	// 查询。
	tx := database.Db.Model(&database.AuctionBidPo{})
	tx = tx.Where("auction_id = ?", req.AuctionId)

	// 查数量。
	var count int64
	err2 := tx.Count(&count).Error
	util.CheckError(err2)

	// 分页。
	offset := (req.PageNo - 1) * req.PageSize
	limit := req.PageSize
	tx.Offset(offset).Limit(limit)

	// 查列表。
	var bids []database.AuctionBidPo
	err3 := tx.Find(&bids).Error
	util.CheckError(err3)

	// 返回。
	var resp GetBidListResp
	resp.TotalSize = count
	resp.BidList = bids
	webReturnOKJson(ctx, &resp)
}
