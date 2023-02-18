package main

import (
	"context"
	"github.com/ClubWeGo/favoritemicro/cmd/rpc"
	"github.com/ClubWeGo/favoritemicro/dal/db"
	favorite "github.com/ClubWeGo/favoritemicro/kitex_gen/favorite"
	"github.com/ClubWeGo/favoritemicro/pack"
	"github.com/ClubWeGo/favoritemicro/pkg/errno"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteMethod implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteMethod(ctx context.Context, request *favorite.FavoriteReq) (resp *favorite.FavoriteResp, err error) {
	if request.ActionType == 1 {
		authorId, err := rpc.GetAuthorId(ctx, request.VideoId)
		if err != nil {
			resp.BaseResp = pack.BuildBaseResp(errno.RPCErr.WithMessage(err.Error()))
			return resp, nil
		}
		err = db.AddFavorite(ctx, request.UserId, request.VideoId, authorId)
		if err != nil {
			resp.BaseResp = pack.BuildBaseResp(errno.DBErr.WithMessage(err.Error()))
			return resp, nil
		}
		resp.BaseResp = pack.BuildBaseResp(errno.Success.WithMessage("点赞成功"))
		return resp, nil
	} else if request.ActionType == 2 {
		if err != nil {
			resp.BaseResp = pack.BuildBaseResp(errno.DBErr.WithMessage(err.Error()))
			return resp, nil
		}
		resp.BaseResp = pack.BuildBaseResp(errno.Success.WithMessage("取消点赞"))
		return resp, nil
	} else {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
}

// FavoriteListMethod implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteListMethod(ctx context.Context, request *favorite.FavoriteListReq) (resp *favorite.FavoriteListResp, err error) {
	favoriteList, err := db.GetFavoriteList(ctx, request.UserId)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.DBErr.WithMessage(err.Error()))
		return resp, nil
	}
	videoIdList := pack.Favorites(favoriteList)
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoIdList = videoIdList
	return resp, nil
}

// FavoriteRelationMethod implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteRelationMethod(ctx context.Context, request *favorite.FavoriteRelationReq) (resp *favorite.FavoriteRelationResp, err error) {
	cnt, err := db.GetFavoriteRelation(ctx, request.UserId, request.VideoId)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.DBErr.WithMessage(err.Error()))
		resp.IsFavorite = false
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.IsFavorite = cnt > 0
	return resp, nil
}

// UserFavoriteCountMethod implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) UserFavoriteCountMethod(ctx context.Context, request *favorite.UserFavoriteCountReq) (resp *favorite.UserFavoriteCountResp, err error) {
	favoriteCnt, favoritedCnt, err := db.CountUserFavorite(ctx, request.UserId)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.DBErr.WithMessage(err.Error()))
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.FavoriteCount = favoriteCnt
	resp.FavoritedCount = favoritedCnt
	return resp, nil
}

// VideoFavoriteCountMethod implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) VideoFavoriteCountMethod(ctx context.Context, request *favorite.VideoFavoriteCountReq) (resp *favorite.VideoFavoriteCountResp, err error) {
	cnt, err := db.CountVideoFavorite(ctx, request.VideoId)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.DBErr.WithMessage(err.Error()))
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.FavoriteCount = cnt
	return resp, nil
}
