package routers

import (
	v1 "gin-blog-server/api/v1"

	"github.com/gin-gonic/gin"
)

func InitSysInfoRouter(routerGrp *gin.RouterGroup) {
	sysInfoRouter := routerGrp.Group("/sysinfo")
	{
		sysInfoRouter.GET("/get", v1.SysInfo)
	}
}
