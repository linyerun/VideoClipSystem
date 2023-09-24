package controller

import (
	"VideoClipSystem/app/controller/dto"
	"VideoClipSystem/app/resp_msg"
	"VideoClipSystem/app/service"
	"github.com/gin-gonic/gin"
)

// ClipVideo 剪辑视频
//@tags 视频模块
//@summary 视频剪辑
//@description 用于剪辑视频
//@accept */*
//@produce */*
//@Param token header string true "token值"
//@Param data body dto.VideoDto true "剪辑一段视频所需信息"
//@success 200 {object} resp_msg.RespMsg
//@failure 400 {object} resp_msg.RespMsg
//@router /video/clip/ [post]
func clipVideo(c *gin.Context) {
	// 读取信息
	var videoDto dto.VideoDto
	ok := readJson(c, &videoDto)
	if !ok {
		return
	}

	// 校验时间是否有误
	if videoDto.StartTime < 0 || videoDto.EndTime < 0 || videoDto.StartTime > videoDto.EndTime {
		code := resp_msg.InvalidParams
		c.JSON(code, resp_msg.NewRespMsg(code, nil))
		return
	}

	// 校验URL，避免恶意的URL
	if len(videoDto.Url) > 500 {
		code := resp_msg.ErrorURLLen
		c.JSON(code, resp_msg.NewRespMsg(code, nil))
		return
	}

	// 开始业务处理
	userId, _ := c.Get("userId")
	userEmail, _ := c.Get("userEmail")
	clippedVideoId, err := service.ClipVideo(userId.(int64), videoDto.Url, userEmail.(string), videoDto.StartTime, videoDto.EndTime)
	if err != nil {
		code := resp_msg.SystemERROR
		c.JSON(code, resp_msg.NewRespMsg(code, nil))
		return
	}
	// 返回信息
	code := resp_msg.SUCCESS
	c.JSON(code, resp_msg.NewRespMsg(code, gin.H{"clipped_video_id": clippedVideoId, "msg": "视频真正剪辑中, 请留意后续邮箱信息!"}))
}

func clipVideoProgress(c *gin.Context) {

}
