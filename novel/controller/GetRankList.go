package controller

import (
	"github.com/gin-gonic/gin"
	"novel-launch/novel/Biu/const"
	"novel-launch/novel/server/handler"
)

func GetRankLists(c *gin.Context) {
	err, rankList := handler.GetRankList(c, constance.RankCount)
	if err != nil {
		c.JSON(500, gin.H{
			"msg":  "error",
			"data": nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"msg":  "ok",
		"data": rankList,
	})
}
