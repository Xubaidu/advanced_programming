package user

import (
	"advanced_programming/common"
	"advanced_programming/constant"
	. "advanced_programming/schema"
	"advanced_programming/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserInfoHandler(c *gin.Context) {
	filter := map[string]interface{}{
		"name": c.Query("name"),
	}
	req := &GetUserInfoRequest{
		Filter: filter,
	}
	resp, err := services.GetUserInfoService(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, common.BuildRespByErr(err))
		return
	}
	c.JSON(constant.OK, common.BuildResp(constant.OK, "OK", resp))
	return
}
