package db

import (
	"context"
	"github.com/ClubWeGo/favoritemicro/dal/model"
	. "github.com/ClubWeGo/favoritemicro/dal/query"
	"github.com/ClubWeGo/favoritemicro/pkg/errno"
)

// GetFavoriteRelation get favorite video info
func GetFavoriteRelation(ctx context.Context, uid int64, vid int64) (bool, error) {
	cnt, err := Favorite.WithContext(ctx).Where(Favorite.UserId.Eq(uid), Favorite.VideoId.Eq(vid)).Count()
	if err != nil {
		return false, errno.DBErr.WithMessage(err.Error())
	}
	return cnt > 0, nil
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
