package config

import (
	"VideoClipSystem/app/global"
	"VideoClipSystem/app/utils"
	"gopkg.in/ini.v1"
)

func Init() {
	//加载config.ini
	file, err := ini.Load("app/config/config.ini")
	if err != nil {
		utils.Logger().Errorf("open config.ini file fail,err = %v\n", err)
		panic(err)
	}
	loadMysqlConfig(file)
	loadJwtConfig(file)
	loadPwdConfig(file)
	loadProcessingCenterConfig(file)
	loadAuthCodeConfig(file)
	loadSmtpConfig(file)
	loadHostConfig(file)
}

func loadMysqlConfig(file *ini.File) {
	global.DbHost = file.Section("mysql").Key("DbHost").String()
	global.DbPort = file.Section("mysql").Key("DbPort").String()
	global.DbUsername = file.Section("mysql").Key("DbUser").String()
	global.DbPassword = file.Section("mysql").Key("DbPassword").String()
	global.DbName = file.Section("mysql").Key("DbName").String()
}

func loadJwtConfig(file *ini.File) {
	global.GoJwtSecret = file.Section("jwt").Key("GoJwtSecret").String()
	var err error
	if global.TokenExpireHour, err = file.Section("jwt").Key("TokenExpireHour").Int(); err != nil {
		panic(err)
	}
}

func loadPwdConfig(file *ini.File) {
	var err error
	global.PasswordCost, err = file.Section("pwd").Key("PasswordCost").Int()
	if err != nil {
		panic(err)
	}
}

func loadProcessingCenterConfig(file *ini.File) {
	var err error
	global.ProcessingCenterPoolSize, err = file.Section("go_pool").Key("ProcessingCenterPoolSize").Uint()
	if err != nil {
		panic(err)
	}
	global.ProcessingCenterChanLen, err = file.Section("go_pool").Key("ProcessingCenterChanLen").Uint()
	if err != nil {
		panic(err)
	}
}

func loadAuthCodeConfig(file *ini.File) {
	var err error
	global.ExpirationSecond, err = file.Section("auto_code").Key("ExpirationSecond").Uint()
	if err != nil {
		panic(err)
	}
}

func loadSmtpConfig(file *ini.File) {
	global.SmtpServer = file.Section("smtp").Key("SmtpServer").String()
	global.SmtpPort = file.Section("smtp").Key("SmtpPort").String()
	global.SmtpUsername = file.Section("smtp").Key("SmtpUsername").String()
	global.SmtpPassword = file.Section("smtp").Key("SmtpPassword").String()
}

func loadHostConfig(file *ini.File) {
	global.HostIp = file.Section("host").Key("ip").String()
	global.HostPort = file.Section("host").Key("port").String()
}
