package model

import "demo4casbin/config"

var db = config.MyDB

// 模型注册
func init() {
	db.AutoMigrate(&User{})
}
