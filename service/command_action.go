package service

import (
	"context"
	"github.com/ClubWeGo/commentmicro/dal/model"
	"github.com/ClubWeGo/commentmicro/dal/query"
	"github.com/ClubWeGo/commentmicro/kitex_gen/comment"
)

type CommentActionService struct {
	ctx context.Context
}

// NewCommentActionService new CommentActionService
func NewCommentActionService(ctx context.Context) *CommentActionService {
	return &CommentActionService{ctx: ctx}
}

// CommentActionService action comment.
func (s *CommentActionService) CommentAction(req *comment.CommentReq) error {
	u := query.Comment
	comment1 := &model.Comment{
		UserID:  int(req.UserId),
		VideoID: int(req.VideoId),
		Content: *req.CommentText,
	}
	if req.ActionType == 1 {
		err := u.Create(comment1)
		if err != nil {
			return err
		}
		return nil
	}
	if req.ActionType == 2 {
		id := *req.CommentId
		findedcomment, err := u.Where(u.ID.Eq(uint(id))).First()
		if err != nil {
			return err
		}
		_, err = u.Delete(findedcomment)
		if err != nil {
			return err
		}
		return err
	}
	return nil
}
