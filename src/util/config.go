package util

import "github.com/spf13/viper"

// 配置项。
var Params ConfigParamInfo

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
	err2 := viper.Unmarshal(&Params)
	if err2 != nil {
		panic(err2)
	}
	Logger.Info("Viper  初始化完成")
}

// 配置的参数。
type ConfigParamInfo struct {
	Datasource DbParamInfo   `mapstructure:"datasource"`
	Eth        EthParamInfo  `mapstructure:"eth"`
	Web        WebParamInfo  `mapstructure:"web"`
	Mock       MockParamInfo `mapstructure:"mock"`
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
	AlchemyKey             string `mapstructure:"alchemyKey"`
	AlchemyNet             string `mapstructure:"alchemyNet"`
}

// web配置
type WebParamInfo struct {
	JwtKey           string `mapstructure:"jwtKey"`
	TokenValidMiutes uint64 `mapstructure:"tokenValidMiutes"`
	Port             uint32 `mapstructure:"port"`
}

// mock测试
type MockParamInfo struct {
	NftContractAddr string `mapstructure:"nftContractAddr"`
	JackAddr        string `mapstructure:"jackAddr"`
	JackPrivateKey  string `mapstructure:"jackPrivateKey"`
	TomAddr         string `mapstructure:"tomAddr"`
	TomPrivateKey   string `mapstructure:"tomPrivateKey"`
}
