package v1

import (
	"gin-blog-server/models/response"
	"gin-blog-server/service"

	"github.com/gin-gonic/gin"
)

func SearchUsername(c *gin.Context) {
	keywords := c.Query("name")
	if keywords == "" {
		response.FailWithMsg("search name cannot be empty", c)
		return
	}
	names, err := service.SearchUsername(keywords)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	response.OKWithData(gin.H{
		"list": names,
	}, c)
}
