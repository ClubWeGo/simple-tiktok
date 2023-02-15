package model

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model      // 自动创建id, created_at, updated_at, deleted_at(用于软删除 Unscoped:https://www.cnblogs.com/guodd/p/14934448.html)
	UserId     uint `gorm:"not null"`
	User       User
	VideoId    uint `gorm:"not null"`
	Video      Video
}
