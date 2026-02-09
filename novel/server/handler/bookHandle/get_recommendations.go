package bookHandle

import (
	"github.com/gin-gonic/gin"
	"novel-launch/novel/middleware/db"
	"novel-launch/novel/server/dao"
)

type Recommendation struct {
	ID     int64
	Title  string
	Author string
	Desc   string
	Tag    string
	Cover  string
	Heat   int
}

func GetRecommendations(ctx *gin.Context) ([]Recommendation, error) {
	var ret []Recommendation
	d := dao.NewBookDAO(db.Get())
	books, err := d.GetRecommendations(ctx, 3)
	if err != nil {
		return nil, err
	}
	for _, b := range books {
		ret = append(ret, Recommendation{
			ID:     b.ID,
			Title:  b.BookName,
			Author: b.Author,
			Desc:   b.Desc,
			Tag:    b.Tag,
			Cover:  b.Cover,
			Heat:   int(b.Heat),
		})
	}
	return ret, nil
}
