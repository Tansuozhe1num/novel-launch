package controller

import (
	"net/http"
	"novel-launch/novel/Biu"
	constance "novel-launch/novel/Biu/const"
	"novel-launch/novel/server/handler"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListChapters(ctx *gin.Context) {
	bidStr := ctx.DefaultQuery("bookId", "0")
	bid, _ := strconv.ParseInt(bidStr, 10, 64)
	if ctx.Request.Method == http.MethodPost && bid == 0 {
		var req struct {
			BookId int64 `json:"bookId" form:"bookId"`
		}
		_ = ctx.ShouldBind(&req)
		bid = req.BookId
	}
	if bid == 0 {
		Biu.Failed(ctx, constance.ParamErr, constance.ParamErrMsg)
		return
	}
	if err, list := handler.ListChapters(ctx, bid); err != nil {
		Biu.Failed(ctx, constance.ParamServerErr, constance.ParamServerErrMsg)
		return
	} else {
		Biu.Success(ctx, list)
	}
}
