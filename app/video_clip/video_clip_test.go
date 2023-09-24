package video_clip

import (
	"fmt"
	"testing"
	"time"
)

func TestDownloadVideo(t *testing.T) {
	videoPath, err := DownloadVideo("https://douyin-video-picture-lyr.oss-cn-shenzhen.aliyuncs.com/video/2022-17-07/1658037158695188800.mp4")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(videoPath)
}

func TestClipVideo(t *testing.T) {
	vp := "E:\\视频录制保存点\\必剪视频保存点\\2022-03-18-17-52-34.mp4"
	pre := time.Now()
	newVideoPath, err := ClipVideo(vp, 0, 59*60+8)
	fmt.Println("运行时间:", time.Now().Sub(pre).Milliseconds())
	if err != nil {
		t.Error(err)
	}
	fmt.Println("VideoURL:", newVideoPath)
}

func TestClipVideoByUrl(t *testing.T) {
	videoPath, err := DownloadVideo("https://douyin-video-picture-lyr.oss-cn-shenzhen.aliyuncs.com/video/2022-17-07/1658037158695188800.mp4")
	if err != nil {
		t.Error(err)
	}
	videoFilePath, err := ClipVideo(videoPath, 0, 1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("clip video file path:", videoFilePath)
}
