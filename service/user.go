package service

import (
	"demo4casbin/model"

	"github.com/sirupsen/logrus"
)

// User 用户服务层
type User struct {
}

// Add 添加用户
func (s *User) Add(u model.User) bool {

	u.SetCreatedInfo()
	u.SetPassowrd(u.Password)
	tx := db.Begin()
	tx.Create(&u)
	if tx.NewRecord(&u) {
		logrus.Error("user info insert to db error ")
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true

}
