package util

import "github.com/spf13/viper"

// 配置项。
var ConfigParams ConfigParamInfo

// 初始化。
func InitViper() {
	viper.AddConfigPath(".")

	viper.SetConfigName("project")
	viper.SetConfigType("yaml")

	err1 := viper.ReadInConfig()
	if err1 != nil {
		panic(err1)
	}

	// 结构化。
	err2 := viper.Unmarshal(&ConfigParams)
	if err2 != nil {
		panic(err2)
	}
}

// 配置的参数。
type ConfigParamInfo struct {
	Datasource DbParamInfo  `mapstructure:"datasource"`
	Eth        EthParamInfo `mapstructure:"eth"`
}

// 数据库
type DbParamInfo struct {
	MysqlUrl       string `mapstructure:"mysqlUrl"`
	NeedDropTables bool   `mapstructure:"needDropTables"`
}

// eth
type EthParamInfo struct {
	RpcUrl                 string `mapstructure:"rpcUrl"`
	AuctionContractAddress string `mapstructure:"auctionContractAddress"`
}
