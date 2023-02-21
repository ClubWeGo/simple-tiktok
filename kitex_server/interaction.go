package kitex_server

import (
	"context"
	"github.com/ClubWeGo/douyin/pack"
	"github.com/ClubWeGo/usermicro/kitex_gen/usermicro"
	"github.com/ClubWeGo/videomicro/kitex_gen/videomicro"

	"github.com/ClubWeGo/douyin/biz/model/interaction"
	"github.com/ClubWeGo/douyin/tools/errno"
	"github.com/ClubWeGo/favoritemicro/kitex_gen/favorite"
)

func AddFavorite(ctx context.Context, uid int64, vid int64) (*interaction.FavoriteResp, error) {
	resp, err := FavoriteClient.FavoriteMethod(ctx, &favorite.FavoriteReq{
		UserId:     uid,
		VideoId:    vid,
		ActionType: 1,
	})
	if err != nil {
		return nil, errno.RPCErr
	}
	return &interaction.FavoriteResp{
		StatusCode: resp.BaseResp.StatusCode,
		StatusMsg:  resp.BaseResp.StatusMsg,
	}, nil
}

func DeleteFavorite(ctx context.Context, uid int64, vid int64) (*interaction.FavoriteResp, error) {
	resp, err := FavoriteClient.FavoriteMethod(ctx, &favorite.FavoriteReq{
		UserId:     uid,
		VideoId:    vid,
		ActionType: 2,
	})
	if err != nil {
		return nil, errno.RPCErr
	}
	return &interaction.FavoriteResp{
		StatusCode: resp.BaseResp.StatusCode,
		StatusMsg:  resp.BaseResp.StatusMsg,
	}, nil
}

func GetFavoriteList(ctx context.Context, uid int64) (favorites *interaction.FavoriteListResp, err error) {
	favoriteListRes, err := FavoriteClient.FavoriteListMethod(ctx, &favorite.FavoriteListReq{
		UserId: uid,
	})
	if err != nil {
		return nil, err
	}
	favoriteIdList := favoriteListRes.VideoIdList
	videosResp, err := Videoclient.GetVideoSetByIdSetMethod(
		ctx, &videomicro.GetVideoSetByIdSetReq{
			IdSet: favoriteIdList,
		})
	if err != nil {
		return nil, err
	}
	videos := videosResp.VideoSet
	authorIdList := make([]int64, 0)
	for _, v := range videos {
		authorIdList = append(authorIdList, v.AuthorId)
	}

	if err != nil {
		return nil, err
	}
	resp, err := Userclient.GetUserSetByIdSetMethod(ctx, &usermicro.GetUserSetByIdSetReq{
		IdSet: authorIdList,
	})
	if err != nil {
		return nil, errno.RPCErr
	}
	authors := resp.UserSet
	isFavorites := make(map[int64]bool)
	//isFollows := make(map[int64]bool)
	isFollowSet, err := GetIsFollowSetByUserIdSet(uid, authorIdList)
	if err != nil {
		return nil, err
	}
	for _, v := range videos {
		isFavorites[v.Id], _ = GetFavoriteRelation(ctx, uid, v.Id)
		//isFollows[v.AuthorId] = Get
	}
	return &interaction.FavoriteListResp{
		StatusCode: favoriteListRes.BaseResp.StatusCode,
		StatusMsg:  favoriteListRes.BaseResp.StatusMsg,
		VideoList:  pack.Videos(videos, authors, isFavorites, isFollowSet),
	}, nil
}

func GetFavoriteRelation(ctx context.Context, uid int64, vid int64) (bool, error) {
	res, err := FavoriteClient.FavoriteRelationMethod(ctx, &favorite.FavoriteRelationReq{
		UserId:  uid,
		VideoId: vid,
	})
	if err != nil {
		return false, errno.RPCErr
	}
	return res.IsFavorite, nil
}

func CountVideoFavorite(ctx context.Context, vid int64) (cnt int64, err error) {
	res, err := FavoriteClient.VideoFavoriteCountMethod(ctx, &favorite.VideoFavoriteCountReq{
		VideoId: vid,
	})
	if err != nil {
		return 0, errno.RPCErr
	}
	return res.FavoriteCount, nil
}

func CountUserFavorite(ctx context.Context, uid int64) (int64, int64, error) {
	res, err := FavoriteClient.UserFavoriteCountMethod(ctx, &favorite.UserFavoriteCountReq{
		UserId: uid,
	})
	if err != nil {
		return 0, 0, errno.RPCErr
	}
	return res.FavoriteCount, res.FavoritedCount, nil
}

// TODO : 传入userId切片，批量查询user对应的favorite， total_favorited
func GetFavoriteCountByUserIdSet(idSet []int64) (favoriteSet, favoritedSet []int64, err error) {
	return []int64{}, []int64{}, nil
}
