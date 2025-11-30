package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"novel-launch/novel/Biu"
	"novel-launch/novel/server/handler/userhandle"
)

func Register(ctx *gin.Context) {
	var user userhandle.RegisterRequest
	if err := ctx.ShouldBind(&user); err != nil {
		Biu.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}

	httpCode, code, msg := userhandle.Register(ctx, user)
	Biu.ResponseMsg(ctx, httpCode, code, msg)
}
