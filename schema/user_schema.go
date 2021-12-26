package schema

import (
	"advanced_programming/models"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Name   string `json:"name"`
	UserID int64  `json:"user_id"`
}

type LoginRequest struct {
	Name     string `json:"name"`
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
