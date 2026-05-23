package mock

import (
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"my.nft.auction/src/eth"
	"my.nft.auction/src/util"
)

// 订阅事件。
func NftSubscribeEvent() {
	ctx, cancel := util.NewContext(3)
	defer cancel()

	// 订阅日志。
	query := ethereum.FilterQuery{
		Addresses: []common.Address{nftContract},
	}
	chanEvent := make(chan types.Log) // 队列。
	sub, err2 := eth.EthClientWS.SubscribeFilterLogs(ctx, query, chanEvent)
	util.CheckError(err2)

	go func() {
		for {
			select {
			case <-sub.Err():
				util.Logger.Error("订阅事件 发生错误", "err", sub.Err())

				sub.Unsubscribe() // 取消订阅。

				NftSubscribeEvent() // 重新订阅。

				return
			case logE := <-chanEvent: // 解析日志

				parseOneEvent(&logE) // 处理日志。

				// updateScanBlockNum(logE.BlockNumber) // 更新DB。
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
	event, err1 := nftABI.EventByID(topic0)
	util.CheckError(err1)

	// 合约地址。
	nftContractAddr := log.Address

	eventName := event.Name
	logx.AddKV("  eventName", eventName)

	// 区分事件
	if "Transfer" == eventName {
		// 解析事件。
		var eventx EventNftTransfer
		// err2 := abiObj.UnpackIntoInterface(&eventx, eventName, log.Data)
		// util.CheckError(err2)

		eventx.NftContract = nftContractAddr
		eventx.From = common.BytesToAddress(log.Topics[1].Bytes())
		eventx.To = common.BytesToAddress(log.Topics[2].Bytes())
		eventx.TokenId = big.NewInt(0).SetBytes(log.Topics[3].Bytes())

		logx.AddKV("  event明细", util.ToJson(eventx))

		// 处理事件。
		handleEventNftTransfer(log, event, &eventx, &logx)
	} else {
		// 未知事件。
		logx.AddKV("  忽略", "未知的event类型")
	}
}
