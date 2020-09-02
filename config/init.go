package config

import (
	"fmt"
	"time"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql" // mysql dirve
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	// 配置文件读取组件
	v *viper.Viper
	// 配置信息结构体
	config Config
	// MyDB 当前数据库连接实例
	MyDB *gorm.DB
	// CasibnAdapter Casibn适配器
	CasibnAdapter *gormadapter.Adapter
	// Enforcer casbin权限判定对象
	Enforcer *casbin.Enforcer
)

func init() {
	readConfig()
	connectionDB()
	// addCasibnAdapter()
}
func readConfig() {
	v = viperReader()
	err := v.Unmarshal(&config)
	if err != nil {
		log.Errorf("unable to decode into struct, %v \n", err)
		panic("不能把配置文件内容序列化到结构体中！")
	}
}

// 添加 Casibn Adapter
func addCasibnAdapter() {
	var err error
	CasibnAdapter, err = gormadapter.NewAdapterByDB(MyDB)
	if err != nil {
		fmt.Println("NewAdapterByDB error : ", err)
	}
	a, e := casbin.NewEnforcer(getProjectDir()+"/resources/rbac_models.conf", CasibnAdapter)
	if e != nil {
		log.Debug("fuck : ", a)
		log.Error("NewEnforcer error : ", err)
	} else {
		log.Debug("fuck : ", a)
	}
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
	res, err := Enforcer.AddPolicy("admin", "/api/v2/ping", "GET")
	if !res {
		fmt.Println("policy is exist")
	} else {
		fmt.Println("policy is not exist, adding")
	}
	// log.Errorln(Enforcer.AddRoleForUser)
	// 将 test 用户加入一个角色中
	// Enforcer.AddRoleForUser("test", "root")
	Enforcer.AddRoleForUser("tom", "admin")
}

// 数据库连接
func connectionDB() {

	sqlLogger := logger.New(
		log.New(), // logrus 组件
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)
	//dbConfig() should like "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dbConfig()), &gorm.Config{
		Logger: sqlLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表名单数
			TablePrefix:   "",   // 表前缀
		},
	})
	MyDB = db.Debug()

	if err != nil {
		// gorm v2 不推荐使用db.Close()方法了
		// 其方法调用隐藏的更深了。
		// 暂时不Close()看看吧
		_db, _ := MyDB.DB()
		_db.Close()
		log.Errorf("database connect error: %v \n", err)
		panic("数据库连接失败！")
	}

}
