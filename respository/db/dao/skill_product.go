package dao

import (
	"context"
	"gin-mall/consts"
	"gin-mall/respository/db/model"
	"gorm.io/gorm"
)

type SkillProductDao struct {
	*gorm.DB
}

func NewSkillProductDao(ctx context.Context) *SkillProductDao {
	return &SkillProductDao{NewDBClient(ctx)}
}

func NewSkillProductDaoByDB(db *gorm.DB) *SkillProductDao {
	return &SkillProductDao{db}
}

func (dao *SkillProductDao) BatchCreate(batch []*model.SkillProduct) error {
	return dao.DB.Model(&model.SkillProduct{}).
		CreateInBatches(&batch, consts.ProductBatchCreate).Error
}

func (dao *SkillProductDao) CreateByList(in []*model.SkillProduct) error {
	return dao.Model(&model.SkillProduct{}).Create(&in).Error
}

func (dao *SkillProductDao) ListSkillGoods() (resp []*model.SkillProduct, err error) {
	err = dao.Model(&model.SkillProduct{}).
		Where("num > 0").Find(&resp).Error

	return
}
