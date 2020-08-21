package service

import (
	"demo4casbin/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	// 获取数据库连接
	db = config.MyDB
}
