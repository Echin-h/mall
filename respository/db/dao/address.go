package dao

import (
	"context"
	"gin-mall/respository/db/model"
	"gorm.io/gorm"
)

type AddressDao struct {
	*gorm.DB
}

func NewAddressDao(ctx context.Context) *AddressDao {
	return &AddressDao{NewDBClient(ctx)}
}

func NewAddressDaoByDB(db *gorm.DB) *AddressDao {
	return &AddressDao{db}
}

func (dao *AddressDao) GetAddressByID(aid, uid uint) (address *model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).
		Where("id = ? AND user_id = ?", aid, uid).
		First(&address).Error
	return
}
