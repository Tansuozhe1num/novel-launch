package controller

import (
	"github.com/gin-gonic/gin"
	"novel-launch/novel/Biu"
	constance "novel-launch/novel/Biu/const"
	"novel-launch/novel/server/handler"
)

func ShowBookById(ctx *gin.Context) {
	var req handler.ShowBookRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		Biu.Failed(ctx, constance.ParamServerErr, constance.ParamServerErrMsg)
		return
	}

	code, msg, book := handler.ShowBookById(ctx, req)
	if code != constance.Success {
		Biu.Failed(ctx, code, msg)
		return
	}
	Biu.Success(ctx, book)
}
