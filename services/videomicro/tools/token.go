package tools

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	Id int64 `json:"id"`
	jwt.RegisteredClaims
}

var JWTSecret = []byte("defaulthmacsamplesecret") // need set while init
// ref : https://github.com/golang-jwt/jwt

func GenerateToken(userid int64) string {
	claims := CustomClaims{
		userid,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(168 * time.Hour)), // 一周过期时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(JWTSecret)
	return tokenString
}

// Token鉴权
func ValidateToken(tokenString string) (valid bool, userid int64, err error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTSecret, nil
	})
	if err != nil {
		return false, 0, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
			return true, claims.Id, nil
		}
	}
	return false, 0, err
}
