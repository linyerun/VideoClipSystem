package db

import "VideoClipSystem/app/entity"

func CreateClippedVideo(userId int64, inUrl string, startTimestamp, endTimestamp int64) (clippedVideoId int64, err error) {
	clippedVideo := entity.NewClippedVideo(userId, inUrl, startTimestamp, endTimestamp)
	err = db.Create(clippedVideo).Error
	return clippedVideo.Id, err
}

func FillInClippedVideoMsgById(cvId int64, outUrl, fileInPath, fileOutPath string) (err error) {
	ev := new(entity.ClippedVideo)
	ev.Id = cvId
	ev.OutUrl = outUrl
	ev.FileIn = fileInPath
	ev.FileOut = fileOutPath
	return db.Updates(ev).Error
}

func SelectClippedVideoFileOutPathById(cvId int64) (fileOutPath string, err error) {
	err = db.Table("clipped_video").Select("file_out").Where("id = ?", cvId).Scan(&fileOutPath).Error
	return
}
