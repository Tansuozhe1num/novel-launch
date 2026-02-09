package dao

import (
	"context"
	"novel-launch/novel/middleware/db"
	models "novel-launch/novel/modells/db"
)

type CategoriesServe interface {
	GetChildren() []Categories
	FindFather() (models.Category, error)
	GetBooks(ctx context.Context) ([]models.Books, error)
}

type Categories struct {
	// mu       sync.RWMutex
	GroupID  int64        `json:"group_id"`
	Name     string       `json:"name"`
	ParentID *int64       `json:"parent_id"`
	Children []Categories `json:"children"`
}

func NewCategoriesServe(ctx context.Context, groupID int64) CategoriesServe {
	var ret Categories
	err := db.Get().Where("group_id = ?", groupID).First(&ret).Error
	if err != nil {
		return nil
	}
	return &ret
}

func GetDefaultCategories(ctx context.Context) ([]models.Category, error) {
	var ret []models.Category
	err := db.Get().Where("parent_id IS NULL").Find(&ret).Error
	return ret, err
}

func (c *Categories) GetChildren() []Categories {
	return c.Children
}

func (c *Categories) FindFather() (models.Category, error) {
	var ret models.Category
	err := db.Get().Where("group_id = ?", c.GroupID).First(&ret).Error
	return ret, err
}

func (c *Categories) GetBooks(ctx context.Context) ([]models.Books, error) {
	var books []models.Books
	err := db.Get().Where("category_id = ?", c.GroupID).Find(&books).Error
	return books, err
}
