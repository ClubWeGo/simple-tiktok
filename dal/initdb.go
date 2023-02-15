package dal

import (
	"github.com/ClubWeGo/favoritemicro/dal/model"
	"github.com/ClubWeGo/favoritemicro/dal/query"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// Dynamic SQL
/*type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}*/

func Init() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./dal/query",
		Mode:    gen.WithDefaultQuery, //| gen.WithQueryInterface, gen.WithoutContext | // generate mode
	})
	RegisterSSH()
	db, _ := gorm.Open(mysql.Open("root:yutian@mysql+ssh(127.0.0.1:3306)/simpletk?charset=utf8mb4&parseTime=True&loc=Local"))
	err := db.AutoMigrate(&model.Favorite{})
	if err != nil {
		return
	}
	g.UseDB(db) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(model.Favorite{}, model.User{}, model.Video{})

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	//g.ApplyInterface(func(Querier){}, model.User{}, model.Company{})

	// Generate the code
	g.Execute()
	query.SetDefault(db)
}
