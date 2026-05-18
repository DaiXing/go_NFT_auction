package eth

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// 创建。
type EventAuctionCreate struct {
	NftContract common.Address
	Seller      common.Address
	TokenId     *big.Int
	AuctionId   *big.Int
	MinPrice    *big.Int
	BeginTime   *big.Int
	EndTime     *big.Int
}

// 退款。
type EventAuctionRefund struct {
	AuctionId *big.Int
	To        common.Address
	Amount    *big.Int
	BidId     *big.Int
}

// 竞拍。
type EventAuctionBid struct {
	AuctionId *big.Int
	Bidder    common.Address
	BidPrice  *big.Int
	BidId     *big.Int
}

// 取消。
type EventAuctionCancel struct {
	AuctionId *big.Int
}

// 结束。
type EventAuctionEnd struct {
	AuctionId *big.Int
	State     uint8
}
