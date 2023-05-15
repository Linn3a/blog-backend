package controller

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type CustomClaims struct {
	// 可根据需要自行添加字段
	UserID               uint   `json:"user_id"`
	Username             string `json:"username"`
	Avatar               string `json:"avatar"`
	jwt.RegisteredClaims        // 内嵌标准的声明
}

// GenToken 生成JWT
var mySigningkey = []byte("secret")

func GenToken(userId uint, username string, avatar string) (string, error) {
	// 创建一个我们自己声明的数据
	claims := CustomClaims{
		userId,
		username, // 自定义字段
		avatar,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)), // 定义过期时间
			Issuer:    "cos",                                              // 签发人
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成签名字符串
	return token.SignedString(mySigningkey)
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析token
	var mc = new(CustomClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return mySigningkey, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
