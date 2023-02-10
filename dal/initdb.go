package dal

import (
	"log"

	"github.com/ClubWeGo/usermicro/dal/query"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dsn string) {
	database, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}
	DB = database
	query.SetDefault(DB)
}
