package book

import (
	"github.com/gin-gonic/gin"
	"novel-launch/novel/Biu"
	constance "novel-launch/novel/Biu/const"
	"novel-launch/novel/server/handler/bookHandle"
)

func GetRecommendations(ctx *gin.Context) {
	recommendations, err := bookHandle.GetRecommendations(ctx)
	if err != nil {
		Biu.Failed(ctx, constance.ParamServerErr, err.Error())
		return
	}
	Biu.Success(ctx, recommendations)
}
