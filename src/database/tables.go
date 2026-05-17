package database

import (
	"time"

	"gorm.io/gorm"
)

// 基本字段
type BasePo struct {
	ID        uint64         `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

// 拍卖信息。
type AuctionInfoPo struct {
	BasePo
	NftContract string `json:"nftContract" gorm:"index;size:100"` // NFT合约地址
	TokenId     uint64 `json:"tokenId" gorm:"index"`              // token
	TokenOwner  string `json:"tokenOwner" gorm:"index;size:100"`  // token owner
	AuctionId   uint64 `json:"auctionId" gorm:"index"`            // 拍卖ID
	MinPrice    uint64 `json:"minPrice" `                         // 起拍价。
	BeginTime   uint64 `json:"beginTime"`                         // 开始时间
	PeriodTime  uint64 `json:"periodTime"`                        // 持续时间。
	Bidder      string `json:"bidder" gorm:"size:100"`            // 当前最高竞拍人
	BidPrice    uint64 `json:"bidPrice"`                          // 当前最高竞拍价格
	State       int    `json:"state"`                             // 状态。
}

// 竞拍信息。 每个人的出价。
type AuctionBidPo struct {
	BasePo
	LogKey    string `json:"logKey" gorm:"uniqueIndex;size:100"` // 日志的唯一标识。
	AuctionId uint64 `json:"auctionId" gorm:"index"`             // 拍卖ID
	Bidder    string `json:"bidder" gorm:"index;size:100"`       // 竞拍人
	BidPrice  uint64 `json:"bidPrice"`                           // 竞拍价格
	BidTime   uint64 `json:"bidTime"`                            // 竞拍时间。
}
