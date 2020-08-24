package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
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
func CreateToken(uid string) (string, error) {
	claims := jwt.MapClaims{}
	claims[Sub] = uid
	claims[Iat] = time.Now()
	claims[Exp] = time.Now().Add(time.Minute * time.Duration(jwtConf.Exp))
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(jwtConf.Secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

// ParseToken 解析token
func ParseToken(token string, secret string) (string, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}
	return claim.Claims.(jwt.MapClaims)["uid"].(string), nil
}
