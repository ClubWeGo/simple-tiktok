package main

import "github.com/ClubWeGo/favoritemicro/dal"
import "github.com/ClubWeGo/favoritemicro/dal/db"

func main() {
	dal.Init()
	err := db.AddFavorite(nil, 1, 1)
	if err != nil {
		return
	}

}
