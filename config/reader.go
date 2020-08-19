package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func viperReader() *viper.Viper {
	v := viper.New()
	v.SetConfigName("app")
	v.SetConfigType("toml")
	v.AddConfigPath("../resources/")
	if err := v.ReadInConfig(); err != nil {
		log.Error(err)
		panic("读取配置文件失败！")
	}
	return v
}

// GetBDConfig 获取数据库配置
func GetBDConfig() DBConfig {

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