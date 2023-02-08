package main

import (
	"context"
	usermicro "usermicro/kitex_gen/usermicro"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// GetUserByNameMethod implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserByNameMethod(ctx context.Context, request *usermicro.GetUserReq) (resp *usermicro.GetUserResp, err error) {
	// TODO: Your code here...
	return
}

// LoginUserMethod implements the UserServiceImpl interface.
func (s *UserServiceImpl) LoginUserMethod(ctx context.Context, request *usermicro.LoginUserReq) (resp *usermicro.LoginUserResp, err error) {
	// TODO: Your code here...
	return
}

// CreateUserMethod implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUserMethod(ctx context.Context, request *usermicro.CreateUserReq) (resp *usermicro.CreateUserResp, err error) {
	// TODO: Your code here...
	return
}

// UpdateUserMethod implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUserMethod(ctx context.Context, request *usermicro.UpdateUserReq) (resp *usermicro.UpdateUserResp, err error) {
	// TODO: Your code here...
	return
}

// FollowUserMethod implements the UserServiceImpl interface.
func (s *UserServiceImpl) FollowUserMethod(ctx context.Context, request *usermicro.FollowUserReq) (resp *usermicro.FollowUserResp, err error) {
	// TODO: Your code here...
	return
}
