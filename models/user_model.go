package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserID   int    `gorm:"column:id" json:"user_id"`
	Name     string `gorm:"column:name" json:"name"`
	Password string `gorm:"password" json:"password"`
	Email    string `gorm:"email" json:"email"`
	Resume   string `gorm:"column:resume" json:"resume"`
}

func (User) TableName() string {
	return "user"
}
