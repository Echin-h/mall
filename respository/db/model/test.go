package model

import "gorm.io/gorm"

// has one
type Usere struct {
	gorm.Model
	CreditCard CreditCard `gorm:"foreignKey:UsereID"`
}

type CreditCard struct {
	gorm.Model
	Number  string
	UsereID uint
}

// belong to
type Userb struct {
	gorm.Model
	Name      string
	CompanyID int
	Company   Company
}

type Company struct {
	ID   int
	Name string
}
