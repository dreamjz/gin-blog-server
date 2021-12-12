package routers

import (
	v1 "gin-blog-server/api/v1"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(routerGrp *gin.RouterGroup) {
	userRouter := routerGrp.Group("user")
	{
		userRouter.GET("/name", v1.SearchUsername)
	}
}
