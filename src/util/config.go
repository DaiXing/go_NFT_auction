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
	Mysql MysqlParamInfo `mapstructure:"mysql"`
}
type MysqlParamInfo struct {
	Url string `mapstructure:"url"`
}
