package db

import (
	"VideoClipSystem/app/global"
	"fmt"
	"testing"
)

func TestCreateClippedVideo(t *testing.T) {
	global.DbHost = "127.0.0.1"
	global.DbPort = "3306"
	global.DbUsername = "root"
	global.DbPassword = "123456"
	global.DbName = "video_clip_system"
	Init()

	videoId, err := CreateClippedVideo(1, "111", 0, 1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("videoId:", videoId)
}

func TestFillInClippedVideoMsgById(t *testing.T) {
	global.DbHost = "127.0.0.1"
	global.DbPort = "3306"
	global.DbUsername = "root"
	global.DbPassword = "123456"
	global.DbName = "video_clip_system"
	Init()

	err := FillInClippedVideoMsgById(1, "222", "333", "444")
	if err != nil {
		t.Error(err)
	}
}

func TestSelectClippedVideoFileOutPathById(t *testing.T) {
	global.DbHost = "127.0.0.1"
	global.DbPort = "3306"
	global.DbUsername = "root"
	global.DbPassword = "123456"
	global.DbName = "video_clip_system"
	Init()

	fileOutPath, err := SelectClippedVideoFileOutPathById(1)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("fileOutPath:", fileOutPath)
}
