package model

import "gorm.io/gorm"

// 这里的钱呐，密码呐都需要密文存储

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	Email          string
	PasswordDigest string
	NickName       string
	Status         string
	Avatar         string
	Money          string
}
