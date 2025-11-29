package handler

import (
	"github.com/gin-gonic/gin"
	"novel-launch/novel/middleware/db"
	models "novel-launch/novel/modells/db"
	"novel-launch/novel/server/dao"
	"strconv"
)

func GetNewBooksList(ctx *gin.Context) (error, []models.Books) {
	countstr := ctx.DefaultQuery("bookCount", "10")
	count, err := strconv.Atoi(countstr)
	if err != nil {
		count = 10
	}

	d := dao.NewBookDAO(db.Get())
	books, err := d.GetNewBooks(ctx, count)
	if err != nil {
		return err, nil
	}

	return nil, books
}
