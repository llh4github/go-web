package config

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const myApp = "go-web"

func viperReader() *viper.Viper {
	v := viper.New()
	v.SetDefault("Port", "8080")
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
// Deprecated
func GetBDConfig() DBConfig {

	return config.Dbconfig
}

// GetURLPrefix 获取配置文件中url前缀
//
// Deprecated
func GetURLPrefix() string {
	fmt.Println("GetURLPerfix : " + config.URLPrefix)
	return config.URLPrefix
}

// GetJwtConfig jwt的配置
// Deprecated
func GetJwtConfig() JwtConfig {
	return config.Jwt
}

// Config 所有配置项
// Deprecated
type Config struct {
	// URLPrefix gin会判断是不是以 / 开头，不用考虑gin配置。为了使用方便请以 / 开头配置此项目
	URLPrefix string

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
