package eth

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"my.nft.auction/src/util"
)

// 客户端。
var Client *ethclient.Client

// 初始化。
func InitClient() {
	ctx, cancel := util.NewContext(3)
	defer cancel()

	// URL
	url := util.ConfigParams.Eth.RpcUrl
	// 连接。
	client, err := ethclient.DialContext(ctx, url)
	if err != nil {
		util.CheckError(err)
	}
	Client = client

	util.Logger.Info("Eth 初始化 client 完成")
}
