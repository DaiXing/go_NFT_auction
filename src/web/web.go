package web

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"my.nft.auction/src/util"
)

// 服务器。
var webServer *gin.Engine

// 初始化web。
func InitWeb() {
	webServer = gin.Default()
	setupPath()

	addr := fmt.Sprintf(":%d", util.Params.Web.Port)
	util.Logger.Info("初始化 web ", "addr", addr)

	webServer.Run(addr)
}

// 设置路径。
func setupPath() {
	webServer.Use(aopLogRequest) // 打日志。

	// 健康检测
	webServer.GET("/health", handleHealth)

	// 代币
	group1 := webServer.Group("/token")
	group1.POST("/get-token-list", handleGetTokenList)

	// 拍卖
	group2 := webServer.Group("/auction")
	group2.POST("/get-auction-list", handleGetAuctionList)

	// 出价
	group3 := webServer.Group("/bid")
	group3.POST("/get-bid-list", handleGetBidList)
}
