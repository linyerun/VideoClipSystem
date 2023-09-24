package entity

type User struct {
	Id       int64  `gorm:"column:id"`       //用户ID
	Email    string `gorm:"column:email"`    //登录账号
	Password string `gorm:"column:password"` //登录密码(加密的)
}

func NewUser(email, password string) *User {
	return &User{
		Email:    email,
		Password: password,
	}
}

// TableName 会将 User 的表名重写为 `user`
func (User) TableName() string {
	return "user"
}
