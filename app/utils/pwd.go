package utils

import (
	"VideoClipSystem/app/global"
	"golang.org/x/crypto/bcrypt"
)

// EncryptPassword  密码加密
func EncryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), global.PasswordCost)
	if err != nil {
		Logger().Errorln(err)
		return "", err
	}
	return string(bytes), nil
}

//CheckoutPassword 校验密码
func CheckoutPassword(encryptionPassword, InputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encryptionPassword), []byte(InputPassword))
	if err != nil {
		Logger().Error(err)
		return false
	}
	return true
}
