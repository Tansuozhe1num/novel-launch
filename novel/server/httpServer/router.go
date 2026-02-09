package httpServer

import (
	"novel-launch/novel/controller/book"
	"novel-launch/novel/controller/user"
	authjwt "novel-launch/novel/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	DefaultPage := r.Group("/defaultPage")
	{
		DefaultPage.GET("/recommendations", book.GetRecommendations)
		DefaultPage.GET("/rankings", book.GetRankLists)
		DefaultPage.GET("/newBooksList", book.GetNewBooksList)
	}

	Book := r.Group("/book") // 读书操作
	{
		Book.GET("/showBookById", book.ShowBookById)
		Book.GET("/listChapters", book.ListChapters)
		Book.GET("/chapterContent", book.GetChapterContent)
	}

	User := r.Group("/user")
	{
		User.POST("/login", user.Login)
		User.POST("/register", user.Register)
		User.GET("/profile", authjwt.AuthJWT(), user.Profile)
		User.POST("/logout", authjwt.AuthJWT(), user.Logout)
	}

	Category := r.Group("/category")
	{
		Category.GET("/list", book.GetDefaultCategories)
	}

	r.Static("/static", "../fronter")
	r.StaticFile("/", "../fronter/index.html")
	r.StaticFile("/book.html", "../fronter/book.html")
	r.StaticFile("/page.html", "../fronter/page.html")
	r.StaticFile("/readercentor.html", "../fronter/readercentor.html")
	return r
}
