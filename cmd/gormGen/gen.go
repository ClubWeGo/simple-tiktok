package main

import (
	"github.com/ClubWeGo/commentmicro/utils"
	"log"

	"github.com/ClubWeGo/commentmicro/dal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {

	g := gen.NewGenerator(gen.Config{
		OutPath: "../../dal/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	//dsn := "root:12345678@tcp(127.0.0.1:3306)/simpletk?charset=utf8&parseTime=True&loc=Local"
	dsn := "root:yutian@tcp(124.221.147.131:3306)/simpletk?charset=utf8&parseTime=True&loc=Local"

	utils.RegisterSSH()
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal(err)
		return
	}

	g.UseDB(db)

	g.ApplyBasic(model.Comment{})

	// g.ApplyInterface(func(model.UserMethod) {}, model.User{})

	g.Execute()
}
