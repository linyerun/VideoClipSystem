package utils

import (
	"VideoClipSystem/app/global"
	"fmt"
	"testing"
)

func TestEncryptAndCheckoutPassword(t *testing.T) {
	global.PasswordCost = 10

	p := "123456"

	encryptPassword, err := EncryptPassword(p)
	fmt.Println("encryptPassword: ", encryptPassword)
	if err != nil {
		t.Error(err)
	}

	res := CheckoutPassword(encryptPassword, p)
	fmt.Println("res:", res)
}
