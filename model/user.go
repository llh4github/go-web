package model

import (
	"demo4casbin/common/utils"

	"github.com/sirupsen/logrus"
)

// User 用户模型
type User struct {
	BasicModel
	Username string `json:"username"`
	Password string `json:"password"`
}

// SetPassowrd 设置密码
func (u *User) SetPassowrd(raw string) {
	pwd, err := utils.HashPassword(raw)
	if err != nil {
		logrus.Error("用户密码加密出错！", err)
		panic("用户密码加密出错！")
	}
	u.Password = pwd
}

// Add 新增用户信息
// 返回操作结果。是否成功。
//
// 根据主键是否为空来判断是否插入成功的
func (u *User) Add() bool {
	u.SetPassowrd(u.Password)
	u.SetCreatedInfo()
	if err := db.Create(u).Error; err != nil {
		logrus.Error("insert error : ", err.Error())
		return false
	}
	return true
}
