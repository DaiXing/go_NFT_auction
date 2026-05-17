package database

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"my.nft.auction/src/util"
)

// 连接。
var Db *gorm.DB

func InitClient() {
	// 连接。
	url := util.ConfigParams.Mysql.Url
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 可选：设置日志级别
	})
	util.CheckError(err)

	Db = db

	util.Logger.Info("DB 初始化完成 ", "url", url)
}
