package main

import (
	"context"
	"github.com/ClubWeGo/favoritemicro/dal/db"
	favorite "github.com/ClubWeGo/favoritemicro/kitex_gen/favorite"
	"github.com/ClubWeGo/favoritemicro/pack"
	"log"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteMethod implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteMethod(ctx context.Context, request *favorite.FavoriteReq) (resp *favorite.FavoriteResp, err error) {
	if request.ActionType == 1 {
		err = db.AddFavorite(ctx, uint(request.UserId), uint(request.VideoId))
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return &favorite.FavoriteResp{
			StatusCode: 0,
			StatusMsg:  &FavoriteSuccess,
		}, nil
	} else {
		err = db.DeleteFavorite(ctx, uint(request.UserId), uint(request.VideoId))
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return &favorite.FavoriteResp{
			StatusCode: 0,
			StatusMsg:  &FavoriteSuccess,
		}, nil
	}
}

// FavoriteListMethod implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteListMethod(ctx context.Context, request *favorite.FavoriteListReq) (resp *favorite.FavoriteListResp, err error) {
	videoList, err := db.GetFavoriteList(ctx, uint(request.UserId))
	if err != nil {
		return nil, err
	}
	return &favorite.FavoriteListResp{
		StatusCode: 0,
		//StatusMsg:   &FavoriteSuccess,
		VideoIdList: pack.Favorites(videoList),
	}, nil
}

// FavoriteRelationMethod implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteRelationMethod(ctx context.Context, request *favorite.FavoriteRelationReq) (resp *favorite.FavoriteRelationResp, err error) {
	cnt, err := db.GetFavoriteRelation(ctx, uint(request.UserId), uint(request.VideoId))
	if err != nil {
		return &favorite.FavoriteRelationResp{
			StatusCode: 1,
			IsFavorite: false,
		}, nil
	}
	return &favorite.FavoriteRelationResp{
		StatusCode: 0,
		StatusMsg:  &FavoriteSuccess,
		IsFavorite: cnt > 0,
	}, nil
}

// UserFavoriteCountMethod implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) UserFavoriteCountMethod(ctx context.Context, request *favorite.UserFavoriteCountReq) (resp *favorite.UserFavoriteCountResp, err error) {
	// TODO: Your code here...
	return
}

// VideoFavoriteCountMethod implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) VideoFavoriteCountMethod(ctx context.Context, request *favorite.VideoFavoriteCountReq) (resp *favorite.VideoFavoriteCountResp, err error) {
	// TODO: Your code here...
	return
}
