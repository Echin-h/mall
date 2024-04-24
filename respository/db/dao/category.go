package dao

import (
	"context"
	"gin-mall/respository/db/model"
	"gorm.io/gorm"
)

type CategoryDao struct {
	*gorm.DB
}

func NewCategoryDao(ctx context.Context) *CategoryDao {
	return &CategoryDao{NewDBClient(ctx)}
}

func NewCategoryDaoByDB(db *gorm.DB) *CategoryDao {
	return &CategoryDao{db}
}

func (dao *CategoryDao) ListCategory() (r []*model.Category, err error) {
	err = dao.DB.Model(&model.Category{}).Find(&r).Error
	return
}
