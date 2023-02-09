package main

import (
	"log"
	"usermicro/dal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {

	g := gen.NewGenerator(gen.Config{
		OutPath: "../../dal/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	dsn := "tk:123456@tcp(127.0.0.1:3306)/simpletk?charset=utf8&parseTime=True&loc=Local"
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
