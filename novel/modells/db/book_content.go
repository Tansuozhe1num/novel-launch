package models

import (
	"gorm.io/gorm"
	"time"
)

type BookContent struct {
	ID         int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	BookID     int64  `gorm:"not null" json:"bookID"`    // 关联Books的ID
	ChapterID  int64  `gorm:"not null" json:"chapterID"` // 关联Chapter的ID
	Content    string `gorm:"type:MEDIUMTEXT" json:"content"`
	CreateTime int64  `gorm:"not null;default:0" json:"createTime"`
	UpdateTime int64  `gorm:"not null;default:0" json:"updateTime"`
}

func (*BookContent) TableName() string {
	return "book_content"
}

func (b *BookContent) BeforeCreate(tx *gorm.DB) error {
	if b.CreateTime == 0 {
		b.CreateTime = time.Now().Unix()
	}
	if b.UpdateTime == 0 {
		b.UpdateTime = time.Now().Unix()
	}
	return nil
}

func (b *BookContent) BeforeUpdate(tx *gorm.DB) error {
	b.UpdateTime = time.Now().Unix()
	return nil
}
