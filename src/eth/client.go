package eth

import (
	"strings"

	"github.com/ethereum/go-ethereum/ethclient"
	"my.nft.auction/src/util"
)

// 客户端。 http , WS
var EthClient *ethclient.Client
var EthClientWS *ethclient.Client

// 初始化。
func InitClient() {
	ctx, cancel := util.NewContext(3)
	defer cancel()

	// URL
	url := util.Params.Eth.RpcUrl
	urlWs := strings.Replace(url, "http", "ws", 1)

	// 连接。
	client, err := ethclient.DialContext(ctx, url)
	util.CheckError(err)
	clientWs, err2 := ethclient.DialContext(ctx, urlWs)
	util.CheckError(err2)

	EthClient = client
	EthClientWS = clientWs

	util.Logger.Info("Eth 初始化 client 完成", "url", url, "urlWs", urlWs)
}
