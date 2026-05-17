package web

import "my.nft.auction/src/database"

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
	TokenOwner string `json:"tokenOwner" ` // token owner
}
type TokenInfo struct {
	NftContract string `json:"nftContract"  ` // NFT合约地址
	TokenId     uint64 `json:"tokenId"  `     // token
	TokenUri    string `json:"tokenUri"  `    // token uri
}
type GetTokenListResp struct {
	BaseResp
	TokenList []TokenInfo `json:"tokenList"` // token列表。
}

// 查询 拍卖列表。
type GetAuctionListReq struct {
	BaseReq
	NftContract string `json:"nftContract"  ` // NFT合约地址
	TokenId     uint64 `json:"tokenId"  `     // token
	TokenOwner  string `json:"tokenOwner" `   // token owner
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
