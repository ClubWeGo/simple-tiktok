package main

import (
	"github.com/ClubWeGo/usermicro/dal/model"
	"github.com/ClubWeGo/usermicro/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"log"
)

func main() {

	g := gen.NewGenerator(gen.Config{
		OutPath: "../../dal/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	dsn := "root:yutian@mysql+ssh(127.0.0.1:3306)/simpletk?charset=utf8&parseTime=True&loc=Local"

	utils.RegisterSSH()
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal(err)
		return
	}

	g.UseDB(db)

	g.ApplyBasic(model.User{})

	// g.ApplyInterface(func(model.UserMethod) {}, model.User{})

	g.Execute()
}
