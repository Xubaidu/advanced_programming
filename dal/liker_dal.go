package dal

import (
	"advanced_programming/clients"
	"advanced_programming/common"
	"advanced_programming/constant"
	"advanced_programming/models"
	"log"
)

func CreateBlogLiker(liker *models.BlogLiker) (err error) {
	var DB = clients.DB

	// 创建记录
	if err := DB.Create(liker).Error; err != nil {
		log.Printf("%+v", err)
		return common.NewError(constant.RecordInsertError, "创建记录失败")
	}

	return nil
}

func GerStarGivers(blogID int) (users []*models.User, err error) {
	var DB = clients.DB

	filter := map[string]interface{}{
		"blog_id": blogID,
	}
	DB = DB.Table("user").Where(filter).Select("user.email").
		Joins("join blog_liker on blog_liker.user_id = user.id").Take(&users)

	if DB.Error != nil {
		return nil, common.NewError(constant.RecordNotFound, "没有查询到相关记录")
	}

	return users, nil
}
