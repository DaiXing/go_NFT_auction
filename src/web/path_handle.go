package web

import (
	"time"

	"github.com/gin-gonic/gin"
	"my.nft.auction/src/database"
	"my.nft.auction/src/eth"
	"my.nft.auction/src/util"
)

// 健康检测。
func pathHealth(ctx *gin.Context) {
	resp := BaseResp{
		Message: "NFT auction : " + time.Now().Format(time.DateTime),
	}
	webReturnOKJson(ctx, resp)
}

// 查询 token列表。
func pathGetTokenList(ctx *gin.Context) {
	// 参数。
	var req GetTokenListReq
	err := ctx.ShouldBindJSON(&req)
	util.CheckError(err)

	// 查 NFT
	nftResp, err2 := eth.AlchemyQueryNft(req.Seller, req.PageSize, true)
	util.CheckError(err2)

	// 返回。
	var resp GetTokenListResp
	webReturnOKJson(ctx, &resp)
}

// 查询 拍卖列表。
func pathGetAuctionList(ctx *gin.Context) {
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
	if len(req.Seller) > 0 {
		tx = tx.Where("seller = ?", req.Seller)
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
func pathGetBidList(ctx *gin.Context) {
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

// 统计。
func pathStatistic(ctx *gin.Context) {
	var resp StatisticResp

	// 拍卖总数
	err1 := database.Db.Model(&database.AuctionInfoPo{}).Count(&resp.CountAuction).Error
	util.CheckError(err1)

	// 出价总数
	err2 := database.Db.Model(&database.AuctionBidPo{}).Count(&resp.CountBid).Error
	util.CheckError(err2)

	// 总锁仓价值
	var sumBidPrice int64
	err3 := database.Db.Model(&database.AuctionBidPo{}).Select("sum(bid_price)").Scan(&sumBidPrice).Error
	util.CheckError(err3)

	resp.SumTvl += sumBidPrice

	webReturnOKJson(ctx, resp)
}
