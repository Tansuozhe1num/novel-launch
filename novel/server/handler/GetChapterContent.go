package handler

import (
    "novel-launch/novel/middleware/db"
    "novel-launch/novel/server/dao"
    "github.com/gin-gonic/gin"
)

type ChapterContentRequest struct {
    BookId    int64 `form:"bookId" json:"bookId"`
    ChapterId int64 `form:"chapterId" json:"chapterId"`
}

type ChapterContentResponse struct {
    ChapterId int64  `json:"chapterId"`
    Content   string `json:"content"`
}

func GetChapterContent(ctx *gin.Context, req ChapterContentRequest) (error, *ChapterContentResponse) {
    d := dao.NewContentDAO(db.Get())
    c, err := d.GetByChapter(ctx, req.BookId, req.ChapterId)
    if err != nil {
        return err, nil
    }
    return nil, &ChapterContentResponse{ChapterId: c.ChapterID, Content: c.Content}
}
