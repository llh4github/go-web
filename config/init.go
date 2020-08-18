package config

import (
	_ "github.com/go-sql-driver/mysql" // mysql dirve
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var v *viper.Viper
var config Config

func init() {
	readConfig()
}
func readConfig() {
	v = viperReader()
	err := v.Unmarshal(&config)
	if err != nil {
		log.Errorf("unable to decode into struct, %v \n", err)
		panic("不能把配置文件内容序列化到结构体中！")
	}
}
func connectionDB() {
	// "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dbConfig())
	defer db.Close()
	if err != nil {
		log.Errorf("database connect error: %v \n", err)
		// panic("数据库连接失败！")
	}
}
