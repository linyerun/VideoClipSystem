package utils

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path"
	"sync"
	"time"
)

var myLog *logrus.Logger
var once sync.Once

func Logger() *logrus.Logger {
	once.Do(
		func() {
			myLog = logrus.New()              //创建logrus
			myLog.SetLevel(logrus.DebugLevel) //设置日志级别
			myLog.SetFormatter(               //设置时间格式
				&logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05"},
			)
			output, err := getOutputFile()
			if err != nil {
				panic(err)
			}
			myLog.SetOutput(output)    // 设置输入文件
			myLog.AddHook(new(myHook)) //添加钩子
		},
	)
	return myLog
}

func getOutputFile() (*os.File, error) {
	//获取绝对路径
	rootDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	fileDir := rootDir + "/video_clip_system_logs/global_logs"
	_, err = os.Stat(fileDir) //判断这个文件夹是否存在
	if os.IsNotExist(err) {
		//不存在这个文件夹就创建
		err := os.MkdirAll(fileDir, 0666) //可读写，不可执行
		if err != nil {
			return nil, err
		}
	}

	//目录存在，那就直接执行下面的就行了
	//开始创建文件
	fileName := time.Now().Format("2006-01-02") + ".log"
	filePath := path.Join(fileDir, fileName)
	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		//文件不存在就创建文件
		//Create采用模式0666（任何人都可读写，不可执行）创建一个名为name的文件，如果文件已存在会截断它（为空文件）
		file, err := os.Create(filePath)
		if err != nil {
			return nil, err
		}
		return file, nil
	}

	return os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
}

type myHook struct {
}

func (m *myHook) Levels() []logrus.Level { // 设置日志级别
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
		logrus.TraceLevel,
	}
}

func (m *myHook) Fire(entry *logrus.Entry) error { // 打印日志信息到控制台
	log.Println(entry.Message)
	return nil
}
