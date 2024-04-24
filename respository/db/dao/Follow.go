package dao

import (
	"context"
	"errors"
	"gin-mall/respository/db/model"
	"gorm.io/gorm"
)

type FollowDao struct {
	*gorm.DB
}

func NewFollowDao(ctx context.Context) *FollowDao {
	return &FollowDao{NewDBClient(ctx)}
}

func (dao *FollowDao) Follow(uId uint, FollowId uint) (err error) {
	var count int64
	dao.DB.Model(&model.Follow{}).Where("user_id = ? AND follow_id = ?", uId, FollowId).Count(&count)
	if count > 0 {
		err = errors.New("你已经关注了")
		return
	}

	var f model.Follow
	f.FollowId = FollowId
	f.UserId = uId

	err = dao.DB.Model(&model.Follow{}).Create(&f).Error
	return
}

func (dao *FollowDao) UnFollow(uid uint, unfollowId uint) (err error) {
	var count int64
	dao.DB.Model(&model.Follow{}).Where("user_id = ? AND follow_id = ?", uid, unfollowId).Count(&count)
	if count == 0 {
		err = errors.New("你本身没关注")
		return
	}

	var f model.Follow
	err = dao.DB.Model(&model.Follow{}).Where("user_id = ? AND follow_id = ?", uid, unfollowId).Delete(&f).Error

	return
}
