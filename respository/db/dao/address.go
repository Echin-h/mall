package dao

import (
	"context"
	"gin-mall/respository/db/model"
	"gin-mall/types"
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

func (dao *AddressDao) CreateAddress(address *model.Address) (err error) {
	err = dao.DB.Create(address).Error
	return
}

func (dao *AddressDao) ListAddress(uid uint, req *types.AddressListReq) (addresses []*model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).
		Where("user_id = ?", uid).
		Offset((req.PageNum - 1) * req.PageSize).
		Find(&addresses).Error
	return
}

func (dao *AddressDao) UpdateAddressByID(uid, aid uint, name, phone, address string) (err error) {
	err = dao.DB.Model(&model.Address{}).
		Where("id = ? AND user_id = ?", aid, uid).
		Updates(map[string]interface{}{"name": name, "phone": phone, "address": address}).Error

	return
}

func (dao *AddressDao) DeleteAddressByID(uid, aid uint) (err error) {
	err = dao.DB.Model(&model.Address{}).
		Where("id = ? AND user_id = ?", aid, uid).
		Delete(&model.Address{}).Error
	return
}
