package main

import (
	"VideoClipSystem/app/cache"
	"VideoClipSystem/app/config"
	"VideoClipSystem/app/controller"
	"VideoClipSystem/app/db"
	"VideoClipSystem/app/processing_center"
	_ "VideoClipSystem/docs" // 用于生成swagger交互文档
)

// @title 视频剪辑项目接口文档
// @version 1.0
// @description author: 林叶润
// @host 127.0.0.1:8888
// @BasePath /
func main() {
	// 初始化全局参数
	config.Init()

	// 初始胡处理中心
	processing_center.Init()

	// 初始化本地缓存模块GeeCache
	cache.Init()

	// 初始化Gorm
	db.Init()

	// Gin必须放在最后进行初始化
	controller.InitGin()
}
