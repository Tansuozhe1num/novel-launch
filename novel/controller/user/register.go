package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"novel-launch/novel/Biu"
	constance "novel-launch/novel/Biu/const"
	"novel-launch/novel/middleware/auth"
	"novel-launch/novel/server/handler/userhandle"
)

func Register(ctx *gin.Context) {
	res, err := auth.AllowRequest(ctx, "register", constance.TokenRate, constance.TokenCapacity)
	if !res {
		Biu.ResponseMsg(ctx, http.StatusTooManyRequests, constance.ParamErr, constance.ParamErrMsg)
		return
	}
	if err != nil {
		Biu.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	var user userhandle.RegisterRequest
	if err := ctx.ShouldBind(&user); err != nil {
		Biu.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}

	httpCode, code, msg := userhandle.Register(ctx, user)
	Biu.ResponseMsg(ctx, httpCode, code, msg)
}
