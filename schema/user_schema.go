package schema

import (
	"advanced_programming/models"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type RegisterResponse struct {
	Name   string `json:"name"`
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"Password"`
}

type LoginResponse struct {
	UserInfo *models.User `json:"user_info"`
}

type GetUserInfoRequest struct {
	UserID int `json:"user_id"`
}

type GetUserInfoResponse struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Resume string `json:"resume"`
}

type GetUserApplyRequest struct {
	UserID int
}

type GetUserApplyResponse struct {
	Jobs []*models.Job `json:"jobs"`
}
