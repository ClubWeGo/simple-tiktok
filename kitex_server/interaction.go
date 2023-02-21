package kitex_server

import (
	"context"
	"sync"

	"github.com/ClubWeGo/commentmicro/kitex_gen/comment"
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
	videoIdList := make([]int64, 0)
	authorIdList := make([]int64, 0)
	for _, v := range videos {
		videoIdList = append(videoIdList, v.Id)
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
	isFavoriteMap, err := GetFavoriteRelations(ctx, uid, videoIdList)
	if err != nil {
		return nil, err
	}
	isFollowMap, err := GetIsFollowMapByUserIdSet(uid, authorIdList)
	if err != nil {
		return nil, err
	}
	return &interaction.FavoriteListResp{
		StatusCode: favoriteListRes.BaseResp.StatusCode,
		StatusMsg:  favoriteListRes.BaseResp.StatusMsg,
		VideoList:  pack.Videos(videos, authors, isFavoriteMap, isFollowMap),
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

func GetFavoriteRelations(ctx context.Context, uid int64, vids []int64) (isFavoriteMap map[int64]bool, err error) {
	res, err := FavoriteClient.FavoriteRelationsMethod(ctx, &favorite.FavoriteRelationsReq{
		UserId:      uid,
		VideoIdList: vids,
	})
	if err != nil {
		return nil, errno.RPCErr
	}
	isFavoriteMap = res.IsFavoriteMap
	return isFavoriteMap, nil
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

// 协程接口
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
func GetIsFavoriteMap(idSet []int64, currentUser int64, respIsFavoriteMap chan map[int64]bool, wg *sync.WaitGroup, errChan chan error) {
	defer wg.Done()

	res, err := FavoriteClient.FavoriteRelationsMethod(context.Background(), &favorite.FavoriteRelationsReq{
		UserId:      currentUser,
		VideoIdList: idSet,
	})
	if err != nil {
		respIsFavoriteMap <- map[int64]bool{}
		errChan <- err
		return
	}
	respIsFavoriteMap <- res.IsFavoriteMap
	errChan <- nil
}

// 传入videoId切片，批量视频的评论数量
func GetCommentCountMap(idSet []int64, respIsFavoriteMap chan map[int64]int64, wg *sync.WaitGroup, errChan chan error) {
	defer wg.Done()

	res, err := CommentClient.VideosFavoriteCountMethod(context.Background(), &comment.VideosCommentCountReq{
		VideoIdList: idSet,
	})
	if err != nil {
		respIsFavoriteMap <- map[int64]int64{}
		errChan <- err
		return
	}
	respIsFavoriteMap <- res.CommentCountMap
	errChan <- nil
}
