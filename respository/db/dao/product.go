package dao

import (
	"context"
	"gin-mall/respository/db/model"
	"gin-mall/types"
	"gorm.io/gorm"
)

type ProductDao struct {
	*gorm.DB
}

func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{NewDBClient(ctx)}
}

func (dao *ProductDao) ListProductByCondition(categoryID uint, page types.BasePage) (products []*model.Product, err error) {
	err = dao.DB.Model(&model.Product{}).Where("category = ?", categoryID).
		Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).
		Find(&products).Error
	if err != nil {
		return
	}
	return
}

func (dao *ProductDao) CreateProduct(product *model.Product) (err error) {
	err = dao.DB.Create(product).Error
	return
}

// 获取商品的数量
func (dao *ProductDao) CountProductByCondition(categoryID uint) (total int64, err error) {
	err = dao.DB.Model(&model.Product{}).
		Where("category = ?", categoryID).Count(&total).Error
	return
}

func (dao *ProductDao) GetBossById(uId uint) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id = ?", uId).
		First(user).Error
	return
}
