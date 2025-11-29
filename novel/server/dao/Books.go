package dao

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	models "novel-launch/novel/modells/db"
)

type BookDAO struct {
	db *gorm.DB
}

func NewBookDAO(db *gorm.DB) *BookDAO {
	return &BookDAO{db: db}
}

func (dao *BookDAO) GetRecommendations(ctx *gin.Context, limit int) ([]models.Books, error) {
	var books []models.Books
	err := dao.db.Where("is_deleted = ?", 0).
		Order("heat DESC").
		Limit(limit).
		Find(&books).Error
	return books, err
}

func (dao *BookDAO) GetHotRanking(ctx *gin.Context, limit int) ([]models.Books, error) {
	var books []models.Books
	err := dao.db.Where("is_deleted = ?", 0).
		Select("id, book_name, author, heat").
		Order("heat DESC").
		Limit(limit).
		Find(&books).Error
	return books, err
}

func (d *BookDAO) GetNewBooks(ctx *gin.Context, limit int) ([]models.Books, error) {
	var books []models.Books
	err := d.db.Where("is_deleted = ?", 0).
		Order("create_time DESC").
		Limit(limit).
		Find(&books).Error
	return books, err
}

func (d *BookDAO) GetBookById(ctx *gin.Context, bookId int64) (models.Books, error) {
	var book models.Books
	err := d.db.Where("id = ?", bookId).
		First(&book).Error
	return book, err
}
