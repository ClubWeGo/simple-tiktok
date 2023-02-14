package main

import (
	"context"
	comment "github.com/ClubWeGo/commentmicro/kitex_gen/comment"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentMethod implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentMethod(ctx context.Context, request *comment.CommentReq) (resp *comment.CommentResp, err error) {
	// TODO: Your code here...
	return
}

// CommentListMethod implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentListMethod(ctx context.Context, request *comment.CommentListReq) (resp *comment.CommentListResp, err error) {
	// TODO: Your code here...
	return
}
