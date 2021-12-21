package models

import "github.com/jinzhu/gorm"

type Blog struct {
	gorm.Model
	AuthorId int
	Title string
}

