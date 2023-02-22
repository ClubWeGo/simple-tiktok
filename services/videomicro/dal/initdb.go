package dal

import (
<<<<<<< HEAD
	"github.com/ClubWeGo/usermicro/utils"
	"gorm.io/driver/mysql"

	"github.com/ClubWeGo/usermicro/dal/query"
	"log"

=======
	"log"

	"github.com/ClubWeGo/videomicro/dal/query"
	"github.com/ClubWeGo/videomicro/utils"

	"gorm.io/driver/mysql"
>>>>>>> video
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
