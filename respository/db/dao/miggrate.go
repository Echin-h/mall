package dao

import (
	"gin-mall/model"
	model2 "gin-mall/respository/db/model"
)

func migrate() (err error) {
	err = _db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
		&model.User{}, &model.Favorite{},
		&model.Order{}, &model.Admin{}, &model.Address{},
		&model.Cart{}, &model.Category{}, &model.Carousel{},
		&model.Notice{}, &model.Notice{}, &model.Product{},
		&model.ProductImg{}, &model2.SkillProduct{}, &model2.SkillProduct2MQ{},
	)

	return
}
