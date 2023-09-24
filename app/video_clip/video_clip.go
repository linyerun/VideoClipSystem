package video_clip

import (
	"VideoClipSystem/app/utils"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func DownloadVideo(url string) (string, error) {
	//1. 读取响应报文中的视频
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			utils.Logger().Error(err)
		}
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		utils.Logger().Error(err)
		return "", err
	}

	// 2. 创建文件夹
	fileDir := getProjectRootPath() + "/video_clip_system_temp_files"
	createDirIfNo(fileDir)

	// 3. 将视频文件保存到临时文件夹中
	fileName := fmt.Sprintf("video_%d", time.Now().UnixMicro())
	filePath := filepath.Join(fileDir, fileName)
	if err = ioutil.WriteFile(filePath, body, os.ModePerm); err != nil {
		utils.Logger().Error(err)
		return "", err
	}

	return filePath, nil
}

func ClipVideo(videoPath string, startTime int64, endTime int64) (string, error) {
	// 生成文件夹
	outputFileDir := getProjectRootPath() + "/video_clip_system_files"
	createDirIfNo(outputFileDir)

	// 生成输出文件路径
	ext := filepath.Ext(videoPath) // 获取文件拓展名
	if ext == "" {
		ext = ".mp4"
	}
	outputFileName := fmt.Sprintf("file_%d_%d_%d%s", time.Now().UnixMicro(), startTime, endTime, ext)
	filePath := filepath.Join(outputFileDir, outputFileName)

	// 使用FFmpeg进行视频剪辑
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-ss", fmt.Sprintf("%d", startTime), "-to", fmt.Sprintf("%d", endTime), "-c", "copy", "-y", filePath)
	utils.Logger().Println("ffmpeg clip cmd:", cmd.String())

	var outputBuf bytes.Buffer
	var errorBuf bytes.Buffer
	cmd.Stdout = &outputBuf
	cmd.Stderr = &errorBuf

	if err := cmd.Run(); err != nil {
		utils.Logger().Errorln("FFmpegErrorBuf:", errorBuf.String())
		utils.Logger().Errorln("FFmpegError:", err)
		return "", err
	}

	return "/video_clip_system_files/" + outputFileName, nil
}
