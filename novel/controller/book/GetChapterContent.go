package book

import (
	"novel-launch/novel/Biu"
	constance "novel-launch/novel/Biu/const"
	"novel-launch/novel/server/handler/bookHandle"

	"github.com/gin-gonic/gin"
)

func GetChapterContent(ctx *gin.Context) {
	var req bookHandle.ChapterContentRequest
	_ = ctx.ShouldBindQuery(&req)
	if req.BookId == 0 || req.ChapterId == 0 {
		Biu.Failed(ctx, constance.ParamErr, constance.ParamErrMsg)
		return
	}
	if err, resp := bookHandle.GetChapterContent(ctx, req); err != nil {
		Biu.Failed(ctx, constance.ParamServerErr, constance.ParamServerErrMsg)
		return
	} else {
		Biu.Success(ctx, resp)
	}
}
