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

	// 先查找用户是否已存在
	filter := map[string]interface{}{
		"email": req.Email,
	}
	if user, _ := dal.GetUser(filter); user != nil {
		log.Printf("用户已存在")
		return nil, err
	}

	// 创建用户
	user := &models.User{
		Name:     req.Name,
		Password: req.Password,
		Email:    req.Email,
	}
	if err := dal.CreateUser(user); err != nil {
		return nil, err
	}

	// 给注册邮箱发一封邮件
	sendMailReq := &SendMailRequest{
		From:        constant.SMTPUsername,
		To:          req.Email,
		Subject:     constant.MailGreetingSubject,
		ContentType: constant.MailContentType,
		Body:        constant.MailGreetingBody,
	}
	if _, err := SendMailService(sendMailReq); err != nil {
		return nil, err
	}

	// 返回响应并返回
	resp = &RegisterResponse{
		Name:   user.Name,
		UserID: user.UserID,
		Email:  user.Email,
	}
	return resp, nil
}

func LoginService(req *LoginRequest) (resp *LoginResponse, err error) {

	// 查找用户
	filter := map[string]interface{}{
		"email": req.Email,
	}
	user, err := dal.GetUser(filter)

	// 数据库校验
	if err != nil {
		log.Printf("用户不存在")
		return nil, common.NewError(constant.UserMissingError, "错误的用户名，用户不存在，你在搞笑？")
	}

	// 密码校验
	if user.Password != req.Password {
		log.Printf("密码错误")
		return nil, common.NewError(constant.PasswordWrongError, "密码都能输错？")
	}

	// 构造响应并返回
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
