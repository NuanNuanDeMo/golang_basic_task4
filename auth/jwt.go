package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const SECRETKEY = "zswtest122222"

type User struct {
	Uid                  uint
	Username             string `json:"username"`
	jwt.RegisteredClaims        // v5版本新加的方法
}

// 生成jwt
func GenerateJWT(username string, uid uint) (string, error) {
	claims := User{
		uid,
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 过期时间24小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     // 生效时间
		},
	}
	// 使用HS256签名算法
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := t.SignedString([]byte(SECRETKEY))
	return s, err
}

// 解析jwt
func ParseJwt(tokenstring string) (*User, error) {
	t, err := jwt.ParseWithClaims(tokenstring, &User{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRETKEY), nil
	})
	if claims, ok := t.Claims.(*User); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
