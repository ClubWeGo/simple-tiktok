package model
import(
	"gorm.io/gorm"
)

type comment struct{
	gorm.Model
	videoid int64 `gorm:"type:varchar(128);not null;index"`  \\视频id
	videouserid int64 `gorm:"type:varchar(128);not null;index"` \\视频主id
	commentuserid int64 `gorm:"not null"`
	content	string `gorm:"type:varchar(256)"`
	is_follow bool
}