package db

import (
	"context"
	"github.com/ClubWeGo/favoritemicro/dal/model"
	. "github.com/ClubWeGo/favoritemicro/dal/query"
)

// GetFavoriteRelation get favorite video info
func GetFavoriteRelation(ctx context.Context, uid uint, vid uint) error {
	_, err := Favorite.WithContext(ctx).Where(Favorite.UserId.Eq(uid), Favorite.VideoId.Eq(vid)).First()
	if err != nil {
		return err
	}
	return nil
}

func AddFavorite(ctx context.Context, uid uint, vid uint) error {
	/*log.Println("add favorite")
	user, err := User.WithContext(ctx).Where(User.ID.Eq(uid)).First()
	if err != nil {
		return err
	}
	log.Println(user.ID)
	return nil*/

	return Q.Transaction(func(tx *Query) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		//1. 新增点赞数据
		if err := tx.Favorite.WithContext(ctx).Create(&model.Favorite{UserId: uid, VideoId: vid}); err != nil {
			return err
		}
		//2.改变 video 表中的 favorite count
		_, err := tx.Video.WithContext(ctx).Where(tx.Video.ID.Eq(vid)).UpdateSimple(tx.Video.Comment_count.Add(1))
		if err != nil {
			return err
		}
		return nil
	})
}

func DeleteFavorite(ctx context.Context, uid uint, vid uint) error {
	return Q.Transaction(func(tx *Query) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		//1. 删除点赞数据
		if _, err := tx.Favorite.WithContext(ctx).Where(tx.Favorite.UserId.Eq(uid), tx.Favorite.VideoId.Eq(vid)).Delete(); err != nil {
			return err
		}
		//2.改变 video 表中的 favorite count
		_, err := tx.Video.WithContext(ctx).Where(tx.Video.ID.Eq(vid)).UpdateSimple(tx.Video.Comment_count.Sub(1))
		if err != nil {
			return err
		}
		return nil
	})
}

// GetFavoriteList get favorite video list
func GetFavoriteList(ctx context.Context, uid uint) ([]*model.Video, error) {
	vidsQuery := Favorite.WithContext(ctx).Select(Favorite.VideoId).Where(Favorite.UserId.Eq(uid))
	videos, err := Video.WithContext(ctx).Where(Favorite.WithContext(ctx).Columns(Video.ID).In(vidsQuery)).Find()
	if err != nil {
		return nil, err
	}
	return videos, nil
}
