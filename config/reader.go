package config

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const myApp = "go-web"

func viperReader() *viper.Viper {
	v := viper.New()
	v.SetConfigName("app")
	v.SetConfigType("toml")
	v.AddConfigPath(getProjectDir() + "/resources/")
	if err := v.ReadInConfig(); err != nil {
		log.Error(err)
		// panic("读取配置文件失败！")
	}
	return v
}

// getWorkDir 获取当前项目目录
func getProjectDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Error("获取当前项目目录失败！", err)
		panic("获取当前项目目录失败！")
	}
	paths := strings.Split(dir, myApp)
	projectDir := paths[0] + "/" + myApp
	if info, e := os.Stat(projectDir); e != nil && !info.IsDir() {
		log.Error("获取当前项目目录失败！", err)
		panic("获取当前项目目录失败！")
	}
	return projectDir
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
