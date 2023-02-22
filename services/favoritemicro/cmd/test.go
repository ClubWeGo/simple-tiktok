package main

import (
	"github.com/ClubWeGo/simple-tiktok/services/favoritemicro/dal"
)

func main() {
	dal.Init()
	dal.InitRedis()
	//err := db.AddFavorite(nil, 1, 1)
	//if err != nil {
	//	return
	//}

}
