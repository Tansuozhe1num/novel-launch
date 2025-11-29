package controller

import (
	"github.com/gin-gonic/gin"
	"novel-launch/novel/server/handler"
)

func GetRecommendations(c *gin.Context) {
	err, recommendations := handler.GetRecommendations(c)
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "获取小说列表失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg":  "ok",
		"data": recommendations,
	})
}
