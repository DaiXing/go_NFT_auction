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

	err := database.Db.Create(&row).Error
	util.CheckError(err)

	logx.AddKV("  插入 AuctionInfoPo", util.ToJson(row))
}
func handleEventAuctionRefund(log *types.Log, event *abi.Event, eventx *EventAuctionRefund, logx *util.LogMaker) {

}
func handleEventAuctionBid(log *types.Log, event *abi.Event, eventx *EventAuctionBid, logx *util.LogMaker) {

}
func handleEventAuctionCancel(log *types.Log, event *abi.Event, eventx *EventAuctionCancel, logx *util.LogMaker) {

}
func handleEventAuctionEnd(log *types.Log, event *abi.Event, eventx *EventAuctionEnd, logx *util.LogMaker) {

}
