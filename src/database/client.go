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
	url := util.ConfigParams.Datasource.MysqlUrl
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 可选：设置日志级别
	})
	util.CheckError(err)

	Db = db

	util.Logger.Info("DB 初始化连接 完成 ", "url", url)
}

// 初始化表。
func InitTables() {
	// 表。
	tableList := []any{
		&AuctionInfoPo{}, &AuctionBidPo{}, &KeyValuePo{},
	}

	// 丢弃表。
	if util.ConfigParams.Datasource.NeedDropTables {
		err := Db.Migrator().DropTable(tableList)
		util.CheckError(err)
	}

	// 自动迁移。
	err := Db.AutoMigrate(tableList)
	util.CheckError(err)

	// 查表。
	tableNames, err3 := Db.Migrator().GetTables()
	util.CheckError(err3)

	util.Logger.Info("DB 初始化表 完成", "tables", tableNames)
}
