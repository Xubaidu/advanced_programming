package common

import (
	"advanced_programming/constant"
	"log"

	"github.com/gin-gonic/gin"
)

func BindParams(c *gin.Context, params interface{}) error {
	err := c.ShouldBind(params)
	if err != nil {
		log.Printf("参数绑定失败 %v\n", err)
		return NewError(constant.ParamsBindError, "参数绑定失败")
	}
	return nil
}
