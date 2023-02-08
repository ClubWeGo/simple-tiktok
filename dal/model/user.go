package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model // 自动创建id, created_at, updated_at, deleted_at(用于软删除 Unscoped:https://www.cnblogs.com/guodd/p/14934448.html)

	Name           string `gorm:"type:varchar(128);not null;uniqueIndex"` // Name加唯一索引，用户名不可重复
	Password       string `gorm:"type:varchar(128);not null"`             // md5 128
	Email          string `gorm:"type:varchar(256)"`                      // max valid email 254
	Follow_count   int64
	Follower_count int64
	Is_follow      bool
}

// 自定义query
// type UserMethod interface {
// 	// SELECT * FROM @@table WHERE name = @name
// 	GetByName(name, role string) ([]gen.T, error)

// 	// SELECT * FROM @@table WHERE id = @id
// 	GetById(id, role string) ([]gen.T, error)

// 	// SELECT * FROM @@table WHERE email = @email
// 	GetByEmail(email, role string) ([]gen.T, error)
// }
