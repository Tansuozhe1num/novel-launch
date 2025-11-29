package controller

import (
	"github.com/gin-gonic/gin"
	"novel-launch/novel/server/handler"
)

func GetNewBooksList(c *gin.Context) {
	err, newBooks := handler.GetNewBooksList(c)
	if err != nil {
		c.JSON(500, gin.H{
			"msg":  "获取最新小说列表失败",
			"data": nil,
		})
	}

	c.JSON(200, gin.H{
		"msg":  "",
		"data": newBooks,
	})
}
