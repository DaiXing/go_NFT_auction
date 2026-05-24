package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"my.nft.auction/src/bean"
)

func webReturnJson(ctx *gin.Context, status int, body any) {
	ctx.JSON(status, body)
}

func webReturnOKJson(ctx *gin.Context, body any) {
	ctx.JSON(http.StatusOK, body)
}

func webAbortError(ctx *gin.Context, status int, err string) {
	ctx.AbortWithStatusJSON(status, &bean.BaseResp{
		Error:       err,
		ErrorStatus: status,
	})
}

func webAbortBadReq(ctx *gin.Context, err string) {
	webAbortError(ctx, http.StatusBadRequest, err)
}
