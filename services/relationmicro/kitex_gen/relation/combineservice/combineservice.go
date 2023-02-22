// Code generated by Kitex v0.4.4. DO NOT EDIT.

package combineservice

import (
	"context"
	relation "github.com/ClubWeGo/simple-tiktok/services/relationmicro/kitex_gen/relation"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

type CombineService interface {
	relation.RelationService
	relation.MessageService
}

func serviceInfo() *kitex.ServiceInfo {
	return combineServiceServiceInfo
}

var combineServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "CombineService"
	handlerType := (*CombineService)(nil)
	methods := map[string]kitex.MethodInfo{
		"FollowMethod":          kitex.NewMethodInfo(followMethodHandler, newRelationServiceFollowMethodArgs, newRelationServiceFollowMethodResult, false),
		"GetFollowInfoMethod":   kitex.NewMethodInfo(getFollowInfoMethodHandler, newRelationServiceGetFollowInfoMethodArgs, newRelationServiceGetFollowInfoMethodResult, false),
		"GetFollowListMethod":   kitex.NewMethodInfo(getFollowListMethodHandler, newRelationServiceGetFollowListMethodArgs, newRelationServiceGetFollowListMethodResult, false),
		"GetFollowerListMethod": kitex.NewMethodInfo(getFollowerListMethodHandler, newRelationServiceGetFollowerListMethodArgs, newRelationServiceGetFollowerListMethodResult, false),
		"GetFriendListMethod":   kitex.NewMethodInfo(getFriendListMethodHandler, newRelationServiceGetFriendListMethodArgs, newRelationServiceGetFriendListMethodResult, false),
		"GetIsFollowsMethod":    kitex.NewMethodInfo(getIsFollowsMethodHandler, newRelationServiceGetIsFollowsMethodArgs, newRelationServiceGetIsFollowsMethodResult, false),
		"GetAllMessageMethod":   kitex.NewMethodInfo(getAllMessageMethodHandler, newMessageServiceGetAllMessageMethodArgs, newMessageServiceGetAllMessageMethodResult, false),
		"SendMessageMethod":     kitex.NewMethodInfo(sendMessageMethodHandler, newMessageServiceSendMessageMethodArgs, newMessageServiceSendMessageMethodResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "relation",
	}
	extra["combine_service"] = true
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func followMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceFollowMethodArgs)
	realResult := result.(*relation.RelationServiceFollowMethodResult)
	success, err := handler.(relation.RelationService).FollowMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceFollowMethodArgs() interface{} {
	return relation.NewRelationServiceFollowMethodArgs()
}

func newRelationServiceFollowMethodResult() interface{} {
	return relation.NewRelationServiceFollowMethodResult()
}

func getFollowInfoMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceGetFollowInfoMethodArgs)
	realResult := result.(*relation.RelationServiceGetFollowInfoMethodResult)
	success, err := handler.(relation.RelationService).GetFollowInfoMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceGetFollowInfoMethodArgs() interface{} {
	return relation.NewRelationServiceGetFollowInfoMethodArgs()
}

func newRelationServiceGetFollowInfoMethodResult() interface{} {
	return relation.NewRelationServiceGetFollowInfoMethodResult()
}

func getFollowListMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceGetFollowListMethodArgs)
	realResult := result.(*relation.RelationServiceGetFollowListMethodResult)
	success, err := handler.(relation.RelationService).GetFollowListMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceGetFollowListMethodArgs() interface{} {
	return relation.NewRelationServiceGetFollowListMethodArgs()
}

func newRelationServiceGetFollowListMethodResult() interface{} {
	return relation.NewRelationServiceGetFollowListMethodResult()
}

func getFollowerListMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceGetFollowerListMethodArgs)
	realResult := result.(*relation.RelationServiceGetFollowerListMethodResult)
	success, err := handler.(relation.RelationService).GetFollowerListMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceGetFollowerListMethodArgs() interface{} {
	return relation.NewRelationServiceGetFollowerListMethodArgs()
}

func newRelationServiceGetFollowerListMethodResult() interface{} {
	return relation.NewRelationServiceGetFollowerListMethodResult()
}

func getFriendListMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceGetFriendListMethodArgs)
	realResult := result.(*relation.RelationServiceGetFriendListMethodResult)
	success, err := handler.(relation.RelationService).GetFriendListMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceGetFriendListMethodArgs() interface{} {
	return relation.NewRelationServiceGetFriendListMethodArgs()
}

