package eth

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// 创建。
// event AuctionCreate(
//
//	address indexed nftContract, // NFT合约地址
//	address tokenOwner, // token owner
//	uint256 indexed tokenId, // token
//	uint256 indexed auctionId, // 拍卖
//	uint256 minPrice, // 起拍价
//	uint256 beginTime, // 开始时间
//	uint256 endTime // 结束时间
//
// );
type EventAuctionCreate struct {
	NftContract common.Address
	TokenOwner  common.Address
	TokenId     *big.Int
	AuctionId   *big.Int
	MinPrice    *big.Int
	BeginTime   *big.Int
	EndTime     *big.Int
}

// 退款。
// event AuctionRefund(
//
//	uint256 indexed auctionId,
//	address indexed to, // 给谁。
//	uint256 amount // 退款金额
//
// );
type EventAuctionRefund struct {
	AuctionId *big.Int
	To        common.Address
	Amount    *big.Int
}

// 竞拍。
// event AuctionBid(
//
//	uint256 indexed auctionId,
//	address indexed bidder, // 竞拍人。
//	uint256 bidPrice // 竞拍金额
//
// );
type EventAuctionBid struct {
	AuctionId *big.Int
	Bidder    common.Address
	BidPrice  *big.Int
}

// 取消。
// event AuctionCancel(uint256 indexed auctionId);
type EventAuctionCancel struct {
	AuctionId *big.Int
}

// 结束。
// event AuctionEnd(uint256 indexed auctionId, AuctionState state);
type EventAuctionEnd struct {
	AuctionId *big.Int
	State     uint8
}
