package kitex_server

import (
	"context"
	"sync"

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

// 传入userId切片，批量查询user对应的favorite， total_favorited
// map[int64][]int64  [FavoriteCount  FavoritedCount]
func GetUsersFavoriteCountMap(idSet []int64, respUsersFavoriteCountMap chan map[int64][]int64, wg *sync.WaitGroup, errChan chan error) {
	defer wg.Done()

	res, err := FavoriteClient.UsersFavoriteCountMethod(context.Background(), &favorite.UsersFavoriteCountReq{
		UserIdList: idSet,
	})
	if err != nil {
		respUsersFavoriteCountMap <- map[int64][]int64{}
		errChan <- err
		return
	}
	respUsersFavoriteCountMap <- res.FavoriteCountMap
	errChan <- nil
}

// 传入videoId切片，批量查询video对应的favorite， favorited
// map[int64]int64  FavoriteCount
func GetVideosFavoriteCountMap(idSet []int64, respVideosFavoriteCountMap chan map[int64]int64, wg *sync.WaitGroup, errChan chan error) {
	defer wg.Done()

	res, err := FavoriteClient.VideosFavoriteCountMethod(context.Background(), &favorite.VideosFavoriteCountReq{
		VideoIdList: idSet,
	})
	if err != nil {
		respVideosFavoriteCountMap <- map[int64]int64{}
		errChan <- err
		return
	}
	respVideosFavoriteCountMap <- res.FavoriteCountMap
	errChan <- nil
}

// 传入videoId切片和当前用户id，批量查询喜欢情况
func GetIsFavoriteMap() (idSet []int64, currentUser int64, respIsFavoriteMap chan map[int64]bool, wg *sync.WaitGroup, errChan chan error) {
	defer wg.Done()

	// res, err := FavoriteClient.FavoriteRelationMethod(context.Background(), &favorite.FavoriteRelationReq{})
	return
}
