package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserID   int64  `gorm:"column:id"`
	Name     string `gorm:"column:name"`
	Password string `gorm:"password"`
	Resume   string `gorm:"column:resume"`
}

func (User) TableName() string {
	return "user"
}
