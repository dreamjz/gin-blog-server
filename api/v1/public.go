package v1

import (
	"gin-blog-server/models/response"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	response.OKWithData("pong", c)
}
