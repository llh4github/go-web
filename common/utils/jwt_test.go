package utils

import "testing"

var uid = "123"
var secret = "llh"

// 测试
func TestEncode(t *testing.T) {
	token, _ := CreateToken(uid, secret)
	t.Log(token)
	_uid, _ := ParseToken(token, secret)
	t.Log(_uid)
}
