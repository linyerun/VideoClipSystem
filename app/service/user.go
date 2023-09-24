package service

import (
	"VideoClipSystem/app/db"
	"VideoClipSystem/app/entity"
	"VideoClipSystem/app/utils"
)

func UserLogin(email string, password string) (token string, ok bool) {
	// 获取user
	user, err := db.SelectUserByEmail(email)
	if err != nil {
		utils.Logger().Error(err)
		return
	}
	// 判断密码是否正确
	ok = utils.CheckoutPassword(user.Password, password)
	if !ok {
		return
	}
	// 生成token
	token, err = utils.GenerateToken(user.Id, user.Email, 0)
	if err != nil {
		utils.Logger().Error(err)
		return "", false
	}
	return
}

func UserRegister(email string, password string) (token string, ok bool) {
	// 判断邮箱是否被注册过了
	isInRepoUser := db.IsInRepoUser(email)
	if isInRepoUser {
		return
	}

	// 注册User
	encryptPassword, err := utils.EncryptPassword(password)
	if err != nil {
		return
	}
	id, err := db.CreateUser(entity.NewUser(email, encryptPassword))
	if err != nil {
		utils.Logger().Error(err)
		return
	}

	// 生成token
	token, err = utils.GenerateToken(id, email, 0)
	if err != nil {
		utils.Logger().Error(err)
		return
	}
	return token, true
}
