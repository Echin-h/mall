package model

// 很抽象，这个CreateAt 和 UpdateAt 无法建立，唉

type Cart struct {
	ID        uint `gorm:"primarykey"`
	UserID    uint
	ProductID uint `gorm:"not null"`
	BossID    uint
	Num       uint
	MaxNum    uint
	Check     bool
}
