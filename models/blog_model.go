package models

import (
	"github.com/jinzhu/gorm"
)

type Blog struct {
	gorm.Model
	BlogID   int    `gorm:"column:id"`
	AuthorID int    `gorm:"column:author_id"`
	Title    string `gorm:"column:title"`
	Content  string `gorm:"column:content"`
	Likes    int    `gorm:"column:likes"`
}

type Comment struct {
	gorm.Model
	AuthorID int `gorm:"column:author_id"`
	BlogID   int `gorm:"column:blog_id"`
}

type BlogLiker struct {
	gorm.Model
	AuthorID int `gorm:"column:user_id"`
	BlogID   int `gorm:"column:blog_id"`
}

// TableName 重命名表，否则 gorm 默认命名为 snake 的复数形式，例如 blogs
func (Blog) TableName() string {
	return "blog"
}

func (Comment) TableName() string {
	return "comment"
}

func (BlogLiker) TableName() string {
	return "blog_liker"
}

// Blogs 实现了堆接口的 Blog slice
type Blogs []*Blog

func (blogs *Blogs) Len() int {
	return len(*blogs)
}

// Less 实现小根堆，likes 小的排在堆顶
func (blogs *Blogs) Less(i, j int) bool {
	return (*blogs)[i].Likes < (*blogs)[j].Likes
}

func (blogs *Blogs) Swap(i, j int) {
	(*blogs)[i], (*blogs)[j] = (*blogs)[j], (*blogs)[i]
}

func (blogs *Blogs) Push(b interface{}) {
	*blogs = append(*blogs, b.(*Blog))
}

func (blogs *Blogs) Pop() interface{} {
	n := len(*blogs)
	ret := (*blogs)[n-1]
	*blogs = (*blogs)[:n-1]
	return ret
}
