package service

import (
	"gitee.com/llh-gitee/go-web/model"
	"github.com/sirupsen/logrus"
)

// Role 角色信息服务层
type Role struct {
}

// Add 添加角色信息
func (s *Role) Add(r model.Role) bool {

	r.SetCreatedInfo()
	tx := db.Begin()
	tx.Create(&r)
	if tx.NewRecord(&r) {
		logrus.Error(" Role info insert to db error ")
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}
