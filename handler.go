package main

import (
	"context"
	"github.com/ClubWeGo/favoritemicro/cmd/rpc"
	"github.com/ClubWeGo/favoritemicro/dal/rdb"
	favorite "github.com/ClubWeGo/favoritemicro/kitex_gen/favorite"
	"github.com/ClubWeGo/favoritemicro/pack"
	"github.com/ClubWeGo/favoritemicro/pkg/errno"
	"log"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteMethod implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteMethod(ctx context.Context, request *favorite.FavoriteReq) (resp *favorite.FavoriteResp, err error) {
	resp = &favorite.FavoriteResp{}
	if request.ActionType != 1 && request.ActionType != 2 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	// 获取作者id
	var authorId int64
	authorId, _ = rdb.GetAuthorId(ctx, request.VideoId)
	// 如果缓存中没有作者id，就去rpc获取
	if authorId == 0 {
		authorId, err = rpc.GetAuthorId(ctx, request.VideoId)
		if err != nil {
			resp.BaseResp = pack.BuildBaseResp(errno.RPCErr.WithMessage(err.Error()))
			return resp, nil
		}
		// 将作者id存入缓存
		rdb.SetAuthorId(ctx, request.VideoId, authorId)
	}

	if request.ActionType == 1 {
		err = rdb.AddFavorite(ctx, request.UserId, request.VideoId, authorId)
		if err != nil {
			resp.BaseResp = pack.BuildBaseResp(err)
			return resp, nil
		}
		log.Println("点赞成功")
		resp.BaseResp = pack.BuildBaseResp(errno.Success.WithMessage("点赞成功"))
		return resp, nil
	} else {
		err = rdb.DeleteFavorite(ctx, request.UserId, request.VideoId, authorId)
		if err != nil {
			resp.BaseResp = pack.BuildBaseResp(err)
			return resp, nil
		}
		resp.BaseResp = pack.BuildBaseResp(errno.Success.WithMessage("取消点赞"))
		return resp, nil
	}
}

// FavoriteListMethod implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteListMethod(ctx context.Context, request *favorite.FavoriteListReq) (resp *favorite.FavoriteListResp, err error) {
	resp = &favorite.FavoriteListResp{}
	favoriteList, err := rdb.GetFavoriteList(ctx, request.UserId)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.DBErr.WithMessage(err.Error()))
		return resp, nil
	}
	//videoIdList := pack.Favorites(favoriteList)
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoIdList = favoriteList
	return resp, nil
}

// FavoriteRelationMethod implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteRelationMethod(ctx context.Context, request *favorite.FavoriteRelationReq) (resp *favorite.FavoriteRelationResp, err error) {
	resp = &favorite.FavoriteRelationResp{}
	res, err := rdb.GetFavoriteRelation(ctx, request.UserId, request.VideoId)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		resp.IsFavorite = false
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success.WithMessage("已点赞"))
	resp.IsFavorite = res
	return resp, nil
}

// UserFavoriteCountMethod implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) UserFavoriteCountMethod(ctx context.Context, request *favorite.UserFavoriteCountReq) (resp *favorite.UserFavoriteCountResp, err error) {
	resp = &favorite.UserFavoriteCountResp{}
	favoriteCnt, favoritedCnt, err := rdb.CountUserFavorite(ctx, request.UserId)
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
	resp = &favorite.VideoFavoriteCountResp{}
	cnt, err := rdb.CountVideoFavorite(ctx, request.VideoId)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.FavoriteCount = cnt
	return resp, nil
}

// VideosFavoriteCountMethod implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) VideosFavoriteCountMethod(ctx context.Context,
	request *favorite.VideosFavoriteCountReq) (resp *favorite.VideosFavoriteCountResp, err error) {
	resp = &favorite.VideosFavoriteCountResp{}
	favoriteCountMap, err := rdb.CountVideosFavorite(ctx, request.VideoIdList)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.FavoriteCountMap = favoriteCountMap
	return
}

// UsersFavoriteCountMethod implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) UsersFavoriteCountMethod(ctx context.Context,
	request *favorite.UsersFavoriteCountReq) (resp *favorite.UsersFavoriteCountResp, err error) {
	resp = &favorite.UsersFavoriteCountResp{}
	favoriteMap, err := rdb.CountUsersFavorite(ctx, request.UserIdList)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.FavoriteCountMap = favoriteMap
	return
}

// FavoriteRelationsMethod implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteRelationsMethod(ctx context.Context, request *favorite.FavoriteRelationsReq) (resp *favorite.FavoriteRelationsResp, err error) {
	resp = &favorite.FavoriteRelationsResp{}
	isFavoriteMap, err := rdb.GetFavoriteRelations(ctx, request.UserId, request.VideoIdList)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.IsFavoriteMap = isFavoriteMap
	return resp, nil
}
