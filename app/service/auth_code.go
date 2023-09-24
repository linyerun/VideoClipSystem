package service

import (
	"VideoClipSystem/app/cache"
	"VideoClipSystem/app/global"
	"VideoClipSystem/app/processing_center"
	"VideoClipSystem/app/utils"
	"fmt"
	"time"
)

func SendAuthCode(email string) {
	//验证码存在
	if authCode, ok := cache.GetAutoCodeByEmail(email); ok {
		timestamp, ok := cache.GetTimestampByAutoCode(authCode)
		// 验证码过期了
		if ok && timestamp+int64(global.ExpirationSecond)*1000 < time.Now().UnixMilli() {
			processing_center.HandleTask(func() {
				cache.DeleteAutoCodeByEmail(email)
				cache.DeleteTimestampByAutoCode(authCode)
			})
		} else { // 验证码没过期
			return
		}
	}

	// 生成验证码
	var authCode string
	for i := 6; ; i++ {
		authCode = utils.GenerateAuthCode(i)
		_, ok := cache.GetTimestampByAutoCode(authCode)
		if !ok {
			break
		}
	}
	// 保存验证码
	cache.AddAutoCodeTimestamp(authCode, time.Now().UnixMilli())
	cache.AddEmailAuthCode(email, authCode)
	// 使用邮件发送验证码
	content := fmt.Sprintln("您的注册验证码:", authCode, "。验证码将在", global.ExpirationSecond, "秒后过期。")
	utils.SendEmail("VideoCliSystem Message", content, email)
}
