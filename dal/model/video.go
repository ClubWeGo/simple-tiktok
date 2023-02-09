package model

import (
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model // 自动创建id, created_at, updated_at, deleted_at(用于软删除 Unscoped:https://www.cnblogs.com/guodd/p/14934448.html)

	Title          string `gorm:"type:varchar(128);not null;Index"` // Title加索引，后续查询用
	Author_id      int64  `gorm:"type:varchar(128);not null;index"` // md5 128
	Play_url       string `gorm:"type:varchar(256);not null"`       // max valid email 254
	Cover_url      string `gorm:"type:varchar(256);not null"`
	Favorite_count int64  `gorm:"not null"`
	Comment_count  int64  `gorm:"not null"`
}
