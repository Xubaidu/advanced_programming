package models

import "github.com/jinzhu/gorm"

type LarkToken struct {
	gorm.Model
	ID        int    `gorm:"id" json:"id"`
	TokenName string `gorm:"token_name" json:"name"`
	Token     string `gorm:"token" json:"token"`
}

func (LarkToken) TableName() string {
	return "lark_token"
}
