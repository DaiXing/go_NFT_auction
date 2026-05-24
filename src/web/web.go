package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"my.nft.auction/src/mock"
	"my.nft.auction/src/util"
)

// 服务器。
var webServer *gin.Engine

// 初始化web。
func InitWeb() {
	webServer = gin.Default()

	setupPath()
	mock.SetupMockPath(webServer)

	addr := fmt.Sprintf(":%d", util.Params.Web.Port)
	util.Logger.Info("初始化 web ", "addr", addr)

	err2 := webServer.Run(addr)
	util.CheckError(err2)
}

// 设置路径。
func setupPath() {
	// 打日志。
	webServer.Use(aopLogRequest)
	// 异常捕获
	webServer.Use(gin.CustomRecovery(func(ctx *gin.Context, err any) {
		errmsg := "Error "
		if err != nil {
			errmsg = errmsg + fmt.Sprint(err)
		}
		webAbortError(ctx, http.StatusInternalServerError, errmsg)
	}))

	// 健康检测
	webServer.GET("/health", pathHealth)

	// 代币
	group1 := webServer.Group("/token")
	group1.POST("/get-token-list", pathGetTokenList)

	// 拍卖
	group2 := webServer.Group("/auction")
	group2.POST("/get-auction-list", pathGetAuctionList)

	// 出价
	group3 := webServer.Group("/bid")
	group3.POST("/get-bid-list", pathGetBidList)

	// 全局
	group4 := webServer.Group("/global")
	group4.GET("/statistic", pathStatistic)
}
