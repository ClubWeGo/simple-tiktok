package db

import (
	"context"
	"github.com/ClubWeGo/simple-tiktok/services/commentmicro/dal/model"
	"github.com/ClubWeGo/simple-tiktok/services/commentmicro/pkg/errno"
)

func CountVideosComment(ctx context.Context, vid []int64) (map[int64]int64, error) {
	var results []struct {
		VideoId int64
		Count   int64
	}
	err := Comment.WithContext(ctx).Select(Comment.VideoID, Comment.VideoID.Count().As("count")).
		Where(Comment.VideoID.In(vid...)).Group(Comment.VideoID).Scan(&results)
	if err != nil {
		return nil, errno.DBErr.WithMessage(err.Error())
	}
	res := make(map[int64]int64)
	for _, r := range results {
		res[r.VideoId] = r.Count
	}
	return res, nil
}

func AddComment(ctx context.Context, uid int64, vid int64, commenttext *string) error {
	err := Comment.WithContext(ctx).Create(&model.Comment{
		VideoID: vid,
		UserID:  uid,
		Content: *commenttext,
	})
	if err != nil {
		return errno.DBErr.WithMessage(err.Error())
	}
	return nil
}
func DeleteComment(ctx context.Context, commentid int64) error {
	_, err := Comment.WithContext(ctx).Where(Comment.ID.Eq(uint(commentid))).Delete()
	if err != nil {
		return errno.DBErr.WithMessage(err.Error())
	}
	return nil
}
func GetCommentList(ctx context.Context, uid int64, vid int64) ([]*model.Comment, error) {
	comments, err := Comment.WithContext(ctx).Where(Comment.UserID.Eq(uid), Comment.VideoID.Eq(vid)).Find()
	if err != nil {
		return nil, errno.DBErr.WithMessage(err.Error())
	}
	return comments, nil
}
