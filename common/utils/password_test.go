package utils

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

// 测试加密和匹配
func TestHash(t *testing.T) {
	raw := "password"
	originPwd := []byte(raw)
	hashPwd, _ := bcrypt.GenerateFromPassword(originPwd, bcrypt.DefaultCost)
	t.Log(string(hashPwd))
	err := bcrypt.CompareHashAndPassword(hashPwd, originPwd)
	if err != nil {
		t.Log("密码匹配失败！")
	} else {
		t.Log("密码匹配成功！")
	}
	other := []byte("other")
	err = bcrypt.CompareHashAndPassword(hashPwd, other)
	if err != nil {
		t.Log("密码匹配失败！")
	} else {
		t.Log("密码匹配成功！")
	}
}
