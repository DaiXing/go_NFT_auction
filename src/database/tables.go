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
	AuctionId   uint64 `json:"auctionId" gorm:"uniqueIndex"`      // 拍卖ID。唯一。
	MinPrice    uint64 `json:"minPrice" `                         // 起拍价。
	BeginTime   uint64 `json:"beginTime"`                         // 开始时间
	PeriodTime  uint64 `json:"periodTime"`                        // 持续时间。
	Bidder      string `json:"bidder" gorm:"size:100"`            // 当前最高竞拍人
	BidPrice    uint64 `json:"bidPrice"`                          // 当前最高竞拍价格
	State       uint8  `json:"state"`                             // 状态。
}

// 竞拍信息。 每个人的出价。
type AuctionBidPo struct {
	BasePo
	AuctionId    uint64 `json:"auctionId" gorm:"index"`       // 拍卖ID
	BidId        uint64 `json:"bidId" gorm:"uniqueIndex"`     // 竞拍的序号。 唯一。
	Bidder       string `json:"bidder" gorm:"index;size:100"` // 竞拍人
	BidPrice     uint64 `json:"bidPrice"`                     // 竞拍价格
	BidTime      uint64 `json:"bidTime"`                      // 竞拍时间。
	RefundAmount uint64 `json:"refundAmount"`                 // 退款金额。 只有被超越的竞拍才有退款金额。
	RefundTime   uint64 `json:"refundTime"`                   // 退款的时间。
}

// KV 数据。
type KeyValuePo struct {
	BasePo
	ParamKey   string `json:"paramKey" gorm:"uniqueIndex;size:200"` // 键。唯一。
	ParamValue string `json:"paramValue" gorm:"size:2000"`          // 值。
}
