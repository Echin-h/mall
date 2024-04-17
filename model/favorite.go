package model

import "gorm.io/gorm"

// 商城系统最好不要使用外键，外键会影响性能 i like this comment

type Favorite struct {
	gorm.Model
	User      User    `gorm:"foreignKey:UserId"`
	UserId    uint    `gorm:"not null"`
	Product   Product `gorm:"foreignKey:ProductId"`
	ProductId uint    `gorm:"not null"`
	Boss      User    `gorm:"foreignKey:BossId"`
	BossId    uint    `gorm:"not null"`
}
