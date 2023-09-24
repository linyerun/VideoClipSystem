package utils

import "testing"

func TestSendEmail(t *testing.T) {
	SendEmail("验证码信息", "123321", "3268242396@qq.com")
}
