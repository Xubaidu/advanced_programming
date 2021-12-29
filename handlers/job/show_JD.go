package job

import (
	"advanced_programming/common"
	"advanced_programming/constant"
	. "advanced_programming/schema"
	"advanced_programming/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowJobHandler(c *gin.Context) {

	req := &GetJobRequest{}
	if jobID, ok := c.GetQuery("job_id"); ok {
		req.JobID, _ = strconv.Atoi(jobID)
	}
	resp, err := services.ShowJobService(req)
	if err != nil {
		c.AbortWithStatusJSON(constant.OK, common.BuildRespByErr(err))
		return
	}
	c.JSON(constant.OK, common.BuildResp(constant.OK, "OK", resp))
}

func ShowJobListHandler(c *gin.Context) {
	req := &GetJobListRequest{}
	if limit, ok := c.GetQuery("limit"); ok {
		req.Limit, _ = strconv.Atoi(limit)
	}
	resp, err := services.ShowJobListService(req)
	if err != nil {
		c.AbortWithStatusJSON(constant.OK, common.BuildRespByErr(err))
		return
	}
	c.JSON(constant.OK, common.BuildResp(constant.OK, "OK", resp))
}
