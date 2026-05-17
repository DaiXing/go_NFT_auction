package eth

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"my.nft.auction/src/database"
	"my.nft.auction/src/util"
)

// 处理事件。
func handleEventAuctionCreate(log *types.Log, event *abi.Event, eventx *EventAuctionCreate, logx *util.LogMaker) {
	// 去重。
	auctionId := eventx.AuctionId.Uint64()
	logx.AddKV("  auctionId", auctionId)
	if database.ExistsAuctionId(auctionId) {
		logx.AddKV("  提示", "auction已经在DB。不再写入。")
		return
	}

	var row database.AuctionInfoPo
	row.NftContract = eventx.NftContract.Hex()
	row.TokenId = eventx.TokenId.Uint64()
	row.TokenOwner = eventx.TokenOwner.Hex()
	row.AuctionId = eventx.AuctionId.Uint64()
	row.MinPrice = eventx.MinPrice.Uint64()
	row.BeginTime = eventx.BeginTime.Uint64()
	row.PeriodTime = eventx.EndTime.Uint64() - eventx.BeginTime.Uint64()
	row.Bidder = ""
	row.BidPrice = 0
	row.State = util.AUCTION_STATE_NORMAL // 进行中。
	logx.AddKV("  插入 AuctionInfoPo", util.ToJson(&row))

	err := database.Db.Create(&row).Error
	if err == nil {
		logx.AddLine("  插入表 成功")
	} else {
		logx.AddKV("  插入表 错误", err)
	}
}

// 退款
func handleEventAuctionRefund(log *types.Log, event *abi.Event, eventx *EventAuctionRefund, logx *util.LogMaker) {
	auctionId := eventx.AuctionId.Uint64()
	bidId := eventx.BidId.Uint64()
	logx.AddKV("  auctionId", auctionId)
	logx.AddKV("  bidId    ", bidId)

	// 查询。
	_, err := database.QueryBidByBidId(bidId)
	if err != nil {
		logx.AddKV("  查bid 错误 ", err)
		return
	}

	// 更新表。
	err2 := database.UpdateBid(bidId, map[string]any{
		"refund_amount": eventx.Amount.Uint64(),
		"refund_time":   log.BlockTimestamp,
	})

	if err2 == nil {
		logx.AddLine(" 更新bid 成功")
	} else {
		logx.AddKV(" 更新bid 错误", err2)
	}
}

// 竞拍
func handleEventAuctionBid(log *types.Log, event *abi.Event, eventx *EventAuctionBid, logx *util.LogMaker) {
	auctionId := eventx.AuctionId.Uint64()
	bidId := eventx.BidId.Uint64()
	logx.AddKV("  auctionId", auctionId)
	logx.AddKV("  bidId    ", bidId)

	// 查询。
	bidInfo, err := database.QueryBidByBidId(bidId)
	if err == nil {
		logx.AddKV("  提示", "bid 已经存在。不写表。"+util.ToJson(&bidInfo))
		return
	}

	// 写表。
	var row database.AuctionBidPo
	row.AuctionId = auctionId
	row.BidId = bidId
	row.Bidder = eventx.Bidder.Hex()
	row.BidPrice = eventx.BidPrice.Uint64()
	row.BidTime = log.BlockTimestamp // 时间戳。
	row.RefundAmount = 0
	row.RefundTime = 0
	logx.AddKV(" 准备写row", util.ToJson(&row))

	// 插入。
	err3 := database.Db.Create(&row).Error
	if err3 != nil {
		logx.AddKV(" 写bid 错误", err3)
		return
	} else {
		logx.AddLine(" 写bid 成功")
	}

	// 更新 auction
	err4 := database.UpdateAuction(auctionId, map[string]any{
		"bidder":    eventx.Bidder.Hex(),
		"bid_price": eventx.BidPrice.Uint64(),
	})
	if err4 == nil {
		logx.AddLine(" 更新auction 成功")
	} else {
		logx.AddKV(" 更新auction 错误", err4)
	}
}

// 取消。
func handleEventAuctionCancel(log *types.Log, event *abi.Event, eventx *EventAuctionCancel, logx *util.LogMaker) {
	auctionId := eventx.AuctionId.Uint64()
	logx.AddKV("  auctionId", auctionId)

	auction, err := database.QueryActionInfoByAuctionId(auctionId)
	if err == nil {
		// 判断状态。
		if auction.State == util.AUCTION_STATE_NORMAL {
			err2 := database.UpdateAuction(auctionId, map[string]any{
				"state": util.AUCTION_STATE_CANCEL,
			})
			if err2 == nil {
				logx.AddKV(" 提示", "更新表 成功")
			} else {
				logx.AddKV(" 更新表 错误", err2)
			}
		} else {
			logx.AddKV(" 提示", "状态不是进行中。无法取消。")
		}
	} else {
		logx.AddKV(" 查表 错误", err)
	}
}

// 结束。
func handleEventAuctionEnd(log *types.Log, event *abi.Event, eventx *EventAuctionEnd, logx *util.LogMaker) {
	auctionId := eventx.AuctionId.Uint64()
	logx.AddKV("  auctionId", auctionId)

	err := database.UpdateAuction(auctionId, map[string]any{
		"state": eventx.State,
	})
	if err == nil {
		logx.AddKV(" 提示", "更新 auction表成功。")
	} else {
		logx.AddKV(" 更新报错", err)
	}
}
