package blog

import (
	"advanced_programming/common"
	"advanced_programming/constant"
	"advanced_programming/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowBlogHandler(c *gin.Context) {
	filter := make(map[string]interface{})
	if blogID, ok := c.GetQuery(constant.BlogID); ok {
		filter[constant.ParamMap[constant.BlogID]], _ = strconv.Atoi(blogID) // 为查询条件注入 id
	}
	resp, err := services.ShowBlog(filter)
	if err != nil {
		c.AbortWithStatusJSON(constant.OK, common.BuildHttpErrorResponse(err))
		return
	}
	c.JSON(constant.OK, common.BuildHttpResponse(constant.OK, "OK", resp))
}

func ShowBlogListHandler(c *gin.Context) {
	filter := make(map[string]interface{})
	if limit, ok := c.GetQuery(constant.Limit); ok {
		filter[constant.ParamMap[constant.Limit]] = limit
	}
	resp, err := services.ShowBlogList(filter)
	if err != nil {
		c.AbortWithStatusJSON(constant.OK, common.BuildHttpErrorResponse(err))
		return
	}
	c.JSON(constant.OK, common.BuildHttpResponse(constant.OK, "OK", resp))
}
