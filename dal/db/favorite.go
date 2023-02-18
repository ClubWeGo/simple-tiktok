package db

import (
	"context"
	"github.com/ClubWeGo/favoritemicro/dal/model"
	. "github.com/ClubWeGo/favoritemicro/dal/query"
)

// GetFavoriteRelation get favorite video info
func GetFavoriteRelation(ctx context.Context, uid int64, vid int64) (int64, error) {
	cnt, err := Favorite.WithContext(ctx).Where(Favorite.UserId.Eq(uid), Favorite.VideoId.Eq(vid)).Count()
	if err != nil {
		return 0, err
	}
	return cnt, nil
}

func AddFavorite(ctx context.Context, uid int64, vid int64, aid int64) error {
	/*log.Println("add favorite")
	user, err := User.WithContext(ctx).Where(User.ID.Eq(uid)).First()
	if err != nil {
		return err
	}
	log.Println(user.ID)
	return nil*/

	err := Favorite.WithContext(ctx).Create(&model.Favorite{
		UserId:   uid,
		VideoId:  vid,
		AuthorId: aid,
	})
	if err != nil {
		return err
	}
	return nil
}

func DeleteFavorite(ctx context.Context, uid int64, vid int64) error {
	_, err := Favorite.WithContext(ctx).Where(Favorite.UserId.Eq(uid), Favorite.VideoId.Eq(vid)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// GetFavoriteList get favorite video list
func GetFavoriteList(ctx context.Context, uid int64) ([]*model.Favorite, error) {
	favorites, err := Favorite.WithContext(ctx).Where(Favorite.UserId.Eq(uid)).Find()
	if err != nil {
		return nil, err
	}
	return favorites, nil
}

func CountVideoFavorite(ctx context.Context, vid int64) (int64, error) {
	cnt, err := Favorite.WithContext(ctx).Where(Favorite.VideoId.Eq(vid)).Count()
	if err != nil {
		return 0, err
	}
	return cnt, nil
}

func CountUserFavorite(ctx context.Context, uid int64) (int64, int64, error) {
	favoriteCnt, err := Favorite.WithContext(ctx).Where(Favorite.UserId.Eq(uid)).Count()

	if err != nil {
		return 0, 0, err
	}
	favoritedCnt, err := Favorite.WithContext(ctx).Where(Favorite.AuthorId.Eq(uid)).Count()
	if err != nil {
		return 0, 0, err
	}
	return favoriteCnt, favoritedCnt, nil
}
