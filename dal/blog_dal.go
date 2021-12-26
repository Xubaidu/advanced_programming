package dal

import (
	"advanced_programming/clients"
	"advanced_programming/common"
	"advanced_programming/constant"
	"advanced_programming/models"
	"log"
)

func CreateBlog(blog *models.Blog) (err error) {
	var DB = clients.DB

	// 创建记录
	if err := DB.Create(blog).Error; err != nil {
		log.Printf("%+v", err)
		return common.NewError(constant.RecordInsertError, "创建记录失败")
	}

	return nil
}

func GetBlogs(filter map[string]interface{}) (blogs []*models.Blog, err error) {
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

func GetBlog(filter map[string]interface{}) (blog *models.Blog, err error) {
	var DB = clients.DB

	blog = &models.Blog{}

	// 查询记录，查不到会返回 ErrRecordNotFound 错误
	if err := DB.Where(filter).Take(blog).Error; err != nil {
		log.Printf("%+v", err)
		return nil, common.NewError(constant.RecordNotFound, "没有查询到相关记录")
	}

	return blog, nil
}

func UpdateBlog(filter, updater map[string]interface{}) (err error) {
	var DB = clients.DB

	// 开启事务
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 判断事务是否开启
	if err := tx.Error; err != nil {
		log.Printf("%+v", err)
		return common.NewError(constant.TransactionBeginError, "开启事务失败")
	}

	// 存在才更新
	if _, err = GetBlog(filter); err != nil {
		log.Printf("%+v", err)
		tx.Rollback()
		return common.NewError(constant.RecordNotFound, "没有查询到相关记录")
	}

	// 更新
	for k, v := range updater {
		if err := DB.Model(&models.Blog{}).Where(filter).Update(k, v).Error; err != nil {
			log.Printf("%+v", err)
			tx.Rollback()
			return common.NewError(constant.RecordUpdateError, "更新记录失败")
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		log.Printf("%+v", err)
		tx.Rollback()
		return common.NewError(constant.TransactionCommitError, "提交事务失败")
	}

	// 正常返回
	return nil
}

func DeleteBlog(filter map[string]interface{}) (err error) {
	var DB = clients.DB

	// 删除
	if err := DB.Where(filter).Delete(&models.Blog{}).Error; err != nil {
		log.Printf("%+v", err)
		return common.NewError(constant.RecordDeleteError, "删除记录失败")
	}

	return nil
}
