package config

import (
	"fmt"
	"time"

	"gitee.com/llh-gitee/go-web/common"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var dbConf = common.MyConf.Dbconfig

func dbConfig() string {
	// should like "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"
	s := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		dbConf.Username, dbConf.Password,
		dbConf.Host, dbConf.Dbname,
		dbConf.Params)
	return s
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

	db, err := gorm.Open(mysql.Open(dbConfig()), &gorm.Config{
		Logger: sqlLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表名单数
			TablePrefix:   "",   // 表前缀
		},
	})
	// 是否开启debug模式
	if dbConf.Debug {
		MyDB = db.Debug()
	} else {
		MyDB = db
	}

	if err != nil {
		// gorm v2 不推荐使用db.Close()方法了
		// 其方法调用隐藏的更深了。
		// 暂时不Close()看看吧
		// _db, _ := MyDB.DB()
		// _db.Close()
		log.Errorf("database connect error: %v \n", err)
		panic("数据库连接失败！")
	}

}
