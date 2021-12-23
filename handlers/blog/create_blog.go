package blog

import (
	"advanced_programming/common"
	"advanced_programming/constant"
	. "advanced_programming/schema"
	"advanced_programming/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBlogHandler(c *gin.Context) {
	req := CreateBlogRequest{}
	err := common.BindParams(c, &req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, common.BuildHttpErrorResponse(err))
		return
	}
	resp, err := services.CreateBlog(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, common.BuildHttpErrorResponse(err))
		return
	}
	c.JSON(constant.OK, common.BuildHttpResponse(constant.OK, "OK", resp))
	return
}
