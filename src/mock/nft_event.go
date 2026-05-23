package mock

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"my.nft.auction/src/database"
	"my.nft.auction/src/util"
)

// nft 转账。
type EventNftTransfer struct {
	NftContract common.Address
	From        common.Address
	To          common.Address
	TokenId     *big.Int
}

// 处理事件。
func handleEventNftTransfer(log *types.Log, event *abi.Event, eventx *EventNftTransfer, logx *util.LogMaker) {
	// 去重。
	nftContract := eventx.NftContract.Hex()
	from := eventx.From.Hex()
	to := eventx.To.Hex()
	tokenId := eventx.TokenId.String()
	logx.AddKV("  nftContract", nftContract)
	logx.AddKV("  from", from)
	logx.AddKV("  to", to)
	logx.AddKV("  tokenId", tokenId)

	// from 是 0 。表示新建。
	if eventx.From == (common.Address{}) {
		logx.AddKV("  类型", "新建token")

		if database.ExistToken(nftContract, tokenId) {
			logx.AddLine("  token已经在DB。不再写入。")
			return
		}

		var one database.MockTokenPo
		one.NftContract = nftContract
		one.Creator = to
		one.Owner = to
		one.TokenId = tokenId
		err2 := database.Db.Create(&one).Error
		if err2 == nil {
			logx.AddLine("  插入表 成功")
		} else {
			logx.AddKV("  插入表 错误", err2)
		}
	} else {
		logx.AddKV("  类型", "转移token")

		// 更新。
		err2 := database.UpdateToken(nftContract, tokenId, map[string]any{
			"owner": to,
		})
		if err2 == nil {
			logx.AddLine("  更新表 成功")
		} else {
			logx.AddKV("  更新表 错误", err2)
		}
	}
}
