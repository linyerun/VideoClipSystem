package db

import (
	"VideoClipSystem/app/global"
	"VideoClipSystem/app/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

var db *gorm.DB

func Init() {
	dsn := strings.Join([]string{global.DbUsername, ":", global.DbPassword, "@tcp(", global.DbHost, ":", global.DbPort, ")/", global.DbName, "?charset=utf8mb4&parseTime=true&loc=Local&timeout=10s&readTimeout=30s&writeTimeout=60s"}, "")
	mySQLInit(dsn)
}

func mySQLInit(dsn string) {
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true, //预编译语句
		Logger:                 newGormLogger(),
	})
	if err != nil {
		utils.Logger().Errorf("open mysql fail,err = %v\n", err)
		panic(err)
	}
}

func newGormLogger() logger.Interface {
	//创建文件夹
	rootDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fileDir := rootDir + "/video_clip_system_logs/gorm_logs"
	if _, err = os.Stat(fileDir); os.IsNotExist(err) {
		//不存在这个文件夹就创建
		err := os.MkdirAll(fileDir, 0666) //可读写，不可执行
		if err != nil {
			panic(err)
		}
	}

	//创建文件
	fileName := "gorm.log"
	filePath := path.Join(fileDir, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}

	return logger.New(
		log.New(file, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,
		},
	)
}
