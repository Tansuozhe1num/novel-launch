package handler

import (
	"novel-launch/novel/middleware/db"
	"novel-launch/novel/server/dao"

	"github.com/gin-gonic/gin"
)

type ChapterItem struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	IsVIP int    `json:"isVip"`
}

func ListChapters(ctx *gin.Context, bookId int64) (error, []ChapterItem) {
	d := dao.NewChapterDAO(db.Get())
	list, err := d.ListByBook(ctx, bookId)
	if err != nil {
		return err, nil
	}
	ret := make([]ChapterItem, 0, len(list))
	for _, c := range list {
		ret = append(ret, ChapterItem{ID: c.ID, Title: c.Title, IsVIP: func(b bool) int {
			if b {
				return 1
			}
			return 0
		}(c.IsVip)})
	}
	return nil, ret
}
