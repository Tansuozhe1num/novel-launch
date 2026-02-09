package book

import (
	"github.com/gin-gonic/gin"
	"novel-launch/novel/Biu"
	"novel-launch/novel/server/handler/category"
)

func GetDefaultCategories(ctx *gin.Context) {
	code, msg, data, err := category.GetDefaultCategories(ctx)
	if err != nil {
		Biu.Failed(ctx, code, msg)
		return
	}
	Biu.Success(ctx, data)
}
