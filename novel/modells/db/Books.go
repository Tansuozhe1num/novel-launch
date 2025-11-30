package models

import (
	"gorm.io/gorm"
	"time"
)

type Books struct {
	ID         int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	BookName   string `gorm:"size:255;not null" json:"bookName"`
	Author     string `gorm:"size:100;not null" json:"author"`
	Desc       string `gorm:"type:text" json:"desc"`
	Tag        string `gorm:"size:100;index:idx_tag" json:"tag"`
	Cover      string `gorm:"size:500" json:"cover"`
	Status     int64  `gorm:"not null;default:0;index:idx_status" json:"status"`
	IsDeleted  int64  `gorm:"not null;default:0" json:"isDeleted"`
	Heat       int64  `gorm:"not null;default:0;index:idx_heat" json:"heat"`
	OnClick    int64  `gorm:"not null;default:0" json:"onClick"`
	CreateTime int64  `gorm:"not null;default:0" json:"createTime"`
	UpdateTime int64  `gorm:"not null;default:0" json:"updateTime"`
}

func (*Books) TableName() string {
	return "books"
}

func (b *Books) BeforeCreate(tx *gorm.DB) error {
	if b.CreateTime == 0 {
		b.CreateTime = time.Now().Unix()
	}
	if b.UpdateTime == 0 {
		b.UpdateTime = time.Now().Unix()
	}
	return nil
}

func (b *Books) BeforeUpdate(tx *gorm.DB) error {
	b.UpdateTime = time.Now().Unix()
	return nil
}
