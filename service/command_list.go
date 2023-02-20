package service

import (
	"context"
	"fmt"
	"github.com/ClubWeGo/commentmicro/dal/query"
	"github.com/ClubWeGo/commentmicro/kitex_gen/comment"
)

type CommentListService struct {
	ctx context.Context
}

// NewCommentActionService new CommentActionService
func NewCommentListService(ctx context.Context) *CommentListService {
	return &CommentListService{
		ctx: ctx,
	}
}

// CommentList return comment list
func (s *CommentListService) CommentList(req *comment.CommentListReq) ([]*comment.Comment, error) {

	u := query.Comment
	id := req.VideoId
	findedcomments, err := u.Where(u.VideoID.Eq(int(id))).Find()
	if err != nil {
		return nil, err
	}
	fmt.Println(findedcomments)
	comments := make([]*comment.Comment, 0)

	for _, v := range findedcomments {
		comments = append(comments, &comment.Comment{
			Id: int64(v.ID),
			User: &comment.User{
				Id: req.UserId,
			},
			Content:    v.Content,
			CreateDate: v.CreatedAt.Format("01-02"),
		})
	}
	return comments, nil
}
