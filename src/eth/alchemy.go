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

func AlchemyQueryNft(owner string, pageSize int, needMeta bool) {
	//owner=0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045&withMetadata=false&pageSize=2
	url1 := alchemyBaseUrl + "/nft/v3/%s/getNFTsForOwner"
	url1 = fmt.Sprintf(url1, util.Params.Eth.AlchemyKey)

	// 参数。
	params:= map[string] string {
		"owner": owner, 
		"pageSize": strconv.FormatInt(int64(pageSize), 10), 
		"withMetadata": strconv.FormatBool(needMeta), 
	}

	util.HttpGetJson2[]()
}
