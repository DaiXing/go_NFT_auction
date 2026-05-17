package database

import (
	"time"

	"gorm.io/gorm"
)

// 基本字段
type BasePo struct {
	ID        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// 拍卖信息。
type AuctionInfoPo struct {
	BasePo
	NftContract string // NFT合约地址
	TokenId     uint64 // token
	TokenOwner  string // token owner
	AuctionId   uint64 // 拍卖ID
	MinPrice    uint64 // 起拍价。
	Bidder      string // 当前最高竞拍人
	BidPrice    uint64 // 当前最高竞拍价格
	State       int    // 状态。
}

// 竞拍信息。 每个人的出价。
type AuctionBidPo struct {
	BasePo
	AuctionId uint64 // 拍卖ID
	Bidder    string // 竞拍人
	BidPrice  uint64 // 竞拍价格
	BidTime   uint64 // 竞拍时间。
}
