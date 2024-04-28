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

func (dao *ProductDao) ShowProduct(id uint) (*model.Product, error) {
	var product model.Product
	err := dao.DB.Model(&model.Product{}).Where("id = ?", id).
		First(&product).Error
	return &product, err
}

func (dao *ProductDao) UpdateProduct(id uint, product *model.Product) (err error) {
	err = dao.DB.Model(&model.Product{}).Where("id = ?", id).
		Updates(product).Error

	return
}

func (dao *ProductDao) DeleteProduct(rid uint, bid uint) (err error) {
	return dao.DB.Model(&model.Product{}).
		Where("id = ? and boss_id = ?", rid, bid).
		Delete(&model.Product{}).Error
}

func (dao *ProductDao) SearchProduct(keyword string, page types.BasePage) (products []*model.Product, err error) {
	err = dao.DB.Model(&model.Product{}).Where("name like ?", "%"+keyword+"%").
		Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).
		Find(&products).Error
	if err != nil {
		return
	}
	return
}

func (dao *ProductDao) CountSearchProduct(keyword string) (total int64, err error) {
	err = dao.DB.Model(&model.Product{}).
		Where("name like ?", "%"+keyword+"%").Count(&total).Error
	return
}

func (dao *ProductDao) GetProductById(id uint) (product *model.Product, err error) {
	err = dao.DB.Model(&model.Product{}).Where("id = ?", id).
		First(&product).Error
	return
}

func (dao *ProductDao) UpdateProductById(id uint, product *model.Product) (err error) {
	return dao.DB.Model(&model.Product{}).Where("id = ?", id).
		Updates(product).Error
}
