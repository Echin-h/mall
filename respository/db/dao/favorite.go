package dao

import (
	"context"
	"gin-mall/respository/db/model"
	"gin-mall/types"
	"gorm.io/gorm"
)

type FavoritesDao struct {
	*gorm.DB
}

func NewFavoritesDao(ctx context.Context) *FavoritesDao {
	return &FavoritesDao{NewDBClient(ctx)}
}

func NewFavoritesDaoByDB(db *gorm.DB) *FavoritesDao {
	return &FavoritesDao{db}
}

// 多表查询的魅力就此显现

func (dao *FavoritesDao) ListFavoritesByUserId(uid uint, pageNum int, pageSize int) (r []*types.FavoriteListResp, total int64, err error) {
	err = dao.DB.Model(&model.Favorite{}).
		Where("user_id = ?", uid).Count(&total).Error
	if err != nil {
		return
	}

	err = dao.DB.Model(&model.Favorite{}).
		Joins("AS f LEFT JOIN user AS u ON u.id = f.boss_id").
		Joins("LEFT JOIN product AS p ON p.id = f.product_id").
		Joins("LEFT JOIN category AS c ON c.id = p.category").
		Where("f.user_id = ?", uid).
		Offset((pageNum - 1) * pageSize).Limit(pageSize).
		Select("f.user_id AS user_id," +
			"f.product_id AS product_id," +
			"UNIX_TIMESTAMP(f.created_at) AS created_at," +
			"p.title AS title," +
			"p.info AS info," +
			"p.name AS name," +
			"c.id AS category_id," +
			"c.category_name AS category_name," +
			"u.id AS boss_id," +
			"u.user_name AS boss_name," +
			"u.avatar AS boss_avatar," +
			"p.price AS price," +
			"p.img_path AS img_path," +
			"p.discount_price AS discount_price," +
			"p.num AS num," +
			"p.on_sale AS on_sale").
		Find(&r).Error

	return
}

func (dao *FavoritesDao) FavoriteExistOrNot(uid uint, productId uint) (bool, error) {
	var count int64
	err := dao.DB.Model(&model.Favorite{}).
		Where("user_id = ? AND product_id = ?", uid, productId).
		Count(&count).Error
	if count == 0 || err != nil {
		return false, err
	}
	return true, nil
}

func (dao *FavoritesDao) CreateFavorite(favorite *model.Favorite) error {
	return dao.DB.Create(&favorite).Error
}

func (dao *FavoritesDao) DeleteFavoriteById(favoriteId uint) error {
	return dao.DB.Where("id = ?", favoriteId).Delete(&model.Favorite{}).Error
}
