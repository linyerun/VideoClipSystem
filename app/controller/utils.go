package controller

import (
	"VideoClipSystem/app/resp_msg"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func readJson(c *gin.Context, data any) (ok bool) {
	body := c.Request.Body
	defer func() { _ = body.Close() }()

	// 读取Json数据
	if err := json.NewDecoder(body).Decode(data); err != nil {
		code := resp_msg.InvalidParams
		c.JSON(code, resp_msg.NewRespMsg(code, nil))
		return false
	}
	return true
}
