package utils

import (
	"VideoClipSystem/app/global"
	"fmt"
	"net/smtp"
	"regexp"
)

// VerifyEmailFormat 校验邮箱格式
func VerifyEmailFormat(email string) bool {
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func SendEmail(title string, content string, emails ...string) {
	// 配置 SMTP 认证信息
	auth := smtp.PlainAuth("", global.SmtpUsername, global.SmtpPassword, global.SmtpServer)

	// 配置信息
	message := "From: " + global.SmtpUsername + "\n" +
		"Subject: " + title + "\n\n" + content

	// 发送邮件
	err := smtp.SendMail(global.SmtpServer+":"+global.SmtpPort, auth, global.SmtpUsername, emails, []byte(message))
	if err != nil {
		fmt.Println("Error sending email:", err)
	} else {
		fmt.Println("Email sent successfully")
	}
}
