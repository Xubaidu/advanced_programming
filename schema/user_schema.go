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
	Filter map[string]interface{} `json:"filter"`
}

type GetUserInfoResponse struct {
	UserInfo *models.User `json:"user_info"`
}
