package utils

import (
	"math"
	"time"

	"gitee.com/llh-gitee/go-web/common"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

const (
	// Iss jwt签发者
	Iss string = "iss"
	// Sub  jwt所面向的用户
	Sub = "sub"
	// Exp jwt的过期时间，这个过期时间必须要大于签发时间
	Exp = "exp"
	// Iat  jwt的签发时间
	Iat = "iat"
)

// CreateToken 创建token
func CreateToken(uid string) string {
	// jwt里的时间请使用时间戳，可以减少jwt长度和时间转换步骤
	claims := jwt.MapClaims{}
	claims[Sub] = uid
	claims[Iat] = time.Now().Unix()
	claims[Iss] = jwtConf.Iss
	claims[Exp] = time.Now().Add(
		time.Minute * time.Duration(jwtConf.Exp)).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(jwtConf.Secret))
	if err != nil {
		logrus.Error("创建token出错！", err)
		common.ExceptionByCode(common.JwtCreate)
	}
	return token
}

// ParseToken 解析token
//
// 会进行过期验证
func ParseToken(token string) jwt.MapClaims {
	// 此处会进行过期验证
	tk, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtConf.Secret), nil
	})
	if err != nil {
		logrus.Error("解析token出错！", err)
		common.ExceptionByCode(common.JwtParse)
	}
	return tk.Claims.(jwt.MapClaims)
}

// GetSub 获取jwt所面向的用户
func GetSub(token string) string {
	m := ParseToken(token)
	return m[Sub].(string)
}

// GetExp 获取jwt的过期时间
//
// 返回时间戳（秒）
func GetExp(token string) int64 {
	m := ParseToken(token)
	// 解析回来就成float64  -_-|||
	tmp := m[Exp].(float64)
	return int64(tmp * math.Pow10(0))
}
