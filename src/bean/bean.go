package bean

import (
	"github.com/golang-jwt/jwt/v5"
	"my.nft.auction/src/database"
)

// token
type JwtTokenInfo struct {
	jwt.RegisteredClaims        // 继承。
	UserId               uint64 `json:"userId"`
	Username             string `json:"username"`
}

// 父类。
type BaseReq struct {
	PageNo   int `json:"pageNo"`
	PageSize int `json:"pageSize"`
}
type BaseResp struct {
	Error       string `json:"error"`
	ErrorStatus int    `json:"errorStatus"`
	Message     string `json:"message"`
	TotalSize   int64  `json:"totalSize"`
}

// 查询 token列表。
type GetTokenListReq struct {
	BaseReq
	Seller string `json:"seller" ` // 卖家。
}
type TokenInfo struct {
	NftContract string                  `json:"nftContract"  ` // NFT合约地址
	TokenId     string                  `json:"tokenId"  `     // token
	TokenUri    string                  `json:"tokenUri"  `    // token uri
	TokenType   string                  `json:"tokenType"`
	FloorPrice  float32                 `json:"floorPrice"` // 地板价
	Description string                  `json:"description"`
	Image1      string                  `json:"image1"`
	Image2      string                  `json:"image2"`
	Image3      string                  `json:"image3"`
	AuctionInfo *database.AuctionInfoPo `json:"auctionInfo"` // 拍卖。
}
type GetTokenListResp struct {
	BaseResp
	TokenList []*TokenInfo `json:"tokenList"` // token列表。
}

// 创建 拍卖
type CreateAuctionReq struct {
	BaseReq
	NftContract string `json:"nftContract"  ` // NFT合约地址
	TokenId     string `json:"tokenId"  `     // token
	MinPrice    int64  `json:"minPrice" `     //
	BeginTime   int64  `json:"beginTime" `    //
	PeriodTime  int64  `json:"periodTime" `   //
}

// 出价
type BidAuctionReq struct {
	BaseReq
	AuctionId  string `json:"auctionId"  ` // 拍卖ID
	BidPrice   int64  `json:"bidPrice" `   // 出价
	BidderName string `json:"bidderName" ` // 用户
}

// 查询 拍卖列表。
type GetAuctionListReq struct {
	BaseReq
	NftContract string `json:"nftContract"  ` // NFT合约地址
	TokenId     string `json:"tokenId"  `     // token
	Seller      string `json:"seller" `       // 卖家
	AuctionId   uint64 `json:"auctionId" `    // 拍卖ID。唯一。
}
type GetAuctionListResp struct {
	BaseResp
	AuctionList []database.AuctionInfoPo `json:"auctionList"` // 拍卖列表。
}

// 查询 出价列表。
type GetBidListReq struct {
	BaseReq
	AuctionId uint64 `json:"auctionId" ` // 拍卖ID。唯一。
}
type GetBidListResp struct {
	BaseResp
	BidList []database.AuctionBidPo `json:"bidList"` // 出价列表。
}

// 统计。
type StatisticResp struct {
	BaseResp
	CountAuction int64 `json:"countAuction"` // 拍卖总数
	CountBid     int64 `json:"countBid"`     // 出价总数
	SumTvl       int64 `json:"sumTvl"`       // 总锁仓价值
}
