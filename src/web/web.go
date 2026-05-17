package web

import "github.com/gin-gonic/gin"

var webServer *gin.Engine

func InitWeb() {
	webServer = gin.Default()

}

func setupPath() {

}
