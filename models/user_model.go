package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserID   int64  `gorm:"column:id" json:"user_id"`
	Name     string `gorm:"column:name" json:"name"`
	Password string `gorm:"password" json:"password"`
	Resume   string `gorm:"column:resume" json:"resume"`
}

func (User) TableName() string {
	return "user"
}
