package entity

type ClippedVideo struct {
	Id             int64  `gorm:"column:id"`
	UserId         int64  `gorm:"column:user_id"`
	InUrl          string `gorm:"column:in_url"`
	OutUrl         string `gorm:"column:out_url"`
	StartTimestamp int64  `gorm:"column:start_timestamp"`
	EndTimestamp   int64  `gorm:"column:end_timestamp"`
	FileIn         string `gorm:"column:file_in"`
	FileOut        string `gorm:"column:file_out"`
}

func NewClippedVideo(userId int64, inUrl string, startTimestamp, endTimestamp int64) *ClippedVideo {
	return &ClippedVideo{
		UserId:         userId,
		InUrl:          inUrl,
		StartTimestamp: startTimestamp,
		EndTimestamp:   endTimestamp,
	}
}

func (ClippedVideo) TableName() string {
	return "clipped_video"
}
