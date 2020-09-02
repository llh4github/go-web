package middleware

import (
	"strings"
	"testing"

	"gitee.com/llh-gitee/go-web/config"
)

// 测试url分割结果
func TestPath(t *testing.T) {
	path := "/a/b/c"
	sp := strings.Split(path, "/a")
	t.Log(sp[1])
}

// 测试
func TestValidRole(t *testing.T) {
	enf := config.Enforcer
	has, err := enf.HasRoleForUser("Tom", "admin")
	if err != nil {
		t.Log(err)
	}
	t.Log("Has Tom role ? ", has)
}
