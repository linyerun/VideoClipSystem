package controller

import (
	"VideoClipSystem/app/cache"
	"VideoClipSystem/app/controller/dto"
	"VideoClipSystem/app/global"
	"VideoClipSystem/app/processing_center"
	"VideoClipSystem/app/resp_msg"
	"VideoClipSystem/app/service"
	"VideoClipSystem/app/utils"
	"github.com/gin-gonic/gin"
	"time"
)

// Login 用户登录接口
//@tags 用户模块
//@summary 登录
//@description 已注册用户可通过登录进入app
//@accept json
//@produce json
//@Param data body dto.UserDto true "登录所需信息"
//@success 200 {object} resp_msg.RespMsg
//@failure 400 {object} resp_msg.RespMsg
//@router /user/login/ [post]
func login(c *gin.Context) {
	// 读取Json数据
	var userDto dto.UserDto
	if ok := readJson(c, &userDto); !ok {
		return
	}

	// 校验账号
	if !utils.VerifyEmailFormat(userDto.Email) {
		code := resp_msg.ErrorEmailFormat
		c.JSON(code, resp_msg.NewRespMsg(code, nil))
		return
	}

	// 校验密码长度
	if len(userDto.Password) > 50 {
		code := resp_msg.ErrorPasswordLen
		c.JSON(code, resp_msg.NewRespMsg(code, nil))
		return
	}

	// 生成token
	token, ok := service.UserLogin(userDto.Email, userDto.Password)
	if !ok { // 登录失败
		code := resp_msg.ErrorLogin
		c.JSON(code, resp_msg.NewRespMsg(code, nil))
		return
	}

	// 登录成功
	code := resp_msg.SUCCESS
	c.JSON(code, resp_msg.NewRespMsg(code, gin.H{"token": token}))
}

// Register 用户注册接口
//@tags 用户模块
//@summary 注册
//@description 用户注册入口
//@accept json
//@produce json
//@Param data body dto.RegisterDto true "注册所需信息"
//@success 200 {object} resp_msg.RespMsg
//@failure 400 {object} resp_msg.RespMsg
//@router /user/register/ [post]
func register(c *gin.Context) {
	// 读取Json数据
	var registerDto dto.RegisterDto
	if ok := readJson(c, &registerDto); !ok {
		return
	}

	// 校验账号
	if !utils.VerifyEmailFormat(registerDto.Email) {
		code := resp_msg.ErrorEmailFormat
		c.JSON(code, resp_msg.NewRespMsg(code, nil))
		return
	}

	// 校验密码长度
	if len(registerDto.Password) > 50 {
		code := resp_msg.ErrorPasswordLen
		c.JSON(code, resp_msg.NewRespMsg(code, nil))
		return
	}

	// 检验验证码是否过期，是否存在
	val, ok := cache.GetTimestampByAutoCode(registerDto.AuthCode)
	processing_center.HandleTask(func() { // 异步删除AuthCode
		cache.DeleteAutoCodeByEmail(registerDto.Email)
		cache.DeleteTimestampByAutoCode(registerDto.AuthCode)
	})
	cur := time.Now().UnixMilli()
	if !ok || val+int64(global.ExpirationSecond)*1000 < cur {
		code := resp_msg.ErrorAuthCodeTimeout
		c.JSON(code, resp_msg.NewRespMsg(code, nil))
		return
	}

	// 注册操作, 返回token
	token, ok := service.UserRegister(registerDto.Email, registerDto.Password)
	if !ok { // 注册失败
		code := resp_msg.ErrorRegister
		c.JSON(code, resp_msg.NewRespMsg(code, nil))
		return
	}

	// 注册成功
	code := resp_msg.SUCCESS
	c.JSON(code, resp_msg.NewRespMsg(code, gin.H{"token": token}))
}
