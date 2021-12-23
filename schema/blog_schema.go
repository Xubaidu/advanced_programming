package schema

import "advanced_programming/models"

type CreateBlogRequest struct {
	UserID  int64  `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreateBlogResponse struct {
	UserID  int64  `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type GetBlogResponse struct {
	Blog *models.Blog `json:"blog"`
}

type GetBlogListResponse struct {
	Blogs []*models.Blog `json:"blogs"`
}

type UpdateBlogRequest struct {
	BlogID  int64  `json:"blog_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateBlogResponse struct {
	BlogID  int64  `json:"blog_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type DeleteBlogRequest struct {
	BlogID int64 `json:"blog_id"`
}

type DeleteBlogResponse struct {
	BlogID int64 `json:"blog_id"`
}
