package services

import (
	"advanced_programming/clients"
	"advanced_programming/common"
	"advanced_programming/constant"
	"advanced_programming/dal"
	"advanced_programming/models"
	. "advanced_programming/schema"
	"fmt"
	"log"
)

// CreateBlogService 创建新 blog
func CreateBlogService(req *CreateBlogRequest) (resp *CreateBlogResponse, err error) {

	// 构造 blog 对象
	blog := &models.Blog{
		UserID:  req.UserID,
		Title:   req.Title,
		Content: req.Content,
	}

	// 插入记录
	if err := dal.CreateBlog(blog); err != nil {
		return nil, err
	}

	// 返回响应
	msg := fmt.Sprintf("成功发布帖子 %d", blog.BlogID)
	resp = &CreateBlogResponse{
		Blog: blog,
		Msg:  msg,
	}
	return resp, nil
}

// ShowBlogService 展示具体的 blog
func ShowBlogService(req *GetBlogRequest) (resp *GetBlogResponse, err error) {

	// 构造过滤器
	filter := map[string]interface{}{
		"id": req.BlogID,
	}

	// 查询记录
	blog, err := dal.GetBlog(filter)
	if err != nil {
		return nil, err
	}

	// 根据 user_id 获取 user_name
	filter = map[string]interface{}{
		"id": blog.UserID,
	}
	userName, err := dal.GetUserName(filter)
	if err != nil {
		return nil, err
	}

	// 返回响应
	msg := fmt.Sprintf("成功展示帖子 %d", blog.BlogID)
	resp = &GetBlogResponse{
		UserName: userName,
		Blog:     blog,
		Msg:      msg,
	}
	return resp, nil
}

// ShowBlogListService 展示一页 blog
func ShowBlogListService(req *GetBlogListRequest) (resp *GetBlogListResponse, err error) {

	// 获取 blog list
	blogs, err := dal.GetLimitBlogs(req.Limit, nil)
	if err != nil {
		return nil, err
	}

	// 根据 blog 维护的 user_id
	// 获取每一个 blog 的作者名
	var userNames []string
	for _, blog := range blogs {
		filter := map[string]interface{}{
			"user_id": blog.UserID,
		}
		if userName, err := dal.GetUserName(filter); err != nil {
			userNames = append(userNames, userName)
		}
	}

	// 返回响应
	msg := fmt.Sprintf("成功展示 %d 条帖子", req.Limit)
	resp = &GetBlogListResponse{
		UserNames: userNames,
		Blogs:     blogs,
		Msg:       msg,
	}
	return resp, nil
}

// UpdateBlogService 更新博客信息
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

	// 构造响应
	msg := fmt.Sprintf("成功更新帖子 %d", req.BlogID)
	resp = &UpdateBlogResponse{
		Msg: msg,
	}

	// 获取点赞者
	users, err := dal.GerStarGivers(req.BlogID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		log.Printf("%+v", err)
		tx.Rollback()
		return nil, err
	}

	// 如果有点赞者，提醒他们
	if users != nil {
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
	}
	return resp, nil
}

// DeleteBlogService 删除 blog
func DeleteBlogService(req *DeleteBlogRequest) (resp *DeleteBlogResponse, err error) {

	// 删除记录
	filter := map[string]interface{}{
		"id": req.BlogID,
	}
	err = dal.DeleteBlog(filter)
	if err != nil {
		return nil, err
	}

	// 构造响应
	msg := fmt.Sprintf("成功删除帖子 %d", req.BlogID)
	resp = &DeleteBlogResponse{
		Msg: msg,
	}
	return resp, nil
}
