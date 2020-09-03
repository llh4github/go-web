package config

import (
	"gitee.com/llh-gitee/go-web/common"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper" // mysql dirve
	"gorm.io/gorm"
)

var (
	// 配置文件读取组件
	v *viper.Viper
	// 配置信息结构体
	// Deprecated
	config Config
	// MyDB 当前数据库连接实例
	// Deprecated
	MyDB *gorm.DB
	// CasibnAdapter Casibn适配器
	// Deprecated
	CasibnAdapter *gormadapter.Adapter
	// Enforcer casbin权限判定对象
	// Deprecated
	Enforcer *casbin.Enforcer
)

func init() {
	readConfig()
	connectionDB()
	// addCasibnAdapter()
}

// 读取配置文件的方法
// 务必让它先执行
func readConfig() {
	v = viperReader()
	err := v.Unmarshal(&common.MyConf)
	if err != nil {
		log.Errorf("unable to decode into struct, %v \n", err)
		panic("不能把配置文件内容序列化到结构体中！")
	}
}
