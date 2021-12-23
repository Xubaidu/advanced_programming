package dal

import (
	"advanced_programming/clients"
	"advanced_programming/common"
	"advanced_programming/constant"
	"advanced_programming/models"
	"log"
)

func CreateBlog(app *models.Blog) {
	var DB = clients.DB
	DB.Create(app)
}

func GetBlogs(filter map[string]interface{}) (blogs []*models.Blog, err error) {
	var DB = clients.DB
	DB = DB.Where(filter).Find(&blogs)
	if blogs == nil || len(blogs) == 0 {
		log.Printf("没有符合条件的查询")
		err = &common.CustomError{
			Code:   constant.DataQueryError,
			ErrMsg: "没有符合条件的查询",
		}
		return nil, err
	}
	return blogs, nil
}

func GetBlog(filter map[string]interface{}) (blog *models.Blog, err error) {
	blogs, err := GetBlogs(filter)
	if err != nil {
		return nil, err
	}
	return blogs[0], nil
}

func UpdateBlog(filter, updater map[string]interface{}) (err error) {
	var DB = clients.DB
	DB = DB.Model(&models.Blog{}).Where(filter)
	for k, v := range updater {
		DB = DB.Update(k, v)
	}
	return nil
}

func DeleteBlog(filter map[string]interface{}) (err error) {
	var DB = clients.DB
	DB = DB.Where(filter).Delete(&models.Blog{})
	return nil
}
