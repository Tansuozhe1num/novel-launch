package dao

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"novel-launch/novel/modells/db"
)

type ContentDAO struct {
	db *gorm.DB
}

func NewContentDAO(db *gorm.DB) *ContentDAO {
	return &ContentDAO{db: db}
}

func (d *ContentDAO) GetByChapter(ctx *gin.Context, bookId int64, chapterId int64) (models.BookContent, error) {
	var c models.BookContent
	err := d.db.WithContext(ctx).Where("book_id = ? AND chapter_id = ?", bookId, chapterId).First(&c).Error
	return c, err
}
