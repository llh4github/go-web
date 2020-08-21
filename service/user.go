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
	if err := db.Create(&u); err != nil {
		logrus.Error("user info insert to db error : ", err)
		return false
	}
	return true

}
