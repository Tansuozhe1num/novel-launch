package dao

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	models "novel-launch/novel/modells/db"
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{db: db}
}

func (d *UserDAO) GetUserByUID(ctx *gin.Context, uid string) (models.User, error) {
	var user models.User
	err := d.db.WithContext(ctx).Where("uid = ?", uid).First(&user).Error
	return user, err
}

func (d *UserDAO) CreateUser(ctx *gin.Context, user models.User) error {
	return d.db.WithContext(ctx).Create(&user).Error
}

func (d *UserDAO) UpdateUser(ctx *gin.Context, user models.User) error {
	return d.db.WithContext(ctx).Save(&user).Error
}

func (d *UserDAO) ListUser(ctx *gin.Context) ([]models.User, error) {
	var users []models.User
	err := d.db.WithContext(ctx).Find(&users).Error
	return users, err
}
