package schema

import "advanced_programming/models"

type AddStarRequest struct {
	UserID int `json:"user_id"`
	BlogID int `json:"blog_id"`
}

type AddStarResponse struct {
	UserID int `json:"user_id"`
	BlogID int `json:"blog_id"`
}

type TopKStarResponse struct {
	Blogs models.Blogs `json:"blogs"`
}
