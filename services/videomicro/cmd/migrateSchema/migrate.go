package main

import (
	"log"

<<<<<<< HEAD
	"github.com/ClubWeGo/usermicro/dal/model"
=======
	"github.com/ClubWeGo/videomicro/dal/model"
>>>>>>> video

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

<<<<<<< HEAD
func InitUser(db *gorm.DB) {
	err := db.AutoMigrate(&model.User{})
=======
func InitVideo(db *gorm.DB) {
	err := db.AutoMigrate(&model.Video{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&model.VideoCount{})
>>>>>>> video
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// utils.RegisterSSH()
	var datetimePrecision = 2
<<<<<<< HEAD
	// dsn := "root:yutian@mysql+ssh(127.0.0.1:3306)/simpletk?charset=utf8&parseTime=True&loc=Local"
	dsn := "tk:123456@tcp(127.0.0.1:3306)/simpletk?charset=utf8&parseTime=True&loc=Local"
=======
	dsn := "root:yutian@mysql+ssh(127.0.0.1:3306)/simpletk?charset=utf8&parseTime=True&loc=Local"
	// dsn := "tk:123456@tcp(127.0.0.1:3306)/simpletk?charset=utf8&parseTime=True&loc=Local"
>>>>>>> video
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,                // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
		DefaultStringSize:         256,                // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
		DisableDatetimePrecision:  true,               // disable datetime precision support, which not supported before MySQL 5.6
		DefaultDatetimePrecision:  &datetimePrecision, // default datetime precision
		DontSupportRenameIndex:    true,               // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,               // use change when rename column, rename rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,              // smart configure based on used version
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}

<<<<<<< HEAD
	InitUser(db)
=======
	InitVideo(db)
>>>>>>> video
}
