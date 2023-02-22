// Code generated by Kitex v0.4.4. DO NOT EDIT.

package userservice

import (
	"context"
	usermicro "github.com/ClubWeGo/simple-tiktok/services/usermicro/kitex_gen/usermicro"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetUserMethod(ctx context.Context, request *usermicro.GetUserReq, callOptions ...callopt.Option) (r *usermicro.GetUserResp, err error)
	GetUserSetByIdSetMethod(ctx context.Context, request *usermicro.GetUserSetByIdSetReq, callOptions ...callopt.Option) (r *usermicro.GetUserSetByIdSetResp, err error)
	LoginUserMethod(ctx context.Context, request *usermicro.LoginUserReq, callOptions ...callopt.Option) (r *usermicro.LoginUserResp, err error)
	CreateUserMethod(ctx context.Context, request *usermicro.CreateUserReq, callOptions ...callopt.Option) (r *usermicro.CreateUserResp, err error)
	UpdateUserMethod(ctx context.Context, request *usermicro.UpdateUserReq, callOptions ...callopt.Option) (r *usermicro.UpdateUserResp, err error)
	UpdateRelationMethod(ctx context.Context, request *usermicro.UpdateRelationCacheReq, callOptions ...callopt.Option) (r *usermicro.UpdateRelationCacheResp, err error)
	UpdateInteractionMethod(ctx context.Context, request *usermicro.UpdateInteractionCacheReq, callOptions ...callopt.Option) (r *usermicro.UpdateInteractionCacheResp, err error)
	UpdateWorkMethod(ctx context.Context, request *usermicro.UpdateWorkCacheReq, callOptions ...callopt.Option) (r *usermicro.UpdateWorkCacheResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) GetUserMethod(ctx context.Context, request *usermicro.GetUserReq, callOptions ...callopt.Option) (r *usermicro.GetUserResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserMethod(ctx, request)
}

func (p *kUserServiceClient) GetUserSetByIdSetMethod(ctx context.Context, request *usermicro.GetUserSetByIdSetReq, callOptions ...callopt.Option) (r *usermicro.GetUserSetByIdSetResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserSetByIdSetMethod(ctx, request)
}

func (p *kUserServiceClient) LoginUserMethod(ctx context.Context, request *usermicro.LoginUserReq, callOptions ...callopt.Option) (r *usermicro.LoginUserResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.LoginUserMethod(ctx, request)
}

func (p *kUserServiceClient) CreateUserMethod(ctx context.Context, request *usermicro.CreateUserReq, callOptions ...callopt.Option) (r *usermicro.CreateUserResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateUserMethod(ctx, request)
}

func (p *kUserServiceClient) UpdateUserMethod(ctx context.Context, request *usermicro.UpdateUserReq, callOptions ...callopt.Option) (r *usermicro.UpdateUserResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateUserMethod(ctx, request)
}

func (p *kUserServiceClient) UpdateRelationMethod(ctx context.Context, request *usermicro.UpdateRelationCacheReq, callOptions ...callopt.Option) (r *usermicro.UpdateRelationCacheResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateRelationMethod(ctx, request)
}

func (p *kUserServiceClient) UpdateInteractionMethod(ctx context.Context, request *usermicro.UpdateInteractionCacheReq, callOptions ...callopt.Option) (r *usermicro.UpdateInteractionCacheResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateInteractionMethod(ctx, request)
}

func (p *kUserServiceClient) UpdateWorkMethod(ctx context.Context, request *usermicro.UpdateWorkCacheReq, callOptions ...callopt.Option) (r *usermicro.UpdateWorkCacheResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateWorkMethod(ctx, request)
}
