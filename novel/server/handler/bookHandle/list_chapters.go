package bookHandle

import (
	"fmt"
	"novel-launch/novel/middleware/db"
	"novel-launch/novel/server/dao"

	"github.com/gin-gonic/gin"
)

type ChapterItem struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	IsVIP int    `json:"isVip"`
}
type ListChapterReq struct {
	BookId int64 `json:"bookId" form:"bookId"`
}

func ListChapters(ctx *gin.Context, req ListChapterReq) (error, []ChapterItem) {
	bookId := req.BookId
	if bookId == 0 {
		return fmt.Errorf("param bookId is required"), nil
	}

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
