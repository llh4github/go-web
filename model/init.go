package model

import "gitee.com/llh-gitee/go-web/config"

var db = config.MyDB

// 模型注册
func init() {
	db.AutoMigrate(&User{})
}
