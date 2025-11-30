package userhandle

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	constance "novel-launch/novel/Biu/const"
	"novel-launch/novel/middleware/db"
	"novel-launch/novel/middleware/sonyflake"
	models "novel-launch/novel/modells/db"
	"novel-launch/novel/server/dao"
	"time"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Register(ctx *gin.Context, req RegisterRequest) (httpCode int, code int, msg string) {
	name := req.Name
	userdb := dao.NewUserDAO(db.Get())
	if user, err := userdb.GetUserByUID(ctx, name); err == nil {
		if user.UName == name {
			return http.StatusOK, constance.UserNameExist, constance.UserNameExistMsg
		}
	}

	if !checkPassword(req.Password) {
		return http.StatusOK, constance.UserPassWordFormatErr, constance.UserPassWordFormatErrMsg
	}

	pwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return http.StatusInternalServerError, constance.ParamServerErr, constance.ParamServerErrMsg
	}

	uid, _ := sonyflake.GenID()
	if err := userdb.CreateUser(ctx, models.User{
		UID:        uid,
		UName:      name,
		Password:   string(pwd),
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}); err != nil {
		return http.StatusServiceUnavailable, constance.ParamServerErr, constance.ParamServerErrMsg
	}

	return http.StatusOK, constance.Success, constance.SuccessMsg
}

func checkPassword(pwd string) bool {
	if len(pwd) < 6 || len(pwd) > 12 {
		return false
	}

	hasDigit := false
	hasLetter := false

	for _, c := range pwd {
		if c >= '0' && c <= '9' {
			hasDigit = true
		} else if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			hasLetter = true
		} else {
			return false
		}
	}

	return hasDigit && hasLetter
}
