package model

import (
	"gin-mall/respository/cache"
	"gorm.io/gorm"
	"strconv"
)

// TODO: 这里是CategoryID,但是columnName = category 而不是 category_id

type Product struct {
	gorm.Model
	Name          string `gorm:"size:255;index"`
	CategoryID    uint   `gorm:"not null;column:category_id"`
	Title         string
	Info          string `gorm:"size:1000"`
	ImgPath       string
	Price         string
	DiscountPrice string
	OnSale        bool `gorm:"default:false"`
	Num           int
	BossID        uint
	BossName      string
	BossAvatar    string
}

// 获取点击数，点击数是存放在redis中的
func (product *Product) View() uint64 {
	res, _ := cache.RedisClient.Get(cache.RedisContext, cache.ProductViewKey(product.ID)).Result()
	count, _ := strconv.ParseUint(res, 10, 64)
	return count
}
