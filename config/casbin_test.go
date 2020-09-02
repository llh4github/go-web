package config

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// 测试 使用适配器直接连接数据库
// 此处会默认使用是casbin的数据库，不存在会创建
func TestConn(t *testing.T) {
	// 配置数据库源时，只能配置如下形式的，不能配置更多的东西了
	a, err := gormadapter.NewAdapter("mysql", "root:root@tcp(127.0.0.1:3306)/")
	e, err := casbin.NewEnforcer(getProjectDir()+"/resources/rbac_models.conf", a)
	if err != nil {
		println(err)
	}
	println(e.GetAdapter())
}

// 测试使用已有DB连接
func TestConn2(t *testing.T) {
	a, err := gormadapter.NewAdapterByDB(dbConn())
	if err != nil {
		println(err)
	}
	println("gormadapter.Adapter  ", a.IsFiltered())
	e, er := casbin.NewEnforcer(getProjectDir()+"/resources/rbac_models.conf", a)

	if er != nil {
		println(er)
	}
	println("Enforcer : ", e)
	e.LoadPolicy()
	println(e.GetAdapter())
}
func dbConn() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // 禁用彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/go_web?charset=utf8&parseTime=True&loc=Local"),
		&gorm.Config{
			Logger: newLogger,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 表名单数
				TablePrefix:   "",   // 表前缀
			},
		})
	if err != nil {
		println("数据库连接失败！")
	}
	return db
}
