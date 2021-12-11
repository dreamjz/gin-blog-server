package routers

import (
	v1 "gin-blog-server/api/v1"

	"github.com/gin-gonic/gin"
)

func InitPublicRouter(routerGrp *gin.RouterGroup) {
	publicRouter := routerGrp.Group("/public")
	{
		publicRouter.GET("ping", v1.Ping)
	}
}
