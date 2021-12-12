package routers

import (
	v1 "gin-blog-server/api/v1"

	"github.com/gin-gonic/gin"
)

func InitArticleRouter(routerGrp *gin.RouterGroup) {
	articleRouter := routerGrp.Group("/article")
	{
		articleRouter.POST("/create", v1.CreateArticle)
		articleRouter.GET("/list", v1.QueryArticleList)
		articleRouter.GET("/detail", v1.QueryArticleByID)
		articleRouter.PUT("/edit", v1.EditArticleByID)
		articleRouter.DELETE("/delete", v1.DeleteArticleByID)
	}
}
