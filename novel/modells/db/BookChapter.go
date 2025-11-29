package models

import (
	"gorm.io/gorm"
	"time"
)

type BookChapter struct {
	ID         int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	BookID     int64  `gorm:"not null" json:"bookID"` // 关联Books的ID表示对应书的章节
	Title      string `gorm:"size:255;not null" json:"title"`
	IsVip      bool   `gorm:"default 0" json:"is_vip"`
	CreateTime int64  `gorm:"not null;default:0" json:"createTime"`
	UpdateTime int64  `gorm:"not null;default:0" json:"updateTime"`
}

func (b *BookChapter) BeforeCreate(tx *gorm.DB) error {
	if b.CreateTime == 0 {
		b.CreateTime = time.Now().Unix()
	}
	if b.UpdateTime == 0 {
		b.UpdateTime = time.Now().Unix()
	}
	return nil
}

func (b *BookChapter) BeforeUpdate(tx *gorm.DB) error {
	b.UpdateTime = time.Now().Unix()
	return nil
}
