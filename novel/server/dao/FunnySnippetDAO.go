package dao

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	models "novel-launch/novel/modells/db"
)

type FunnySnippetDAO struct {
	db *gorm.DB
}

func NewFunnySnippetDAO(db *gorm.DB) *FunnySnippetDAO {
	return &FunnySnippetDAO{db: db}
}

func (d *FunnySnippetDAO) GetRandomSnippetList(ctx *gin.Context, count int) ([]models.FunnySnippet, error) {
	var snippets []models.FunnySnippet
	if count <= 0 {
		count = 5
	}
	if count >= 100 {
		return snippets, fmt.Errorf("count too large")
	}
	err := d.db.WithContext(ctx).
		Order("RAND()").
		Limit(count).
		Find(&snippets).Error
	return snippets, err
}
