package services

import (
	"advanced_programming/dal"
	"advanced_programming/models"
	. "advanced_programming/schema"
)

func CreateBlogService(req *CreateBlogRequest) (resp *CreateBlogResponse, err error) {
	blog := &models.Blog{
		AuthorID: req.UserID,
		Title:    req.Title,
		Content:  req.Content,
	}
	dal.CreateBlog(blog)
	resp = &CreateBlogResponse{
		UserID:  req.UserID,
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
	filter := map[string]interface{}{
		"id": req.BlogID,
	}
	updater := map[string]interface{}{
		"title":   req.Title,
		"content": req.Content,
	}
	err = dal.UpdateBlog(filter, updater)
	if err != nil {
		return nil, err
	}
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
