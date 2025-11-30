package Biu

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, data any) {
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}

func Failed(c *gin.Context, code int, msg string) {
	c.JSON(200, gin.H{
		"code": code,
		"msg":  msg,
	})
}

func JSON(ctx *gin.Context, status int, resp gin.H) {
	ctx.JSON(status, resp)
}

func ResponseMsg(ctx *gin.Context, status int, code int, msg string) {
	ctx.JSON(status, gin.H{
		"code": code,
		"msg":  msg,
	})
}
func ResponseData(ctx *gin.Context, status int, code, msg string, data interface{}) {
	ctx.JSON(status, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
