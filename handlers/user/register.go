package user

import (
	"advanced_programming/common"
	"advanced_programming/constant"
	"advanced_programming/dal"
	"advanced_programming/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func RegisterHandler(c *gin.Context) {
	req := RegisterRequest{}
	err := common.BindParams(c, &req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, common.BuildHttpErrorResponse(err))
		return
	}
	app := &models.User{
		Name:     req.Name,
		Password: req.Password,
	}
	dal.CreateUser(app)
	c.JSON(200, common.BuildHttpResponse(constant.OK, "OK", req))
	return
}
