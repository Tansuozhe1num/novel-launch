package controller

import (
	"github.com/gin-gonic/gin"
	"novel-launch/novel/Biu"
	"novel-launch/novel/Biu/const"
	"novel-launch/novel/server/handler"
)

func GetRankLists(c *gin.Context) {
	err, rankList := handler.GetRankList(c, constance.RankCount)
	if err != nil {
		Biu.Failed(c, constance.ParamServerErr, err.Error())
		return
	}
	Biu.Success(c, rankList)
}
