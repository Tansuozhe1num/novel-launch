package user

import (
	"net/http"
	"novel-launch/novel/Biu"
	"novel-launch/novel/server/handler/userhandle"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var req userhandle.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		Biu.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}

	resp := userhandle.Login(ctx, req)
	if resp.Code != http.StatusOK {
		Biu.ResponseMsg(ctx, http.StatusOK, resp.Code, resp.Msg)
		return
	}
	Biu.Success(ctx, gin.H{"token": resp.Token})
}
