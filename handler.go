package main

import (
	"context"
	"github.com/ClubWeGo/commentmicro/dal/db"
	comment "github.com/ClubWeGo/commentmicro/kitex_gen/comment"
	"github.com/ClubWeGo/commentmicro/pack"
	"github.com/ClubWeGo/commentmicro/pkg/errno"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentMethod implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentMethod(ctx context.Context, request *comment.CommentReq) (resp *comment.CommentResp, err error) {
	resp = &comment.CommentResp{}
	if request.ActionType == 1 {
		err = db.AddComment(ctx, request.UserId, request.VideoId, request.CommentText)
		if err != nil {
			resp.StatusCode = errno.DBErrCode
			msg := "add comment failed."
			resp.StatusMsg = msg
			return resp, nil
		}
		resp.StatusCode = errno.SuccessCode
		msg := "add comment success."
		resp.StatusMsg = msg
		return resp, nil
	} else if request.ActionType == 2 {
		err = db.DeleteComment(ctx, *request.CommentId)
		if err != nil {
			resp.StatusCode = errno.DBErrCode
			msg := "delete comment failed."
			resp.StatusMsg = msg
			return resp, nil
		}
		resp.StatusCode = errno.SuccessCode
		msg := "delete comment success."
		resp.StatusMsg = msg
		return resp, nil
	} else {
		resp.StatusCode = errno.ParamErrCode
		msg := "param error."
		resp.StatusMsg = msg
		return resp, nil
	}
}

// CommentListMethod implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentListMethod(ctx context.Context, request *comment.CommentListReq) (resp *comment.CommentListResp, err error) {
	resp = &comment.CommentListResp{}
	commentList, err := db.GetCommentList(ctx, request.UserId, request.VideoId)
	if err != nil {
		resp.StatusCode = errno.DBErrCode
		msg := "get comment list failed."
		resp.StatusMsg = msg
		return resp, nil
	}
	msg := "get comment list success."
	respcommentList, err := pack.Comments(commentList)

	resp.StatusCode = errno.SuccessCode
	resp.StatusMsg = msg
	resp.CommentList = respcommentList
	return resp, nil
}

// VideosFavoriteCountMethod implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) VideosFavoriteCountMethod(ctx context.Context, request *comment.VideosCommentCountReq) (resp *comment.VideosCommentCountResp, err error) {
	resp = &comment.VideosCommentCountResp{}
	commentCountMap, err := db.CountVideosComment(ctx, request.VideoIdList)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.CommentCountMap = commentCountMap
	return
}
