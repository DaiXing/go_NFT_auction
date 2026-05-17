package eth

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"my.nft.auction/src/database"
	"my.nft.auction/src/util"
)

// 扫描历史事件。
func ScanHistoryEvent() {
	logx := util.LogMaker{}
	defer logx.LogString()
	logx.AddLine(">> 扫描历史事件。")

	// 取DB保存的标记。
	value := database.QueryKeyValue(util.KEY_SCAN_BLOCK_NUM, "0")
	blockNumScan, err := strconv.ParseUint(value, 10, 64)
	util.CheckError(err)
	logx.AddKV("  已经处理的blockNum", blockNumScan)

	ctx, cancel := util.NewContext(3)
	defer cancel()

	// 取最新的区块号。
	blockNumLast, err2 := EthClient.BlockNumber(ctx)
	util.CheckError(err2)
	logx.AddKV("  最新的blockNum", blockNumLast)

	// 扫描日志。 一次处理几个区块。
	onceSize := 3 // 每次处理2个区块。
	for blockNum := blockNumScan + 1; blockNum <= blockNumLast; blockNum += uint64(onceSize) {
		begin := blockNum
		end := begin + uint64(onceSize) - 1
		logx.AddKV(" 查询 blockNum 范围：  ", fmt.Sprintf("%d ~ %d ", begin, end))

		// 查询一次。
		query := ethereum.FilterQuery{
			Addresses: []common.Address{
				contractAddr, // 合约地址
			},
			FromBlock: big.NewInt(int64(begin)), // 区块范围
			ToBlock:   big.NewInt(int64(end)),   // 区块范围
		}
		logList, err2 := EthClient.FilterLogs(ctx, query)
		util.CheckError(err2)
		logx.AddKV("  log 数量", len(logList))

		// 处理日志。
		for _, log := range logList {
			parseEvent(&log)
		}
	}

	// 更新DB。 往前一个区块，可能只有部分数据。
	valueNew := strconv.FormatUint(blockNumLast-1, 10)
	database.UpdateKeyValue(util.KEY_SCAN_BLOCK_NUM, valueNew)
	logx.AddKV("  更新已经处理的blockNum", valueNew)
}
