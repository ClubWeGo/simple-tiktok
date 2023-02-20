package main

import (
	"context"
	comment "github.com/ClubWeGo/commentmicro/kitex_gen/comment"
	"github.com/ClubWeGo/commentmicro/service"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentMethod implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentMethod(ctx context.Context, request *comment.CommentReq) (resp *comment.CommentResp, err error) {
	err = service.NewCommentActionService(ctx).CommentAction(request)
	if err != nil {
		statumsg := "error"
		return &comment.CommentResp{
			StatusCode: 201,
			StatusMsg:  &statumsg,
		}, err
	}
	statumsg := "success"
	return &comment.CommentResp{
		StatusCode: 200,
		StatusMsg:  &statumsg,
	}, nil
}

// CommentListMethod implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentListMethod(ctx context.Context, request *comment.CommentListReq) (resp *comment.CommentListResp, err error) {

	comments, err := service.NewCommentListService(ctx).CommentList(request)
	if err != nil {
		statumsg := "error"
		return &comment.CommentListResp{
			StatusCode: 201,
			StatusMsg:  &statumsg,
		}, err
	}
	//fmt.Println(comments)
	statumsg := "success"
	return &comment.CommentListResp{
		StatusCode:  200,
		StatusMsg:   &statumsg,
		CommentList: comments,
	}, nil
}
