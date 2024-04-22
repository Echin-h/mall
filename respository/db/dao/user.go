package dao

import (
	"context"
	"gin-mall/respository/db/model"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

// 创建用户
func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Model(&model.User{}).Create(user).Error
}

func (dao *UserDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.User{}).Where("user_name = ?", userName).Count(&count).Error
	if count == 0 {
		return user, false, err
	}
	err = dao.DB.Model(&model.User{}).Where("user_name = ?", userName).First(&user).Error
	if err != nil {
		return user, false, err
	}
	return user, true, nil
}

func (dao *UserDao) GetUserById(userId uint) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id = ?", userId).
		First(&user).Error
	return
}

func (dao *UserDao) UpdateUserById(uid uint, user *model.User) (err error) {
	return dao.DB.Model(&model.User{}).Where("id = ?", uid).
		Updates(&user).Error
}
