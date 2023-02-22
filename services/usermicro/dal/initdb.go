package dal

import (
	"github.com/ClubWeGo/usermicro/utils"
	"gorm.io/driver/mysql"

	"github.com/ClubWeGo/usermicro/dal/query"
	"log"

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
