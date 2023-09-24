package video_clip

import (
	"VideoClipSystem/app/utils"
	"os"
)

func getProjectRootPath() string {
	rootDir, err := os.Getwd()
	if err != nil {
		utils.Logger().Error(err)
		return ""
	}
	return rootDir
}

func createDirIfNo(filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) { //不存在这个文件夹就创建
		if err := os.MkdirAll(filePath, 0777); err != nil {
			utils.Logger().Error(err)
		}
	}
}
