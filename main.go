package main

import (
	"my.nft.auction/src/database"
	"my.nft.auction/src/eth"
	"my.nft.auction/src/web"
)

func main() {
	// 数据库
	database.InitClient() // 客户端连接
	database.InitTables() // 初始化表

	// 以太坊
	eth.InitClient()       // 客户端连接
	eth.InitContract()     // 合约
	eth.ScanHistoryEvent() // 扫描历史事件
	eth.SubscribeEvent()   // 订阅事件

	// web
	web.InitWeb() // web 服务器
}
