package service

import (
	"gitee.com/llh-gitee/go-web/model"
)

// Role 角色信息服务层
type Role struct {
}

// Add 添加角色信息
func (s *Role) Add(r model.Role) bool {

	r.SetCreatedInfo()
	result := db.Create(&r)
	return result.RowsAffected == 1
}
