package v1

import (
	"gin-blog-server/models"
	"gin-blog-server/models/request"
	"gin-blog-server/models/response"
	"gin-blog-server/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func CreateArticle(c *gin.Context) {
	var article models.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		log.Println("Bind data error: ", err.Error())
		response.FailWithMsg(err.Error(), c)
		return
	}
	if err := service.CreateArticle(&article); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	response.OK(c)
}

func QueryArticleList(c *gin.Context) {
	var pagination request.Pagination
	if err := c.ShouldBindQuery(&pagination); err != nil {
		log.Print("Bind pagination error: ", err.Error())
		response.FailWithMsg(err.Error(), c)
		return
	}
	list, err := service.QueryArticleList(pagination)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	response.OKWithData(list, c)
}

func QueryArticleContentByID(c *gin.Context) {
	id, err := cast.ToUintE(c.Query("id"))
	if err != nil {
		log.Print("Get article id error: ", err.Error())
		response.FailWithMsg(err.Error(), c)
		return
	}
	content, err := service.QueryArticleContentByID(id)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	response.OKWithData(content, c)
}

func EditArticleByID(c *gin.Context) {
	var article models.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		log.Print("Bind article data error: ", err.Error())
		response.FailWithMsg(err.Error(), c)
		return
	}
	if err := service.UpdateArticleByID(&article); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	response.OK(c)
}

func DeleteArticleByID(c *gin.Context) {
	id, err := cast.ToUintE(c.Query("id"))
	if err != nil {
		log.Print("Get id error: ", err.Error())
		response.FailWithMsg(err.Error(), c)
		return
	}
	if err := service.DeleteArticleByID(id); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	response.OK(c)
}
