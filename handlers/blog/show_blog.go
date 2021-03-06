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

	req := &GetBlogRequest{}
	if blogID, ok := c.GetQuery("blog_id"); ok {
		req.BlogID, _ = strconv.Atoi(blogID)
	}
	resp, err := services.ShowBlogService(req)
	if err != nil {
		c.AbortWithStatusJSON(constant.OK, common.BuildRespByErr(err))
		return
	}
	c.JSON(constant.OK, common.BuildResp(constant.OK, "OK", resp))
}

func ShowBlogListHandler(c *gin.Context) {
	req := &GetBlogListRequest{}
	if limit, ok := c.GetQuery("limit"); ok {
		req.Limit, _ = strconv.Atoi(limit)
	}
	resp, err := services.ShowBlogListService(req)
	if err != nil {
		c.AbortWithStatusJSON(constant.OK, common.BuildRespByErr(err))
		return
	}
	c.JSON(constant.OK, common.BuildResp(constant.OK, "OK", resp))
}
