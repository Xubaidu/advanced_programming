package blog

import (
	"advanced_programming/common"
	"advanced_programming/constant"
	"advanced_programming/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func TopKBlogHandler(c *gin.Context) {
	k, _ := strconv.Atoi(c.Query("k"))
	resp, err := services.TopKStarService(k)
	if err != nil {
		c.AbortWithStatusJSON(constant.OK, common.BuildRespByErr(err))
		return
	}
	c.JSON(constant.OK, common.BuildResp(constant.OK, "OK", resp))
}
