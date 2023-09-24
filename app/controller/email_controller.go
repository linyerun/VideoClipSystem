package controller

import (
	"VideoClipSystem/app/db"
	"VideoClipSystem/app/processing_center"
	"VideoClipSystem/app/resp_msg"
	"VideoClipSystem/app/service"
	"VideoClipSystem/app/utils"
	"github.com/gin-gonic/gin"
)

// SendAuthCodeByEmail 验证码发送接口
//@tags 邮箱模块
//@summary 发送验证码
//@description 发送验证码
//@accept json
//@produce json
//@param email query string true "用户邮箱"
//@success 200 {object} resp_msg.RespMsg
//@failure 400 {object} resp_msg.RespMsg
//@router /email/authCode/ [post]
func sendAuthCodeByEmail(c *gin.Context) {
	email := c.Query("email")

	// 校验邮箱
	if !utils.VerifyEmailFormat(email) {
		code := resp_msg.ErrorEmailFormat
		c.JSON(code, resp_msg.NewRespMsg(code, nil))
		return
	}

	if db.IsInRepoUser(email) { //又被注册
		code := resp_msg.BeRegistered
		c.JSON(code, resp_msg.NewRespMsg(code, nil))
		return
	}

	// 异步发送验证码
	processing_center.HandleTask(func() {
		service.SendAuthCode(email)
	})

	code := resp_msg.SUCCESS
	c.JSON(code, resp_msg.NewRespMsg(code, nil))
}
