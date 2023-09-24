package service

import (
	"VideoClipSystem/app/db"
	"VideoClipSystem/app/global"
	"VideoClipSystem/app/processing_center"
	"VideoClipSystem/app/utils"
	"VideoClipSystem/app/video_clip"
	"fmt"
	"os"
	"time"
)

func ClipVideo(userId int64, inUrl, email string, startTimestamp, endTimestamp int64) (clippedVideoId int64, err error) {
	clippedVideoId, err = db.CreateClippedVideo(userId, inUrl, startTimestamp, endTimestamp)
	if err != nil {
		utils.Logger().Error(err)
		return
	}
	// 进行异步任务处理
	processing_center.HandleTask(clipVideoTask(clippedVideoId, inUrl, email, startTimestamp, endTimestamp))
	return
}

func clipVideoTask(clippedVideoId int64, url, email string, startTimestamp, endTimestamp int64) func() {
	return func() {
		// 通过URl获取需要剪辑的视频，对视频进行剪辑
		filePath, err := video_clip.DownloadVideo(url)
		if err != nil {
			utils.Logger().Error(err)
			utils.SendEmail("VideoClipSystem Message", "剪辑过程出现错误: "+err.Error(), email)
			return
		}

		// 删除临时文件操作
		defer func() {
			processing_center.HandleTask(func() {
				for i := 0; i < 10; i++ {
					err := os.Remove(filePath)
					if err == nil {
						break
					}
					utils.Logger().Error(err)
					time.Sleep(time.Millisecond * 100)
				}
			})
		}()

		// 进行剪辑操作
		videoFilePath, err := video_clip.ClipVideo(filePath, startTimestamp, endTimestamp)
		if err != nil {
			utils.Logger().Error(err)
			utils.SendEmail("VideoClipSystem Message", "剪辑过程出现错误: "+err.Error(), email)
			return
		}

		outURL := "http://" + global.HostIp + ":" + global.HostPort + "/file/video/" + fmt.Sprintf("%d", clippedVideoId)

		// 对剪辑信息进行持久化
		for {
			err = db.FillInClippedVideoMsgById(clippedVideoId, outURL, filePath, videoFilePath)
			if err == nil {
				break
			}
			utils.Logger().Error(err)
			time.Sleep(100 * time.Millisecond)
		}

		// 发送邮件告知用户访问的URL
		utils.SendEmail("VideoClipSystem Message", "The New Video URL: "+outURL, email)
	}
}
