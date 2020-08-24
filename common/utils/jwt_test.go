package utils

import "testing"

var uid = "123"

// 测试
func TestEncode(t *testing.T) {
	token := CreateToken(uid)
	t.Log(token)
	_uid := ParseToken(token)
	t.Log(_uid)
}

// 测试过期
func TestValid(t *testing.T) {
	const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTgyNTY5NzAsImlhdCI6MTU5ODI1NjA3MCwiaXNzIjoibGxoIiwic3ViIjoiMTIzIn0.abP6Nv2VjPi_Fa2u69EWx99FmX70_JVupsIh5Ua0x38"
	t.Log(GetExp(token))
}

// 测试获取过期时间
func TestExp(t *testing.T) {
	token := CreateToken(uid)
	t.Log(GetExp(token))
}
