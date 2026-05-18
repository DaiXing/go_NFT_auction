package web

import (
	"time"

	"github.com/gin-gonic/gin"
	"my.nft.auction/src/util"
)

// 打日志。
func aopLogRequest(ctx *gin.Context) {
	startTime := time.Now()
	uri := ctx.Request.URL.Path
	method := ctx.Request.Method

	ctx.Next()

	costMillis := time.Since(startTime).Milliseconds()
	util.Logger.Info("HTTP REQ", "method", method, "uri", uri, "costMillis", costMillis)
}

// AOP 验证token 。
func aopVerifyToken(ctx *gin.Context) {
	// 解析token
	tokenx := ctx.GetHeader(util.KEY_USER_TOKEN)
	pxToken := jwtVerifyToken(tokenx)

	// 保存。
	ctx.Set(util.KEY_USER_ID, pxToken.UserId)
	ctx.Set(util.KEY_USER_NAME, pxToken.Username)
}
