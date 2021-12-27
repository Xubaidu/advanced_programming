package blog

import (
	"advanced_programming/common"
	"advanced_programming/constant"
	. "advanced_programming/schema"
	"advanced_programming/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddStarHandler(c *gin.Context) {
	req := &AddStarRequest{}
	err := common.BindParams(c, req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, common.BuildRespByErr(err))
		return
	}
	resp, err := services.AddStarService(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, common.BuildRespByErr(err))
		return
	}
	c.JSON(constant.OK, common.BuildResp(constant.OK, "OK", resp))
	return
}
