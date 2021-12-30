package models

import (
	"github.com/jinzhu/gorm"
)

type Job struct {
	gorm.Model
	JobID      int    `gorm:"column:id" json:"job_id"`
	UserID     int    `gorm:"column:user_id" json:"user_id"` // 发布者
	Company    string `gorm:"column:company" json:"company"` // 公司名称
	Base       string `gorm:"column:base" json:"base"`       // 岗位 base
	Title      string `gorm:"column:title" json:"title"`     // 岗位名称
	Content    string `gorm:"column:content" json:"content"` // JD 内容
	Applicants int    `gorm:"column:applicants"`             // 申请总人数
}

type Apply struct {
	gorm.Model
	ID     int `gorm:"id"`
	JobID  int `gorm:"column:job_id" json:"job_id"`   // 申请岗位
	UserID int `gorm:"column:user_id" json:"user_id"` // 申请者
}

func (Job) TableName() string {
	return "job"
}

func (Apply) TableName() string {
	return "apply"
}
