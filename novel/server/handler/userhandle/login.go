package userhandle

import (
	"context"
	constance "novel-launch/novel/Biu/const"
	"novel-launch/novel/middleware/db"
	authjwt "novel-launch/novel/middleware/jwt"
	redismw "novel-launch/novel/middleware/redis"
	"novel-launch/novel/server/dao"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Token string `json:"token"`
}

func Login(ctx *gin.Context, req LoginRequest) LoginResponse {
	var resp LoginResponse
	name := req.Name
	d := dao.NewUserDAO(db.Get())
	user, err := d.GetUserByUName(ctx, name)
	if err != nil {
		resp.Code = constance.ParamErr
		resp.Msg = constance.ParamErrMsg
		return resp
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		resp.Code = constance.UserPassWordErr
		resp.Msg = constance.UserPassWordFormatErrMsg
		return resp
	}
	token, err := authjwt.GenerateToken(user.UID, user.UName)
	if err != nil {
		resp.Code = constance.ParamServerErr
		resp.Msg = constance.ParamServerErrMsg
		return resp
	}
	user.Token = token
	if r := redismw.Get(); r != nil {
		_ = r.Set(context.Background(), "user:token:"+strconv.FormatUint(user.UID, 10), token, 7*24*time.Hour).Err()
	}
	_ = d.UpdateUser(ctx, user)
	resp.Code = constance.Success
	resp.Msg = constance.SuccessMsg
	resp.Token = token
	return resp
}
