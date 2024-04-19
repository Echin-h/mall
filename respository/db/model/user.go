package model

import (
	conf "gin-mall/conf/sql"
	"github.com/CocaineCong/secret"
	"github.com/spf13/cast"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	Email          string
	PasswordDigest string
	NickName       string
	Status         string
	Avatar         string `gorm:"size:1000"`
	Money          string
	Relations      []User `gorm:"many2many:relation;"`
}

const (
	PassWordCost        = 12       // 密码加密难度
	Active       string = "active" // 激活用户
)

func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	u.PasswordDigest = string(bytes)
	return nil
}

// 加密Money，这里是使用的是其他人的一个库
func (u *User) EncryptMoney(key string) (money string, err error) {
	aes, err := secret.NewAesEncrypt(conf.Config.EncryptSecret.MoneySecret, key, "", secret.AesEncrypt128, secret.AesModeTypeCBC)
	if err != nil {
		return
	}
	money = aes.SecretEncrypt(u.Money)
	return
}

// 解密金额
func (u *User) CheckPassword(password string) (money float64, err error) {
	aes, err := secret.NewAesEncrypt(conf.Config.EncryptSecret.MoneySecret, password, "", secret.AesEncrypt128, secret.AesModeTypeCBC)
	if err != nil {
		return
	}
	money = cast.ToFloat64(aes.SecretDecrypt(u.Money))
	return
}
