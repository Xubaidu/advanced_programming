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

func DownloadResumeHandler(c *gin.Context) {

	// 	获取 user_id，用于后续查询
	userID, _ := strconv.Atoi(c.Query("user_id"))

	// 构造请求
	req := &DownloadFileRequest{
		Filter: map[string]interface{}{
			"id": userID,
		},
	}

	// 获取响应
	resp, err := services.DownloadFileService(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, common.BuildRespByErr(err))
		return
	}

	// 重点：设置一些 header 信息
	header := map[string]string{
		"Content-Type":              "application/octet-stream",
		"Content-Disposition":       "inline;filename = " + resp.FileName, // 浏览器下载或预览
		"Content-Transfer-Encoding": "binary",
		"Cache-Control":             "no-cache",
	}
	for k, v := range header {
		c.Header(k, v)
	}
	c.JSON(constant.OK, common.BuildResp(constant.OK, "OK", resp))
}
