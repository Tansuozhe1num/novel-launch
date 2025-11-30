package book

import (
	"github.com/gin-gonic/gin"
	"novel-launch/novel/Biu"
	constance "novel-launch/novel/Biu/const"
	"novel-launch/novel/server/handler/bookHandle"
)

func GetNewBooksList(c *gin.Context) {
	err, newBooks := bookHandle.GetNewBooksList(c)
	if err != nil {
		Biu.Failed(c, constance.ParamServerErr, err.Error())
		return
	}

	Biu.Success(c, newBooks)
}
