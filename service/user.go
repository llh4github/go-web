package service

import (
	"strconv"

	"gitee.com/llh-gitee/go-web/common"
	"gitee.com/llh-gitee/go-web/common/utils"
	"gitee.com/llh-gitee/go-web/model"

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

// Login 用户登录
func (s *User) Login(u model.User) string {
	var user model.User
	db.Where("username = ? and remove_flag = false", u.Username).First(&user)
	if user.ID < 1 {
		common.ExceptionByCode(common.PwdError)
	}
	if utils.MatchPassword(u.Password, user.Password) {
		return utils.CreateToken(strconv.Itoa(user.ID))
	}
	common.ExceptionByCode(common.PwdError)
	return ""
}
