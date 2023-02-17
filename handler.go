package main

import (
	"context"
	"fmt"
	"github.com/ClubWeGo/favoritemicro/dal/db"
	favorite "github.com/ClubWeGo/favoritemicro/kitex_gen/favorite"
	"log"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteMethod implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteMethod(ctx context.Context, request *favorite.FavoriteReq) (resp *favorite.FavoriteResp, err error) {
	// TODO: Your code here...
	//query.Favorite.Create(&query.FavoriteModel{request.Token, request.VideoId})
	//if request.ActionType == 1 {
	fmt.Println("add favorite")
	err = db.AddFavorite(ctx, 1, 1)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//query.Favorite.Create(&query.FavoriteModel{request.Token, request.VideoId})
	//} else {
	//query.Favorite.Delete(&query.FavoriteModel{request.Token, request.VideoId})
	//}
	return
}

// FavoriteListMethod implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteListMethod(ctx context.Context, request *favorite.FavoriteListReq) (resp *favorite.FavoriteListResp, err error) {
	// TODO: Your code here...
	return
}
