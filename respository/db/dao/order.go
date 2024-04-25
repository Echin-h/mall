package dao

import (
	"context"
	"gin-mall/respository/db/model"
	"gin-mall/types"
	"gorm.io/gorm"
)

type OrderDao struct {
	*gorm.DB
}

func NewOrderDao(ctx context.Context) *OrderDao {
	return &OrderDao{NewDBClient(ctx)}
}

func NewOrderDaoByDB(db *gorm.DB) *OrderDao {
	return &OrderDao{db}
}

func (dao *OrderDao) CreateOrder(order *model.Order) (err error) {
	return dao.DB.Create(&order).Error
}

// Join非常的耗性能，后续可以考虑走缓存

func (dao *OrderDao) GetOrderList(uid uint, req *types.OrderListReq) (orders []*types.OrderListResp, count int64, err error) {
	d := dao.DB.Model(&model.Order{}).Where("user_id =?", uid)
	if req.Type != 0 {
		d.Where("type = ?", req.Type)
	}
	err = d.Count(&count).Error

	db := dao.DB.Model(&model.Order{}).
		Joins("AS o LEFT JOIN product AS p ON p.id = o.product_id").
		Joins("LEFT JOIN address AS a ON a.id = o.address_id").
		Where("o.user_id = ?", uid)
	if req.Type != 0 {
		db.Where("o.type = ?", req.Type)
	}
	err = db.Offset((req.PageNum - 1) * req.PageSize).
		Limit(req.PageSize).Order("created_at DESC").
		Select("o.id AS id," +
			"o.order_num AS order_num," +
			"UNIX_TIMESTAMP(o.created_at) AS created_at," +
			"UNIX_TIMESTAMP(o.updated_at) AS updated_at," +
			"o.user_id AS user_id," +
			"o.product_id AS product_id," +
			"o.boss_id AS boss_id," +
			"o.num AS num," +
			"o.type AS type," +
			"p.name AS name," +
			"p.discount_price AS discount_price," +
			"p.img_path AS img_path," +
			"a.name AS address_name," +
			"a.phone AS address_phone," +
			"a.address AS address").
		Find(&orders).Error

	return
}
