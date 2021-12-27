package services

import (
	"advanced_programming/clients"
	"advanced_programming/common"
	"advanced_programming/constant"
	"advanced_programming/dal"
	"advanced_programming/models"
	. "advanced_programming/schema"
	"log"
)

func CreateBlogService(req *CreateBlogRequest) (resp *CreateBlogResponse, err error) {
	blog := &models.Blog{
		AuthorID: req.UserID,
		Title:    req.Title,
		Content:  req.Content,
	}
	if err := dal.CreateBlog(blog); err != nil {
		return nil, err
	}
	resp = &CreateBlogResponse{
		UserID:  req.UserID,
		BlogID:  blog.BlogID,
		Title:   req.Title,
		Content: req.Content,
	}
	return resp, nil
}

func ShowBlogService(req *GetBlogRequest) (resp *GetBlogResponse, err error) {
	blog, err := dal.GetBlog(req.Filter)
	if err != nil {
		return nil, err
	}
	resp = &GetBlogResponse{
		Blog: blog,
	}
	return resp, nil
}

func ShowBlogListService(req *GetBlogListRequest) (resp *GetBlogListResponse, err error) {
	blogs, err := dal.GetBlogs(req.Filter)
	if err != nil {
		return nil, err
	}
	resp = &GetBlogListResponse{
		Blogs: blogs,
	}
	return resp, nil
}

func UpdateBlogService(req *UpdateBlogRequest) (resp *UpdateBlogResponse, err error) {

	// 开启事务
	var DB = clients.DB
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 判断事务是否开启
	if err := tx.Error; err != nil {
		log.Printf("%+v", err)
		tx.Rollback()
		return nil, common.NewError(constant.TransactionBeginError, "开启事务失败")
	}

	// 更新 blog
	filter := map[string]interface{}{
		"id": req.BlogID,
	}
	updater := map[string]interface{}{
		"title":   req.Title,
		"content": req.Content,
	}
	if err := dal.UpdateBlog(filter, updater); err != nil {
		tx.Rollback()
		return nil, err
	}

	// 获取点赞者
	users, err := dal.GerStarGivers(req.BlogID)
	if err != nil {
		return nil, err
	}

	// 提醒点赞者
	for _, user := range users {
		req := SendMailRequest{
			From:        constant.SMTPUsername,
			To:          user.Email,
			Subject:     "更新提示",
			ContentType: constant.MailContentType,
			Body:        "你关注的帖子更新了！",
		}
		if _, err := SendMailService(&req); err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// 构造响应并返回
	resp = &UpdateBlogResponse{
		Title:   req.Title,
		Content: req.Content,
		BlogID:  req.BlogID,
	}
	return resp, nil
}

func DeleteBlogService(req *DeleteBlogRequest) (resp *DeleteBlogResponse, err error) {
	filter := map[string]interface{}{
		"id": req.BlogID,
	}
	err = dal.DeleteBlog(filter)
	if err != nil {
		return nil, err
	}
	resp = &DeleteBlogResponse{
		BlogID: req.BlogID,
	}
	return resp, nil
}
