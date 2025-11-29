package controller

import (
	"github.com/gin-gonic/gin"
	"novel-launch/novel/Biu"
	constance "novel-launch/novel/Biu/const"
	"novel-launch/novel/server/handler"
)

func GetRecommendations(ctx *gin.Context) {
	recommendations, err := handler.GetRecommendations(ctx)
	if err != nil {
		Biu.Failed(ctx, constance.ParamServerErr, err.Error())
		return
	}
	Biu.Success(ctx, recommendations)
}
