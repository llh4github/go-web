package service

import (
	"testing"

	"gitee.com/llh-gitee/go-web/model"
)

// 测试
func TestAdd(t *testing.T) {
	m := model.Role{
		RoleName: "admin1",
		Remark:   "admin2",
	}
	s := Role{}

	s.Add(m)
}
