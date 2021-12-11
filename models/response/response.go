package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS = 2000
	ERROR   = 2001
)

var (
	CodeMsgMap = map[int]string{
		SUCCESS: "Success",
		ERROR:   "Error",
	}
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func GetCodeMsg(code int) string {
	if msg, ok := CodeMsgMap[code]; ok {
		return msg
	}
	return ""
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func OK(c *gin.Context) {
	Result(SUCCESS, "", GetCodeMsg(SUCCESS), c)
}

func OKWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, GetCodeMsg(SUCCESS), c)
}

func Fail(c *gin.Context) {
	Result(ERROR, "", GetCodeMsg(ERROR), c)
}

func FailWithMsg(msg string, c *gin.Context) {
	Result(ERROR, "", msg, c)
}

func FailWithCode(code int, c *gin.Context) {
	Result(code, "", GetCodeMsg(code), c)
}
