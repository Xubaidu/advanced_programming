package blog

import (
	"advanced_programming/common"
	"advanced_programming/constant"
	. "advanced_programming/schema"
	"advanced_programming/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowBlogHandler(c *gin.Context) {
	filter := make(map[string]interface{})
	if blogID, ok := c.GetQuery(constant.BlogID); ok {
		filter[constant.ParamMap[constant.BlogID]], _ = strconv.Atoi(blogID) // 为查询条件注入 id
	}
	req := &GetBlogRequest{
		Filter: filter,
	}
	resp, err := services.ShowBlogService(req)
	if err != nil {
		c.AbortWithStatusJSON(constant.OK, common.BuildRespByErr(err))
		return
	}
	c.JSON(constant.OK, common.BuildResp(constant.OK, "OK", resp))
}

func ShowBlogListHandler(c *gin.Context) {
	filter := make(map[string]interface{})
	if limit, ok := c.GetQuery(constant.Limit); ok {
		filter[constant.ParamMap[constant.Limit]] = limit
	}
	req := &GetBlogListRequest{
		Filter: filter,
	}
	resp, err := services.ShowBlogListService(req)
	if err != nil {
		c.AbortWithStatusJSON(constant.OK, common.BuildRespByErr(err))
		return
	}
	c.JSON(constant.OK, common.BuildResp(constant.OK, "OK", resp))
}
