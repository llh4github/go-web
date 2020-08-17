package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var v *viper.Viper

func init() {
	v = viperReader()
}

func viperReader() *viper.Viper {
	v := viper.New()
	v.SetConfigName("app")
	v.SetConfigType("toml")
	v.AddConfigPath("../resources/")
	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("err:%s\n", err)
	}
	return v
}

// GetBDConfig 获取数据库配置
func GetBDConfig() DBConfig {
	var config Config
	err := v.Unmarshal(&config)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v \n", err)
	}
	return config.Dbconfig
}

// Config 所有配置项
type Config struct {
	Name     string
	Dbconfig DBConfig
}

// DBConfig 数据库连接配置
type DBConfig struct {
	Username, Password, Host, Dbname, Params string
}
