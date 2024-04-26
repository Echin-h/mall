package dao

import (
	"context"
	"errors"
	"gin-mall/pkg/e"
	"gin-mall/respository/db/model"
	"gin-mall/types"
	"gorm.io/gorm"
)

type CartDao struct {
	*gorm.DB
}

func NewCartDao(ctx context.Context) *CartDao {
	return &CartDao{NewDBClient(ctx)}
}

func NewCartDaoByDB(db *gorm.DB) *CartDao {
	return &CartDao{db}
}

func (dao *CartDao) CreateCart(uid uint, req *types.CartCreateReq) (cart *model.Cart, status int, err error) {
	cart, err = dao.GetCartById(req.ProductId, uid, req.BossID)
	if errors.Is(err, gorm.ErrRecordNotFound) || cart == nil {
		cart = &model.Cart{
			UserID:    uid,
			ProductID: req.ProductId,
			BossID:    req.BossID,
			Num:       1,
			MaxNum:    20,
			Check:     false,
		}
		err = dao.DB.Create(&cart).Error
		if err != nil {
			return
		}
		return cart, 0, err
	}
	if cart.Num < cart.MaxNum {
		cart.Num++
		err = dao.DB.Save(&cart).Error
		if err != nil {
			return
		}
		return cart, 1, err
	}

	return cart, e.ErrorProductMoreCart, err

}

func (dao *CartDao) GetCartById(pId, uId, bId uint) (cart *model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).
		Where("product_id = ? AND user_id = ? AND boss_id = ?", pId, uId, bId).
		First(&cart).Error
	return
}

func (dao *CartDao) ListCart(uid uint, req *types.CartListReq) (carts []*types.CartResp, err error) {
	err = dao.DB.Model(&model.Cart{}).
		Joins("AS c LEFT JOIN product AS p ON c.product_id = p.id").
		Where("c.user_id = ?", uid).
		Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize).
		Select("c.id AS id," +
			"c.user_id AS user_id," +
			"c.product_id AS product_id," +
			"c.num AS num," +
			"c.max_num AS max_num," +
			"c.check AS check_," +
			"p.img_path AS img_path," +
			"p.boss_id AS boss_id," +
			"p.boss_name AS boss_name," +
			"p.info AS info," +
			"p.discount_price AS discount_price").
		Find(&carts).Error

	return
}

func (dao *CartDao) UpdateCartById(uid uint, req *types.UpdateCartServiceReq) (err error) {
	err = dao.DB.Model(&model.Cart{}).
		Where("id = ? AND user_id = ?", req.Id, uid).
		Update("num", req.Num).Error

	return
}

func (dao *CartDao) DeleteCartById(uid uint, cid uint) (err error) {
	err = dao.DB.Model(&model.Cart{}).
		Where("id = ? AND user_id = ?", cid, uid).
		Delete(&model.Cart{}).Error
	return
}
