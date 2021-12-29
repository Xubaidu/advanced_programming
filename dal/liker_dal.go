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

// GerStarGivers 将 blog_liker 和 user join
// 获取为 blogID 的帖子点赞的所有 user 的 email
// 然后 notify
func GerStarGivers(blogID int) (users []*models.User, err error) {
	var DB = clients.DB

	filter := map[string]interface{}{
		"id": blogID,
	}

	// 先查询 blog 的喜欢数
	blog, err := GetBlog(filter)
	if err != nil {
		return nil, err
	}

	// 如果为 0 直接返回，不用 notify
	if blog.Likes == 0 {
		return users, nil
	}

	// select user.email from user join blog_liker on blog_liker.user_id = user.id
	// where blog_id = blogID;
	DB = DB.Table("user").Where(filter).Select("user.email").
		Joins("join blog_liker on blog_liker.user_id = user.id").Take(&users)

	if DB.Error != nil {
		return nil, common.NewError(constant.RecordNotFound, "没有查询到相关记录")
	}

	return users, nil
}
