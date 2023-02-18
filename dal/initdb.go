package dal

import (
	"log"

	"github.com/ClubWeGo/videomicro/dal/query"
	"github.com/ClubWeGo/videomicro/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dsn string) {
	utils.RegisterSSH()
	database, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}
	DB = database
	query.SetDefault(DB)
}
