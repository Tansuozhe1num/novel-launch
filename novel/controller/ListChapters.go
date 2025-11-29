package controller

import (
	"github.com/gin-gonic/gin"
	"novel-launch/novel/Biu"
	constance "novel-launch/novel/Biu/const"
	"novel-launch/novel/server/handler"
)

func ListChapters(ctx *gin.Context) {
	var req handler.ListChapterReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		Biu.Failed(ctx, constance.ParamErr, constance.ParamErrMsg)
		return
	}

	if err, list := handler.ListChapters(ctx, req); err != nil {
		Biu.Failed(ctx, constance.ParamServerErr, constance.ParamServerErrMsg)
		return
	} else {
		Biu.Success(ctx, list)
	}
}
