package job

import (
	"advanced_programming/common"
	"advanced_programming/constant"
	. "advanced_programming/schema"
	"advanced_programming/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CampaignJobHandler 发布一个职位
func CampaignJobHandler(c *gin.Context) {

	// 绑定请求
	req := &CampaignJobRequest{}
	err := common.BindParams(c, req)

	// 获取响应
	resp, err := services.CampaignJobService(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, common.BuildRespByErr(err))
		return
	}

	// 返回响应
	c.JSON(constant.OK, common.BuildResp(constant.OK, "OK", resp))
}
