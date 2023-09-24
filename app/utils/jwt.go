package utils

import (
	"VideoClipSystem/app/global"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TokenClaims struct {
	Id        int64
	Username  string
	Authority int
	jwt.StandardClaims
}

//GenerateToken 签发用户Token
func GenerateToken(id int64, username string, authority int) (string, error) {
	expireTime := time.Now().Add(time.Duration(global.TokenExpireHour) * time.Hour)
	claim := TokenClaims{
		Id:        id,
		Username:  username,
		Authority: authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.UnixMilli(),
			Issuer:    "video-clip-system",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := tokenClaims.SignedString([]byte(global.GoJwtSecret))
	return token, err
}

//ParseToken 验证用户token
func ParseToken(token string) (*TokenClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.GoJwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims == nil {
		return nil, errors.New("your TokenClaims is nil")
	}
	if claim, ok := tokenClaims.Claims.(*TokenClaims); ok && tokenClaims.Valid {
		return claim, nil
	}
	return nil, errors.New("your TokenClaims.TokenClaims can not be utils.TokenClaims")
}
