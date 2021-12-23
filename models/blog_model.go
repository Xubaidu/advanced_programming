package models

import "github.com/jinzhu/gorm"

type Blog struct {
	gorm.Model
	BlogID   int64  `gorm:"id"`
	AuthorID int64  `gorm:"column:author_id"`
	Title    string `gorm:"column:title"`
	Content  string `gorm:"column:content"`
}

type Comment struct {
	gorm.Model
	AuthorID int64 `gorm:"column:author_id"`
	BlogID   int64 `gorm:"column:blog_id"`
}

// TableName 重命名表，否则 gorm 默认命名为 snake 的复数形式，例如 blogs
func (Blog) TableName() string {
	return "blog"
}

func (Comment) TableName() string {
	return "comment"
}
