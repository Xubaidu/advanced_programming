package user

import (
	"advanced_programming/common"
	"advanced_programming/constant"
	. "advanced_programming/schema"
	"advanced_programming/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UploadResumeHandler(c *gin.Context) {

	// 获取 file, header 等参数
	resume, header, err := c.Request.FormFile("resume")
	userID, _ := strconv.Atoi(c.PostForm("user_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, common.BuildRespByCodeMsg(constant.ParamsParseError, "文件读取失败"))
		return
	}

	// 校验 file 大小
	if header.Size > constant.FileSizeLimit {
		c.AbortWithStatusJSON(http.StatusOK, common.BuildRespByCodeMsg(constant.ParamsValidateError, "文件太大"))
		return
	}

	// 构造上传文件的请求
	req := &UploadFileRequest{
		File:     resume,
		FileName: header.Filename,
		UserID:   int64(userID),
	}

	// 获取上传文件的响应
	resp, err := services.UploadFileService(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, common.BuildRespByErr(err))
		return
	}
	c.JSON(constant.OK, common.BuildResp(constant.OK, "OK", resp))
}
