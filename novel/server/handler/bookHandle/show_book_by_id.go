package bookHandle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	constance "novel-launch/novel/Biu/const"
	"novel-launch/novel/middleware/db"
	"novel-launch/novel/server/dao"
)

type ShowBookRequest struct {
	BookId int64 `form:"bookId"`
}

type ShowBookResponse struct {
	BookName   string `json:"bookName"`
	Author     string `json:"author"`
	Desc       string `json:"desc"`
	Tag        string `json:"tag"`
	Cover      string `json:"cover"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
	Status     int64  `json:"status"`
	IsDeleted  int64  `json:"isDeleted"`
	Heat       int64  `json:"heat"`
}

func ShowBookById(ctx *gin.Context, req ShowBookRequest) (code int, msg string, response *ShowBookResponse) {
	book, err := dao.NewBookDAO(db.Get()).GetBookById(ctx, req.BookId)
	if err != nil {
		fmt.Printf("ERROR: %x, GG, %d", book, req.BookId)
		return constance.ParamServerErr, constance.ParamServerErrMsg, response
	}

	// TODO redis 记录点击数， 写一个脚本定时刷一遍redis数据到数据库并删除redis数据，更新heat值
	response = &ShowBookResponse{
		BookName:   book.BookName,
		Author:     book.Author,
		Desc:       book.Desc,
		Tag:        book.Tag,
		Cover:      book.Cover,
		CreateTime: book.CreateTime,
		UpdateTime: book.UpdateTime,
		Status:     book.Status,
		IsDeleted:  book.IsDeleted,
		Heat:       book.Heat,
	}
	return constance.Success, constance.SuccessMsg, response
}
