package bookHandle

import (
	"novel-launch/novel/middleware/db"
	"novel-launch/novel/server/dao"

	"github.com/gin-gonic/gin"
)

type RankItem struct {
	ID       int64
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
			ID:       b.ID,
			BookName: b.BookName,
			Author:   b.Author,
			Heat:     b.Heat,
		})
	}
	return nil, ret
}
