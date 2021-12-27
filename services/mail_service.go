package services

import (
	"advanced_programming/common"
	"advanced_programming/constant"
	. "advanced_programming/schema"
	"log"

	"gopkg.in/gomail.v2"
)

func SendMailService(req *SendMailRequest) (resp *SendMailResponse, err error) {

	// 构造连接
	m := gomail.NewMessage()

	// 设置 header
	header := map[string]string{
		"From":    req.From,    // 发件人
		"To":      req.To,      // 收件人
		"Subject": req.Subject, // 主题
	}
	for k, v := range header {
		m.SetHeader(k, v)
	}

	// 设置内容
	m.SetBody(req.ContentType, req.Body)

	// 进行连接
	d := gomail.NewDialer(
		constant.SMTPHost,
		constant.SMTPPort,
		constant.SMTPUsername,
		constant.SMTPPassword,
	)

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		log.Printf("%+v", err)
		return nil, common.NewError(constant.MailSendError, "发送失败")
	}

	// 构造响应
	resp = &SendMailResponse{
		Msg: "成功发送",
	}
	return resp, nil
}
