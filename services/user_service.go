package services

import (
	"advanced_programming/common"
	"advanced_programming/constant"
	"advanced_programming/dal"
	"advanced_programming/models"
	. "advanced_programming/schema"
	"log"
)

func RegisterService(req *RegisterRequest) (resp *RegisterResponse, err error) {
	filter := map[string]interface{}{
		"name": req.Name,
	}
	if user, _ := dal.GetUser(filter); user != nil {
		log.Printf("用户名已存在")
		return nil, common.NewError(constant.UserDuplicateError, "用户名已存在")
	}
	app := &models.User{
		Name:     req.Name,
		Password: req.Password,
	}
	err = dal.CreateUser(app)
	if err != nil {
		return nil, err
	}
	resp = &RegisterResponse{
		Name:   app.Name,
		UserID: app.UserID,
	}
	return resp, nil
}

func LoginService(req *LoginRequest) (resp *LoginResponse, err error) {
	filter := map[string]interface{}{
		"name": req.Name,
	}
	user, _ := dal.GetUser(filter)
	if user == nil {
		log.Printf("用户不存在")
		return nil, common.NewError(constant.UserMissingError, "错误的用户名，用户不存在，你在搞笑？")
	}
	if user.Password != req.Password {
		log.Printf("密码错误")
		return nil, common.NewError(constant.PasswordWrongError, "密码都能输错？")
	}
	resp = &LoginResponse{
		UserInfo: user,
	}
	return resp, nil
}

func GetUserInfoService(req *GetUserInfoRequest) (resp *GetUserInfoResponse, err error) {
	user, err := dal.GetUser(req.Filter)
	if err != nil {
		return nil, err
	}
	resp = &GetUserInfoResponse{
		UserInfo: user,
	}
	return resp, nil
}
