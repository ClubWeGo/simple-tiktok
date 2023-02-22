package db

import (
	"context"
	"github.com/ClubWeGo/simple-tiktok/services/favoritemicro/dal"
	"github.com/ClubWeGo/simple-tiktok/services/favoritemicro/dal/model"
	"github.com/ClubWeGo/simple-tiktok/services/favoritemicro/pkg/errno"
	"strconv"
)

// GetFavoriteRelation get favorite video info
func GetFavoriteRelation(ctx context.Context, uid int64, vid int64) (bool, error) {
	cnt, err := Favorite.WithContext(ctx).Where(Favorite.UserId.Eq(uid), Favorite.VideoId.Eq(vid)).Count()
	if err != nil {
		return false, errno.DBErr.WithMessage(err.Error())
	}
	return cnt > 0, nil
}

func GetFavoriteRelations(ctx context.Context, uid int64, vids []int64) (map[int64]bool, error) {
	res := make(map[int64]bool)
	for _, vid := range vids {
		status, err := GetFavoriteRelation(ctx, uid, vid)
		if err != nil {
			return nil, err
		}
		res[vid] = status
	}
	return res, nil
}

func AddFavorite(ctx context.Context, uid int64, vid int64, aid int64) error {
	count, err := Favorite.WithContext(ctx).Where(Favorite.UserId.Eq(uid), Favorite.VideoId.Eq(vid)).Count()
	if err != nil {
		return errno.DBErr.WithMessage(err.Error())
	}
	if count > 0 {
		return errno.RecordAlreadyExistErr.WithMessage("用户已经点赞过了")
	}
	err = Favorite.WithContext(ctx).Create(&model.Favorite{
		UserId:   uid,
		VideoId:  vid,
		AuthorId: aid,
	})
	if err != nil {
		return errno.DBErr.WithMessage(err.Error())
	}
	_, err = dal.Redis.SAdd(ctx, strconv.FormatInt(uid, 10), vid).Result()
	if err != nil {
		return errno.DBErr.WithMessage(err.Error())
	}
	return nil
}

func DeleteFavorite(ctx context.Context, uid int64, vid int64) error {
	res, err := Favorite.WithContext(ctx).Where(Favorite.UserId.Eq(uid), Favorite.VideoId.Eq(vid)).Delete()
	if err != nil {
		return errno.DBErr.WithMessage(err.Error())
	}
	if res.RowsAffected == 0 {
		return errno.RecordNotExistErr.WithMessage("重复取消点赞")
	}
	_, err = dal.Redis.SRem(ctx, strconv.FormatInt(uid, 10), vid).Result()
	if err != nil {
		return errno.DBErr.WithMessage(err.Error())
	}
	return nil
}

// GetFavoriteList get favorite video list
func GetFavoriteList(ctx context.Context, uid int64) ([]*model.Favorite, error) {
	favorites, err := Favorite.WithContext(ctx).Where(Favorite.UserId.Eq(uid)).Find()
	if err != nil {
		return nil, errno.DBErr.WithMessage(err.Error())
	}
	return favorites, nil
}

func CountVideoFavorite(ctx context.Context, vid int64) (int64, error) {
	cnt, err := Favorite.WithContext(ctx).Where(Favorite.VideoId.Eq(vid)).Count()
	if err != nil {
		return 0, errno.DBErr.WithMessage(err.Error())
	}
	return cnt, nil
}

func CountUserFavorite(ctx context.Context, uid int64) (int64, int64, error) {
	favoriteCnt, err := Favorite.WithContext(ctx).Where(Favorite.UserId.Eq(uid)).Count()

	if err != nil {
		return 0, 0, errno.DBErr.WithMessage(err.Error())
	}
	favoritedCnt, err := Favorite.WithContext(ctx).Where(Favorite.AuthorId.Eq(uid)).Count()
	if err != nil {
		return 0, 0, errno.DBErr.WithMessage(err.Error())
	}
	return favoriteCnt, favoritedCnt, nil
}

func CountVideosFavorite(ctx context.Context, vid []int64) (map[int64]int64, error) {
	var results []struct {
		VideoId int64
		Count   int64
	}
	err := Favorite.WithContext(ctx).Select(Favorite.VideoId, Favorite.VideoId.Count().As("count")).
		Where(Favorite.VideoId.In(vid...)).Group(Favorite.VideoId).Scan(&results)
	if err != nil {
		return nil, errno.DBErr.WithMessage(err.Error())
	}
	res := make(map[int64]int64)
	for _, r := range results {
		res[r.VideoId] = r.Count
	}
	return res, nil
}

func CountUsersFavorite(ctx context.Context, uid []int64) (map[int64][]int64, error) {
	var favoriteResults []struct {
		UserId        int64
		FavoriteCount int64
	}
	var favoritedResults []struct {
		UserId         int64
		FavoritedCount int64
	}

	err := Favorite.WithContext(ctx).Select(Favorite.UserId, Favorite.UserId.Count().As("favorite_count")).
		Where(Favorite.UserId.In(uid...)).Group(Favorite.UserId).Scan(&favoriteResults)
	if err != nil {
		return nil, errno.DBErr.WithMessage(err.Error())
	}

	err = Favorite.WithContext(ctx).Select(Favorite.AuthorId.As("user_id"), Favorite.AuthorId.Count().As("favorited_count")).
		Where(Favorite.AuthorId.In(uid...)).Group(Favorite.AuthorId).Scan(&favoritedResults)
	if err != nil {
		return nil, errno.DBErr.WithMessage(err.Error())
	}

	res := make(map[int64][]int64)
	for _, r := range favoriteResults {
		res[r.UserId] = []int64{r.FavoriteCount}
	}
	for _, r := range favoritedResults {
		res[r.UserId] = append(res[r.UserId], r.FavoritedCount)
	}
	return res, nil
}
