package blog

import (
	"advanced_programming/common"
	"advanced_programming/constant"
	. "advanced_programming/schema"
	"advanced_programming/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateBlogHandler(c *gin.Context) {
	req := &UpdateBlogRequest{}
	err := common.BindParams(c, req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, common.BuildRespByErr(err))
		return
	}
	resp, err := services.UpdateBlogService(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, common.BuildRespByErr(err))
		return
	}
	c.JSON(constant.OK, common.BuildResp(constant.OK, "OK", resp))
	return
}
