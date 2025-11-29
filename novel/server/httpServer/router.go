package httpServer

import (
	"novel-launch/novel/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	DefaultPage := r.Group("/defaultPage")
	{
		DefaultPage.GET("/recommendations", controller.GetRecommendations)
		DefaultPage.GET("/rankings", controller.GetRankLists)
		DefaultPage.GET("/newBooksList", controller.GetNewBooksList)
	}

	r.Static("/static", "../fronter")
	r.StaticFile("/", "../fronter/reader.html")
	return r
}
