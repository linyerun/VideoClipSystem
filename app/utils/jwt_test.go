package utils

import (
	"VideoClipSystem/app/global"
	"fmt"
	"testing"
)

func TestJwt(t *testing.T) {
	global.GoJwtSecret = "tttt"
	global.TokenExpireHour = 2

	token, err := GenerateToken(1, "123456", 0)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("token:", token)

	claims, err := ParseToken(token)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("msg: ", claims.Id, claims.Username, claims.Authority)
}
