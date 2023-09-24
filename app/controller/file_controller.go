package controller

import (
	"VideoClipSystem/app/db"
	"VideoClipSystem/app/resp_msg"
	"github.com/gin-gonic/gin"
	"strconv"
)

func clippedVideoFileReturn(c *gin.Context) {
	clippedVideoId, err := strconv.ParseInt(c.Param("clippedVideoId"), 10, 64)
	if err != nil {
		code := resp_msg.FileLoadError
		c.JSON(code, resp_msg.NewRespMsg(code, nil))
		return
	}
	fileOutPath, err := db.SelectClippedVideoFileOutPathById(clippedVideoId)
	if err != nil {
		code := resp_msg.FileLoadError
		c.JSON(code, resp_msg.NewRespMsg(code, nil))
		return
	}
	c.File("." + fileOutPath)
}
