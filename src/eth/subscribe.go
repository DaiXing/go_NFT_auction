package eth

import (
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

}
