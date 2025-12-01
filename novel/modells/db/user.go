package models

type User struct {
	ID            int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	UID           uint64 `gorm:"size:255;not null;unique" json:"uid"`
	UName         string `gorm:"size:255;not null" json:"name"`
	Password      string `gorm:"size:255;not null" json:"password"`
	Avatar        string `gorm:"size:255;default null" json:"avatar"`
	Token         string `gorm:"size:255;not null" json:"token"`
	Email         string `gorm:"size:255;default null" json:"email"`
	Status        int    `gorm:"not null;default:0" json:"status"`        // 0 未登录， 1在线， 2 封禁, 3 离线
	UserLikeCount int    `gorm:"not null;default:0" json:"userLikeCount"` // 用户喜欢系数，用来计算推荐算法
	CreateTime    int64  `gorm:"not null;default:0" json:"createTime"`
	UpdateTime    int64  `gorm:"not null;default:0" json:"updateTime"`
}

func (*User) TableName() string {
	return "user"
}
