package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func viperReader() *viper.Viper {
	v := viper.New()
	v.SetConfigName("app")
	v.SetConfigType("toml")
	v.AddConfigPath("./resources/")
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

// GetJwtConfig jwt的配置
func GetJwtConfig() JwtConfig {
	return config.Jwt
}

// Config 所有配置项
type Config struct {
	Name     string
	Dbconfig DBConfig
	Jwt      JwtConfig
}

// DBConfig 数据库连接配置
type DBConfig struct {
	Username, Password, Host, Dbname, Params string
}

// JwtConfig 生成jwt的配置
type JwtConfig struct {
	Secret string
	Iss    string
	// Exp 过期时间（分钟）。务必写入正整数
	Exp int
}
