package kitex_server

import (
	"context"
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

func GetFavoriteList(ctx context.Context, uid int64) (favorites []*interaction.FavoriteListReq, err error) {
	//res, err := FavoriteClient.FavoriteListMethod(ctx, &favorite.FavoriteListReq{
	//	UserId: uid,
	//})
	//Videoclient.GetVideoMethod(ctx, &videomicro.GetVideoReq{})
	//if err != nil {
	//	return nil, errno.RPCErr
	//}

	return
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
