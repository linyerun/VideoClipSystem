package db

import (
	"VideoClipSystem/app/entity"
	"VideoClipSystem/app/global"
	"fmt"
	"testing"
)

func TestSelectUserByEmail(t *testing.T) {
	global.DbHost = "127.0.0.1"
	global.DbPort = "3306"
	global.DbUsername = "root"
	global.DbPassword = "123456"
	global.DbName = "video_clip_system"
	Init()

	user, err := SelectUserByEmail("2338244917@qq.com")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(user)
}

func TestIsInRepoUser(t *testing.T) {
	global.DbHost = "127.0.0.1"
	global.DbPort = "3306"
	global.DbUsername = "root"
	global.DbPassword = "123456"
	global.DbName = "video_clip_system"
	Init()

	fmt.Println(IsInRepoUser("2338244917@qq.com"))
	fmt.Println(IsInRepoUser("233824491711@qq.com"))
}

func TestCreateUser(t *testing.T) {
	global.DbHost = "127.0.0.1"
	global.DbPort = "3306"
	global.DbUsername = "root"
	global.DbPassword = "123456"
	global.DbName = "video_clip_system"
	Init()
	u := entity.NewUser("3268242396@qq.com", "$2a$10$PtjNmmcgQJU41a0auCMzU.YlU0/usgrbaoFKtT3/BWWHi2uKTRTqO")
	id, err := CreateUser(u)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("id: ", id)
}
