package model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	VideoID int64  `gorm:"index:idx_videoid;not null"`
	UserID  int64  `gorm:"index:idx_userid;not null"`
	Content string `gorm:"type:varchar(255);not null"`
}
