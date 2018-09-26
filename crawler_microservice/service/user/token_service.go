package main

import (
	"GoCrawler/crawler_microservice/service/user/proto"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Authable interface {
	Decode(tokenStr string) (*CustomClaims, error)
	Encode(user *user.User) (string, error)
}

// 定义加盐哈希密码时所用的盐
// 要保证其生成和保存都足够安全
// 比如使用 md5 来生成
var privateKey = []byte("`xs#a_1-!")

// 自定义的 metadata
// 在加密后作为 JWT 的第二部分返回给客户端
type CustomClaims struct {
	User *user.User
	jwt.StandardClaims
}

//
// 将 JWT 字符串解密为 CustomClaims 对象
//
func (srv *TokenService) Decode(tokenStr string) (*CustomClaims, error) {
	t, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return privateKey, nil
	})
	// 解密转换类型并返回
	if claims, ok := t.Claims.(*CustomClaims); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

type TokenService struct {
}

//
// 将 User 用户信息加密为 JWT 字符串
//
func (srv *TokenService) Encode(user *user.User) (string, error) {
	// 三天后过期
	expireTime := time.Now().Add(time.Hour * 24 * 3).Unix()
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			Issuer:    "go.micro.srv.user", // 签发者
			ExpiresAt: expireTime,
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString(privateKey)
}
