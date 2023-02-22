package rpc

import (
	"context"
	"github.com/ClubWeGo/simple-tiktok/services/favoritemicro/kitex_gen/favorite"
	"log"
)

func Favorite() {
	req := &favorite.FavoriteReq{UserId: 1, VideoId: 1, ActionType: 1}
	res, err := FavoriteClient.FavoriteMethod(context.Background(), req)
	if err != nil {
		log.Println(err)
	}
	log.Println(res)
}
