package service

import (
	"gitee.com/llh-gitee/go-web/config"

	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	// 获取数据库连接
	db = config.MyDB
}
