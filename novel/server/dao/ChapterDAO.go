package dao

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	models "novel-launch/novel/modells/db"
)

type ChapterDAO struct {
	db *gorm.DB
}

func NewChapterDAO(db *gorm.DB) *ChapterDAO {
	return &ChapterDAO{db: db}
}

func (d *ChapterDAO) ListByBook(ctx *gin.Context, bookId int64) ([]models.BookChapter, error) {
	var chapters []models.BookChapter
	err := d.db.WithContext(ctx).Model(&models.BookChapter{}).Where("book_id = ?", bookId).Order("id ASC").Find(&chapters).Error
	return chapters, err
}

func (d *ChapterDAO) CountByBook(ctx *gin.Context, bookId int64) (int64, error) {
	var cnt int64
	err := d.db.WithContext(ctx).Model(&models.BookChapter{}).Where("book_id = ?", bookId).Count(&cnt).Error
	return cnt, err
}
