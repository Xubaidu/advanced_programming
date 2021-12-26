package services

import (
	"advanced_programming/clients"
	"advanced_programming/constant"
	"advanced_programming/dal"
	. "advanced_programming/schema"
	"io"
	"log"
	"os"
	"path"
	"strconv"
)

func UploadFileService(req *UploadFileRequest) (resp *UploadFileResponse, err error) {

	// 构造存储路径
	dstPath := constant.SavedFilesPath + "user" + strconv.Itoa(int(req.UserID)) + "_" + req.FileName

	// 打开源文件
	src := req.File
	defer func() {
		err := src.Close()
		if err != nil {
			log.Printf("%v", err)
		}
	}()

	// 打开目标文件
	out, err := os.Create(dstPath) // ignore_security_alert
	if err != nil {
		return nil, err
	}
	defer func() {
		err := out.Close()
		if err != nil {
			log.Printf("%v", err)
		}
	}()

	var DB = clients.DB

	// 开启事务
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 落库
	filter := map[string]interface{}{
		"id": req.UserID,
	}
	updater := map[string]interface{}{
		"resume": dstPath,
	}
	if err := dal.UpdateUser(filter, updater); err != nil {
		log.Printf("%+v", err)
		tx.Rollback()
		return nil, err
	}

	// 文件另存到本地
	if _, err := io.Copy(out, src); err != nil {
		log.Printf("%+v", err)
		tx.Rollback()
		return nil, err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		log.Printf("%+v", err)
		tx.Rollback()
		return nil, err
	}

	// 构造响应
	resp = &UploadFileResponse{
		DstPath: dstPath,
	}
	return resp, nil
}

func DownloadFileService(req *DownloadFileRequest) (resp *DownloadFileResponse, err error) {

	// 查询用户的 resume 地址
	dstPath, err := dal.GetUserResume(req.Filter)
	if err != nil {
		return nil, err
	}

	// 打开文件
	_, err = os.Open(dstPath)
	if err != nil {
		return nil, err
	}

	// 返回响应
	resp = &DownloadFileResponse{
		FileName: path.Base(dstPath),
	}
	return resp, nil
}
