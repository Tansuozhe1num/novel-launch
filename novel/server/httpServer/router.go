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

	Book := r.Group("/book") // 读书操作
	{
		Book.GET("/showBookById", controller.ShowBookById)
		Book.GET("/listChapters", controller.ListChapters)
	}

	r.Static("/static", "../fronter")
	r.StaticFile("/", "../fronter/reader.html")
	r.StaticFile("/book.html", "../fronter/book.html")
	return r
}
