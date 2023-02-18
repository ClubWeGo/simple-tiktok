package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model // 自动创建id, created_at, updated_at, deleted_at(用于软删除 Unscoped:https://www.cnblogs.com/guodd/p/14934448.html)

	Name            string `gorm:"type:varchar(32);not null;uniqueIndex"` // Name加唯一索引，用户名不可重复，app业务要求最长32
	Password        string `gorm:"type:varchar(128);not null"`            // md5 128
	Email           string `gorm:"type:varchar(256)"`                     // max valid email 254
	FollowCount     int64
	FollowerCount   int64
	Avatar          string // 头像地址url
	BackgroundImage string // 背景图地址
	Signature       string // 个人简介
	TotalFavorited  int64  // 获赞数量
	WorkCount       int64  // 作品数
	FavoriteCount   int64  // 喜欢数
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
