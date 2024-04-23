package model

type Follow struct {
	ID           uint `gorm:"primaryKey"`
	User         User `gorm:"foreignKey:FollowId"` // 关注人
	FollowId     uint
	FollowedUser User `gorm:"foreignKey:UserId"` // 被关注人
	UserId       uint
	CreateTime   uint `gorm:"autoCreateTime"`
}

/*
 ForeignKey Specifies the column name of the current model used as a foreign key in the join table.
 references Specifies the column name of the association table used as a foreign key in the join table.
 join table 联结表
 一个是 current model 一个是 association table
*/
