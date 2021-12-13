package v1

import (
	"gin-blog-server/models/response"
	"gin-blog-server/service"

	"github.com/gin-gonic/gin"
)

func SysInfo(c *gin.Context) {
	server, err := service.SystemInfo()
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	response.OKWithData(server, c)
}
