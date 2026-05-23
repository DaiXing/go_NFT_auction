package eth

import (
	"fmt"
	"strconv"

	"my.nft.auction/src/util"
)

var alchemyBaseUrl string

func InitAlchemy() {
	// URL
	alchemyBaseUrl = fmt.Sprintf("https://%s.g.alchemy.com", util.Params.Eth.AlchemyNet)
}

// 查询 NFT
// https://eth-mainnet.g.alchemy.com/nft/v3/umAjVzLoWaqMIxZh9OFn2/getNFTsForOwner?owner=0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045&withMetadata=false&pageSize=2
func AlchemyQueryNft(owner string, pageSize int, needMeta bool) (*AxNftResp, error) {
	url1 := alchemyBaseUrl + "/nft/v3/%s/getNFTsForOwner"
	url1 = fmt.Sprintf(url1, util.Params.Eth.AlchemyKey)

	// 参数。
	params := map[string]string{
		"owner":        owner,
		"pageSize":     strconv.FormatInt(int64(pageSize), 10),
		"withMetadata": strconv.FormatBool(needMeta),
	}

	// 查询
	resp, err := util.HttpGetJson2[AxNftResp](url1, params)
	return resp, err
}

type AxContract struct {
	Name            string      `json:"name"`
	Address         string      `json:"address"`
	Symbol          string      `json:"symbol"`
	TokenType       string      `json:"tokenType"`
	OpenSeaMetadata AxTokenMeta `json:"openSeaMetadata"`
}
type AxToken struct {
	TokenId     string       `json:"tokenId"`
	TokenType   string       `json:"tokenType"`
	Description string       `json:"description"`
	TokenUri    string       `json:"tokenUri"`
	Balance     string       `json:"balance"`
	Contract    AxContract   `json:"contract"`
	Image       AxTokenImage `json:"image"`
}
type AxTokenImage struct {
	CachedUrl    string `json:"cachedUrl"`
	ThumbnailUrl string `json:"thumbnailUrl"`
	PngUrl       string `json:"pngUrl"`
}
type AxNftResp struct {
	OwnedNfts  []AxToken `json:"ownedNfts"`
	TotalCount int32     `json:"totalCount"`
	// PageKey    string    `json:"pageKey"`
}
type AxTokenMeta struct {
	FloorPrice float32 `json:"floorPrice"`
	ImageUrl   string  `json:"imageUrl"`
}
