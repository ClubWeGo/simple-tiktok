package model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	VideoID int    `gorm:"index:idx_videoid;not null"`
	UserID  int    `gorm:"index:idx_userid;not null"`
	Content string `gorm:"type:varchar(255);not null"`
}
