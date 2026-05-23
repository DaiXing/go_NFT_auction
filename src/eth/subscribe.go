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
	chanEvent := make(chan types.Log) // 队列。
	sub, err2 := EthClientWS.SubscribeFilterLogs(ctx, query, chanEvent)
	util.CheckError(err2)

	go func() {
		for {
			select {
			case <-sub.Err():
				util.Logger.Error("订阅事件 发生错误", "err", sub.Err())

				sub.Unsubscribe() // 取消订阅。

				SubscribeEvent() // 重新订阅。

				return
			case logE := <-chanEvent: // 解析日志

				parseOneEvent(&logE) // 处理日志。

				updateScanBlockNum(logE.BlockNumber) // 更新DB。
			}
		}
	}()
}

// 解析log
func parseOneEvent(log *types.Log) {
	logx := util.LogMaker{}
	logx.AddLine(">> 解析一个event ")
	defer logx.LogString()

	defer func() {
		err := recover() // 捕获解析事件的panic，避免程序崩溃。
		logx.AddKV(" 异常", err)
	}()

	// 找出具体的事件定义
	topic0 := log.Topics[0]
	logx.AddKV("  topic0", topic0)
	event, err1 := abiObj.EventByID(topic0)
	util.CheckError(err1)

	eventName := event.Name
	logx.AddKV("  eventName", eventName)

	// 区分事件
	if "AuctionCreate" == eventName {
		// 解析事件。
		var eventx EventAuctionCreate
		err2 := abiObj.UnpackIntoInterface(&eventx, eventName, log.Data)
		util.CheckError(err2)

		eventx.Seller = common.BytesToAddress(log.Topics[1].Bytes())
		eventx.TokenId = big.NewInt(0).SetBytes(log.Topics[2].Bytes())
		eventx.AuctionId = big.NewInt(0).SetBytes(log.Topics[3].Bytes())

		logx.AddKV("  event明细", util.ToJson(eventx))

		// 处理事件。
		handleEventAuctionCreate(log, event, &eventx, &logx)
	} else if "AuctionRefund" == eventName {
		// 解析事件。
		var eventx EventAuctionRefund
		err2 := abiObj.UnpackIntoInterface(&eventx, eventName, log.Data)
		util.CheckError(err2)

		eventx.AuctionId = big.NewInt(0).SetBytes(log.Topics[1].Bytes())
		eventx.To = common.BytesToAddress(log.Topics[2].Bytes())

		logx.AddKV("  event明细", util.ToJson(eventx))

		// 处理事件。
		handleEventAuctionRefund(log, event, &eventx, &logx)
	} else if "AuctionBid" == eventName {
		// 解析事件。
		var eventx EventAuctionBid
		err2 := abiObj.UnpackIntoInterface(&eventx, eventName, log.Data)
		util.CheckError(err2)

		eventx.AuctionId = big.NewInt(0).SetBytes(log.Topics[1].Bytes())
		eventx.Bidder = common.BytesToAddress(log.Topics[2].Bytes())

		logx.AddKV("  event明细", util.ToJson(eventx))

		// 处理事件。
		handleEventAuctionBid(log, event, &eventx, &logx)
	} else if "AuctionCancel" == eventName {
		// 解析事件。
		var eventx EventAuctionCancel
		err2 := abiObj.UnpackIntoInterface(&eventx, eventName, log.Data)
		util.CheckError(err2)

		eventx.AuctionId = big.NewInt(0).SetBytes(log.Topics[1].Bytes())

		logx.AddKV("  event明细", util.ToJson(eventx))

		// 处理事件。
		handleEventAuctionCancel(log, event, &eventx, &logx)
	} else if "AuctionEnd" == eventName {
		// 解析事件。
		var eventx EventAuctionEnd
		err2 := abiObj.UnpackIntoInterface(&eventx, eventName, log.Data)
		util.CheckError(err2)

		eventx.AuctionId = big.NewInt(0).SetBytes(log.Topics[1].Bytes())

		logx.AddKV("  event明细", util.ToJson(eventx))

		// 处理事件。
		handleEventAuctionEnd(log, event, &eventx, &logx)
	} else {
		// 未知事件。
		logx.AddKV("  忽略", "未知的event类型")
	}
}
