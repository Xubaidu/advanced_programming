package schema

import "advanced_programming/models"

type CreateBlogRequest struct {
	UserID  int    `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreateBlogResponse struct {
	UserID  int    `json:"user_id"`
	BlogID  int    `json:"blog_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type GetBlogRequest struct {
	Filter map[string]interface{} `json:"filter"`
}

type GetBlogResponse struct {
	Blog *models.Blog `json:"blog"`
}

type GetBlogListRequest struct {
	Filter map[string]interface{} `json:"filter"`
}

type GetBlogListResponse struct {
	Blogs []*models.Blog `json:"blogs"`
}

type UpdateBlogRequest struct {
	UserID  int    `json:"user_id"`
	BlogID  int    `json:"blog_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateBlogResponse struct {
	BlogID  int    `json:"blog_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type DeleteBlogRequest struct {
	BlogID int `json:"blog_id"`
}

type DeleteBlogResponse struct {
	BlogID int `json:"blog_id"`
}
