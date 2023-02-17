package db

import (
	"context"
	. "github.com/ClubWeGo/favoritemicro/dal/query"
	"log"
)

//var Q = query.Q

// GetFavoriteRelation get favorite video info
/*func GetFavoriteRelation(ctx context.Context, uid int64, vid int64) ([]*model.Video, error) {
	//user := new(model.User)
	user, err := userQ.WithContext(ctx).Where(userQ.ID.Eq(uint(uid))).First()
	if err != nil {
		return nil, err
	}

	//video := new(Video)
	videos, err := favoriteQ.WithContext(ctx).Where(favoriteQ.UserId.Eq(int(uid))).Find()
	if err != nil {
		return nil, err
	}

	//if err := DB.WithContext(ctx).Model(&user).Association("FavoriteVideos").Find(&video, vid); err != nil {
	//	return nil, err
	//}
	return videos, nil
}*/

func AddFavorite(ctx context.Context, uid uint, vid uint) error {
	log.Println("add favorite")
	user, err := User.WithContext(ctx).Where(User.ID.Eq(uid)).First()
	if err != nil {
		return err
	}
	log.Println(user.ID)
	return nil

	//return Q.Transaction(func(tx *query.Query) error {
	//	// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
	//	//1. 新增点赞数据
	//	//user, err := tx.User.WithContext(ctx).Where(tx.User.ID.Eq(uid)).First()
	//	//if err != nil {
	//	//	return err
	//	//}
	//	//
	//	//video, err := tx.Video.WithContext(ctx).Where(tx.Video.ID.Eq(vid)).First()
	//	//if err != nil {
	//	//	return err
	//	//}
	//
	//	if err := tx.Favorite.WithContext(ctx).Create(&Favorite{UserId: uid, VideoId: vid}); err != nil {
	//		return err
	//	}
	//	//2.改变 video 表中的 favorite count
	//	res, err := tx.Video.WithContext(ctx).Where(tx.Video.ID.Eq(vid)).UpdateSimple(tx.Video.Comment_count.Add(1))
	//	if err != nil {
	//		return err
	//	}
	//
	//	if res.RowsAffected != 1 {
	//		return err
	//	}
	//	return nil
	//})
}