func newRelationServiceGetFriendListMethodResult() interface{} {
	return relation.NewRelationServiceGetFriendListMethodResult()
}

func getIsFollowsMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceGetIsFollowsMethodArgs)
	realResult := result.(*relation.RelationServiceGetIsFollowsMethodResult)
	success, err := handler.(relation.RelationService).GetIsFollowsMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceGetIsFollowsMethodArgs() interface{} {
	return relation.NewRelationServiceGetIsFollowsMethodArgs()
}

func newRelationServiceGetIsFollowsMethodResult() interface{} {
	return relation.NewRelationServiceGetIsFollowsMethodResult()
}

func getAllMessageMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.MessageServiceGetAllMessageMethodArgs)
	realResult := result.(*relation.MessageServiceGetAllMessageMethodResult)
	success, err := handler.(relation.MessageService).GetAllMessageMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMessageServiceGetAllMessageMethodArgs() interface{} {
	return relation.NewMessageServiceGetAllMessageMethodArgs()
}

func newMessageServiceGetAllMessageMethodResult() interface{} {
	return relation.NewMessageServiceGetAllMessageMethodResult()
}

func sendMessageMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.MessageServiceSendMessageMethodArgs)
	realResult := result.(*relation.MessageServiceSendMessageMethodResult)
	success, err := handler.(relation.MessageService).SendMessageMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMessageServiceSendMessageMethodArgs() interface{} {
	return relation.NewMessageServiceSendMessageMethodArgs()
}

func newMessageServiceSendMessageMethodResult() interface{} {
	return relation.NewMessageServiceSendMessageMethodResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) FollowMethod(ctx context.Context, request *relation.FollowReq) (r *relation.FollowResp, err error) {
	var _args relation.RelationServiceFollowMethodArgs
	_args.Request = request
	var _result relation.RelationServiceFollowMethodResult
	if err = p.c.Call(ctx, "FollowMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFollowInfoMethod(ctx context.Context, request *relation.GetFollowInfoReq) (r *relation.GetFollowInfoResp, err error) {
	var _args relation.RelationServiceGetFollowInfoMethodArgs
	_args.Request = request
	var _result relation.RelationServiceGetFollowInfoMethodResult
	if err = p.c.Call(ctx, "GetFollowInfoMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFollowListMethod(ctx context.Context, request *relation.GetFollowListReq) (r *relation.GetFollowListResp, err error) {
	var _args relation.RelationServiceGetFollowListMethodArgs
	_args.Request = request
	var _result relation.RelationServiceGetFollowListMethodResult
	if err = p.c.Call(ctx, "GetFollowListMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFollowerListMethod(ctx context.Context, request *relation.GetFollowerListReq) (r *relation.GetFollowerListResp, err error) {
	var _args relation.RelationServiceGetFollowerListMethodArgs
	_args.Request = request
	var _result relation.RelationServiceGetFollowerListMethodResult
	if err = p.c.Call(ctx, "GetFollowerListMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFriendListMethod(ctx context.Context, request *relation.GetFriendListReq) (r *relation.GetFriendListResp, err error) {
	var _args relation.RelationServiceGetFriendListMethodArgs
	_args.Request = request
	var _result relation.RelationServiceGetFriendListMethodResult
	if err = p.c.Call(ctx, "GetFriendListMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetIsFollowsMethod(ctx context.Context, request *relation.GetIsFollowsReq) (r *relation.GetIsFollowsResp, err error) {
	var _args relation.RelationServiceGetIsFollowsMethodArgs
	_args.Request = request
	var _result relation.RelationServiceGetIsFollowsMethodResult
	if err = p.c.Call(ctx, "GetIsFollowsMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetAllMessageMethod(ctx context.Context, request *relation.GetAllMessageReq) (r *relation.GetAllMessageResp, err error) {
	var _args relation.MessageServiceGetAllMessageMethodArgs
	_args.Request = request
	var _result relation.MessageServiceGetAllMessageMethodResult
	if err = p.c.Call(ctx, "GetAllMessageMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SendMessageMethod(ctx context.Context, request *relation.SendMessageReq) (r *relation.SendMessageResp, err error) {
	var _args relation.MessageServiceSendMessageMethodArgs
	_args.Request = request
	var _result relation.MessageServiceSendMessageMethodResult
	if err = p.c.Call(ctx, "SendMessageMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
