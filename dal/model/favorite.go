package model

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model       // 自动创建id, created_at, updated_at, deleted_at
	UserId     int64 `gorm:"not null;index"`
	VideoId    int64 `gorm:"not null;index"`
	AuthorId   int64 `gorm:"not null;index"`
}
