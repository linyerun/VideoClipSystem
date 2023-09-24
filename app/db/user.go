package db

import (
	"VideoClipSystem/app/entity"
	"VideoClipSystem/app/utils"
)

func SelectUserByEmail(email string) (user *entity.User, err error) {
	user = new(entity.User)
	err = db.Where("email = ?", email).First(user).Error
	return
}

func CreateUser(user *entity.User) (id int64, err error) {
	err = db.Create(user).Error
	return user.Id, err
}

func IsInRepoUser(email string) bool {
	cnt := int64(0)
	// 没有隐式指明表就需要自己指明
	if err := db.Table("user").Select("email").Where("email = ?", email).Count(&cnt).Error; err != nil {
		utils.Logger().Error(err)
	}
	return cnt == 1
}
