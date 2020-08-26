package config

import (
	"fmt"

	"github.com/casbin/casbin"
	gormadapter "github.com/casbin/gorm-adapter"
	_ "github.com/go-sql-driver/mysql" // mysql dirve
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var v *viper.Viper
var config Config

// MyDB 当前数据库连接实例
var MyDB *gorm.DB

// CasibnAdapter Casibn适配器
var CasibnAdapter *gormadapter.Adapter

// Enforcer casbin权限判定对象
var Enforcer *casbin.Enforcer

func init() {
	readConfig()
	connectionDB()
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
	var err error
	MyDB, err = gorm.Open("mysql", dbConfig())

	CasibnAdapter = gormadapter.NewAdapterByDB(MyDB)
	Enforcer = casbin.NewEnforcer(getProjectDir()+"/resources/rbac_models.conf", CasibnAdapter)
	// 开启权限认证日志
	Enforcer.EnableLog(true)
	// 加载数据库中的策略
	err = Enforcer.LoadPolicy()
	if err != nil {
		log.Errorf("加载数据库中的策略失败: %v \n", err)
		// panic("数据库连接失败！")
	}
	// 创建一个角色,并赋于权限
	// admin 这个角色可以访问GET 方式访问 /api/v2/ping
	res := Enforcer.AddPolicy("admin", "/api/v2/ping", "GET")
	if !res {
		fmt.Println("policy is exist")
	} else {
		fmt.Println("policy is not exist, adding")
	}
	log.Errorln(Enforcer.AddRoleForUser)
	// 将 test 用户加入一个角色中
	// Enforcer.AddRoleForUser("test", "root")
	Enforcer.AddRoleForUser("tom", "admin")
	MyDB.LogMode(true)
	if err != nil {
		defer MyDB.Close()
		log.Errorf("database connect error: %v \n", err)
		panic("数据库连接失败！")
	}
	MyDB.SingularTable(true)
}
