package dal

import (
	"advanced_programming/clients"
	"advanced_programming/common"
	"advanced_programming/constant"
	"advanced_programming/models"
	"log"
)

// CreateBlog 创建一条 blog 记录
func CreateBlog(blog *models.Blog) (err error) {
	var DB = clients.DB

	// 创建记录
	if err := DB.Create(blog).Error; err != nil {
		log.Printf("%+v", err)
		return common.NewError(constant.RecordInsertError, "创建记录失败")
	}

	return nil
}

// GetAllBlogs 查询所有 blog
func GetAllBlogs(filter map[string]interface{}) (blogs models.Blogs, err error) {
	var DB = clients.DB

	// 查询记录
	if err := DB.Where(filter).Find(&blogs).Error; err != nil {
		log.Printf("%+v", err)
		return nil, common.NewError(constant.RecordNotFound, "没有查询到相关记录")
	}

	// 异常处理
	if len(blogs) == 0 {
		log.Printf("没有符合条件的查询")
		return nil, common.NewError(constant.DataQueryError, "没有符合条件的查询")
	}

	return blogs, nil
}

// GetLimitBlogs 查询 limit 个 blog
func GetLimitBlogs(limit int, filter map[string]interface{}) (blogs models.Blogs, err error) {
	var DB = clients.DB

	// 查询记录
	if err := DB.Limit(limit).Where(filter).Find(&blogs).Error; err != nil {
		log.Printf("%+v", err)
		return nil, common.NewError(constant.RecordNotFound, "没有查询到相关记录")
	}

	// 异常处理
	if len(blogs) == 0 {
		log.Printf("没有符合条件的查询")
		return nil, common.NewError(constant.DataQueryError, "没有符合条件的查询")
	}

	return blogs, nil
}

// GetBlog 查询一个 blog
func GetBlog(filter map[string]interface{}) (blog *models.Blog, err error) {

	var DB = clients.DB

	// 查询记录，查不到会返回 ErrRecordNotFound 错误
	blog = &models.Blog{}
	if err := DB.Where(filter).Take(blog).Error; err != nil {
		log.Printf("%+v", err)
		return nil, common.NewError(constant.RecordNotFound, "没有查询到相关记录")
	}

	return blog, nil
}

// UpdateBlog 修改 blog
func UpdateBlog(filter, updater map[string]interface{}) (err error) {
	var DB = clients.DB

	// 判断是否存在
	if _, err = GetBlog(filter); err != nil {
		log.Printf("%+v", err)
		return common.NewError(constant.RecordNotFound, "没有查询到相关记录")
	}

	// 更新
	for k, v := range updater {
		if err := DB.Model(&models.Blog{}).Where(filter).Update(k, v).Error; err != nil {
			log.Printf("%+v", err)
			return common.NewError(constant.RecordUpdateError, "更新记录失败")
		}
	}

	// 正常返回
	return nil
}

// DeleteBlog 删除 blog
func DeleteBlog(filter map[string]interface{}) (err error) {
	var DB = clients.DB

	// 删除
	if err := DB.Where(filter).Delete(&models.Blog{}).Error; err != nil {
		log.Printf("%+v", err)
		return common.NewError(constant.RecordDeleteError, "删除记录失败")
	}

	return nil
}
