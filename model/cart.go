package model

// Cart 购物车

type Cart struct {
	ID        uint `gorm:"primarykey"`
	UserId    uint `gorm:"not null"`
	ProductId uint `gorm:"not null"`
	BossId    uint `gorm:"not null"`
	Num       uint `gorm:"not null"`
	MaxNum    uint `gorm:"not null"`
	Check     bool // 是否支付
}
