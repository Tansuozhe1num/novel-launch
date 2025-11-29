package controller

import (
	"github.com/gin-gonic/gin"
	"novel-launch/novel/Biu"
	constance "novel-launch/novel/Biu/const"
	"novel-launch/novel/server/handler"
)

func GetNewBooksList(c *gin.Context) {
	err, newBooks := handler.GetNewBooksList(c)
	if err != nil {
		Biu.Failed(c, constance.ParamServerErr, err.Error())
		return
	}

	Biu.Success(c, newBooks)
}
