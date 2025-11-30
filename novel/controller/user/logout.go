package user

import (
	"context"
	"net/http"
	"strconv"

	"novel-launch/novel/Biu"
	constance "novel-launch/novel/Biu/const"
	"novel-launch/novel/middleware/db"
	redismw "novel-launch/novel/middleware/redis"
	"novel-launch/novel/server/dao"

	"github.com/gin-gonic/gin"
)

func Logout(ctx *gin.Context) {
	uidVal, ok := ctx.Get("uid")
	if !ok {
		Biu.Failed(ctx, http.StatusUnauthorized, "unauthorized")
		return
	}
	uid := uidVal.(uint64)
	if r := redismw.Get(); r != nil {
		r.Del(context.Background(), "user:token:"+strconv.FormatUint(uid, 10))
	}
	d := dao.NewUserDAO(db.Get())
	if u, err := d.GetUserByUID(ctx, strconv.FormatUint(uid, 10)); err == nil {
		u.Token = ""
		_ = d.UpdateUser(ctx, u)
	}
	Biu.ResponseMsg(ctx, http.StatusOK, constance.Success, constance.SuccessMsg)
}
