package main

import (
	"my.nft.auction/src/database"
	"my.nft.auction/src/eth"
	"my.nft.auction/src/util"
	"my.nft.auction/src/web"
)

func main() {
	// 基础服务。
	util.InitLogger()
	util.InitViper()

	// 数据库
	database.InitDb()     // 客户端连接
	database.InitTables() // 初始化表

	// 以太坊
	eth.InitAlchemy()  // NFT查询
	eth.InitClient()   // 客户端连接
	eth.InitContract() // 合约
	// eth.ScanHistoryEvent() // 扫描历史事件
	eth.SubscribeEvent() // 订阅事件

	// web
	web.InitWeb() // web 服务器
}
