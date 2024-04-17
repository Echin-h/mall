package model

import "gorm.io/gorm"

// Order 订单
type Order struct {
	gorm.Model
	UserId    uint `gorm:"not null"`
	ProductId uint `gorm:"not null"`
	BossId    uint `gorm:"not null"`
	AddressId uint `gorm:"not null"`
	Num       uint `gorm:"not null"`
	OrderNum  uint64
	Type      uint    // 1.未支付 2.已支付 3.已发货 4.已收货 5.已评价
	Money     float64 // 这里明文存储
}
