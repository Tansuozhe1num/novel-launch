package models

type Category struct {
	GroupID  int64      `gorm:"primaryKey;autoIncrement" json:"group_id"`
	Name     string     `gorm:"size:255;not null" json:"name"`
	ParentID *int64     `json:"parent_id"` // 使用指针，可以为nil表示根节点
	Children []Category `gorm:"foreignKey:ParentID" json:"children"`
}

func (*Category) TableName() string {
	return "category"
}
