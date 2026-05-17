package eth

import (
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"my.nft.auction/src/util"
)

// 订阅事件。
func SubscribeEvent() {
	ctx, cancel := util.NewContext(3)
	defer cancel()

	// 订阅日志。
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddr},
	}
	chanEvent := make(chan types.Log)
	sub, err2 := Client.SubscribeFilterLogs(ctx, query, chanEvent)
	util.CheckError(err2)

	go func() {
		for {
			select {
			case <-sub.Err():
				util.Logger.Error("订阅事件发生错误", "err", sub.Err())
				return
			case logE := <-chanEvent: // 解析日志
				parseEvent(logE)
			}
		}
	}()
}

// 解析log
func parseEvent(log *types.Log) {
	// 找出具体的事件定义
	topic0 := log.Topics[0]
	event, err1 := abiObj.EventByID(topic0)
	util.CheckError(err1)

	eventName := event.Name

	// 区分事件
	if "AuctionCreate" == eventName {
		// 解析事件。
		var eventx EventAuctionCreate
		err2 := abiObj.UnpackIntoInterface(&eventx, eventName, log.Data)
		util.CheckError(err2)

		eventx.NftContract = common.BytesToAddress(log.Topics[1].Bytes())
		eventx.TokenId = big.NewInt(0).SetBytes(log.Topics[2].Bytes())
		eventx.AuctionId = big.NewInt(0).SetBytes(log.Topics[3].Bytes())

		// 处理事件。
		handleEventAuctionCreate(log, event, &eventx)
	} else if "AuctionRefund" == eventName {
		// 解析事件。
		var eventx EventAuctionRefund
		err2 := abiObj.UnpackIntoInterface(&eventx, eventName, log.Data)
		util.CheckError(err2)

		eventx.AuctionId = big.NewInt(0).SetBytes(log.Topics[1].Bytes())
		eventx.To = common.BytesToAddress(log.Topics[2].Bytes())

		// 处理事件。
		handleEventAuctionRefund(log, event, &eventx)
	} else if "AuctionBid" == eventName {
		// 解析事件。
		var eventx EventAuctionBid
		err2 := abiObj.UnpackIntoInterface(&eventx, eventName, log.Data)
		util.CheckError(err2)

		eventx.AuctionId = big.NewInt(0).SetBytes(log.Topics[1].Bytes())
		eventx.Bidder = common.BytesToAddress(log.Topics[2].Bytes())

		// 处理事件。
		handleEventAuctionBid(log, event, &eventx)
	} else if "AuctionCancel" == eventName {
		// 解析事件。
		var eventx EventAuctionCancel
		err2 := abiObj.UnpackIntoInterface(&eventx, eventName, log.Data)
		util.CheckError(err2)

		eventx.AuctionId = big.NewInt(0).SetBytes(log.Topics[1].Bytes())

		// 处理事件。
		handleEventAuctionCancel(log, event, &eventx)
	} else if "AuctionEnd" == eventName {
		// 解析事件。
		var eventx EventAuctionEnd
		err2 := abiObj.UnpackIntoInterface(&eventx, eventName, log.Data)
		util.CheckError(err2)

		eventx.AuctionId = big.NewInt(0).SetBytes(log.Topics[1].Bytes())

		// 处理事件。
		handleEventAuctionEnd(log, event, &eventx)
	} else {
		// 未知事件。
	}

}
