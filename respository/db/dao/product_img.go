package dao

import (
	"context"
	"gin-mall/respository/db/model"
	"gorm.io/gorm"
)

type ProductImgDao struct {
	*gorm.DB
}

func NewProductImgDao(ctx context.Context) *ProductImgDao {
	return &ProductImgDao{NewDBClient(ctx)}
}

func NewProductImgDaoByDB(db *gorm.DB) *ProductImgDao {
	return &ProductImgDao{db}

}

func (dao *ProductImgDao) CreateProductImg(productImg *model.ProductImg) (err error) {
	err = dao.DB.Model(&model.ProductImg{}).Create(&productImg).Error

	return
}

func (dao *ProductImgDao) ListProductImgByProductId(pId uint) (r []*model.ProductImg, err error) {
	err = dao.DB.Model(&model.ProductImg{}).Where("product_id=?", pId).Find(&r).Error

	return
}
