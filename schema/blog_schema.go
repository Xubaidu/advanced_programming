package schema

import "advanced_programming/models"

type CreateBlogRequest struct {
	UserID  int    `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreateBlogResponse struct {
	Blog *models.Blog `json:"blog"`
	Msg  string       `json:"msg"`
}

type GetBlogRequest struct {
	BlogID int `json:"blog_id"`
}

type GetBlogResponse struct {
	UserName string       `json:"user_name"`
	Blog     *models.Blog `json:"blog"`
	Msg      string       `json:"msg"`
}

type GetBlogListRequest struct {
	Limit int
}

type GetBlogListResponse struct {
	UserNames []string       `json:"user_names"` // blogs 对应的作者名
	Blogs     []*models.Blog `json:"blogs"`
	Msg       string         `json:"msg"`
}

type UpdateBlogRequest struct {
	UserID  int    `json:"user_id"`
	BlogID  int    `json:"blog_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateBlogResponse struct {
	Msg string `json:"msg"`
}

type DeleteBlogRequest struct {
	BlogID int `json:"blog_id"`
}

type DeleteBlogResponse struct {
	Msg string `json:"msg"`
}
