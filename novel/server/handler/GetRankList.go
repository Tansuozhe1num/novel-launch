package handler

import (
	"github.com/gin-gonic/gin"
	"novel-launch/novel/middleware/db"
	"novel-launch/novel/server/dao"
)

type RankItem struct {
	BookName string
	Author   string
	Heat     int64
}

func GetRankList(ctx *gin.Context, rankCount int) (error, []RankItem) {
	d := dao.NewBookDAO(db.Get())
	books, err := d.GetHotRanking(ctx, rankCount)
	if err != nil {
		return err, nil
	}
	var ret []RankItem
	for _, b := range books {
		ret = append(ret, RankItem{
			BookName: b.BookName,
			Author:   b.Author,
			Heat:     b.Heat,
		})
	}
	return nil, ret
}
