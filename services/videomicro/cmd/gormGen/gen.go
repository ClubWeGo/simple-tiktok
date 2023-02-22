package main

import (
	"log"

<<<<<<< HEAD
	"github.com/ClubWeGo/usermicro/dal/model"
=======
	"github.com/ClubWeGo/videomicro/dal/model"
	"github.com/ClubWeGo/videomicro/utils"
>>>>>>> video
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
<<<<<<< HEAD

=======
>>>>>>> video
	g := gen.NewGenerator(gen.Config{
		OutPath: "../../dal/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
<<<<<<< HEAD
	// dsn := "root:yutian@mysql+ssh(127.0.0.1:3306)/simpletk?charset=utf8&parseTime=True&loc=Local"
	dsn := "tk:123456@tcp(127.0.0.1:3306)/simpletk?charset=utf8&parseTime=True&loc=Local"

	// utils.RegisterSSH()
=======
	dsn := "root:yutian@mysql+ssh(127.0.0.1:3306)/simpletk?charset=utf8&parseTime=True&loc=Local"
	// dsn := "tk:123456@tcp(127.0.0.1:3306)/simpletk?charset=utf8&parseTime=True&loc=Local"

	utils.RegisterSSH()
>>>>>>> video
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal(err)
		return
	}

	g.UseDB(db)

<<<<<<< HEAD
	g.ApplyBasic(model.User{})
=======
	g.ApplyBasic(model.Video{})

	g.ApplyBasic(model.VideoCount{})
>>>>>>> video

	// g.ApplyInterface(func(model.UserMethod) {}, model.User{})

	g.Execute()
}
