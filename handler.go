package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"

	"github.com/ClubWeGo/usermicro/dal/model"
	"github.com/ClubWeGo/usermicro/dal/query"
	usermicro "github.com/ClubWeGo/usermicro/kitex_gen/usermicro"
	service "github.com/ClubWeGo/usermicro/service"
)

func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// GetUserMethod implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserMethod(ctx context.Context, request *usermicro.GetUserReq) (resp *usermicro.GetUserResp, err error) {
	// TODO: Your code here...

	user, err := service.QueryUserByIdOrNameOrEmail(request.Id, request.Name, request.Email)
	if err != nil {
		return &usermicro.GetUserResp{
			Status: false,
		}, err
	}

	userinstance, err := user.First()
	if err != nil {
		return &usermicro.GetUserResp{
			Status: false,
		}, err
	}

	return &usermicro.GetUserResp{
		Status: true,
		User: &usermicro.UserInfo{
			Id:            int64(userinstance.ID),
			Name:          userinstance.Name,
			Email:         &userinstance.Email,
			FollowCount:   userinstance.Follow_count,
			FollowerCount: userinstance.Follower_count,
		},
	}, nil
}

// LoginUserMethod implements the UserServiceImpl interface.
func (s *UserServiceImpl) LoginUserMethod(ctx context.Context, request *usermicro.LoginUserReq) (resp *usermicro.LoginUserResp, err error) {
	// TODO: Your code here...
	user, err := service.QueryUserByIdOrNameOrEmail(nil, request.Name, request.Email)
	if err != nil {
		return &usermicro.LoginUserResp{
			Status: false,
			User:   &usermicro.UserInfo{},
		}, err
	}

	userinstance, err := user.First()
	if err != nil {
		return &usermicro.LoginUserResp{
			Status: false,
			User:   &usermicro.UserInfo{},
		}, err
	}

	if MD5(request.Password) == userinstance.Password {
		return &usermicro.LoginUserResp{ // succeed
			Status: true,
			User: &usermicro.UserInfo{
				Id:            int64(userinstance.ID),
				Name:          userinstance.Name,
				Email:         &userinstance.Email,
				FollowCount:   userinstance.Follow_count,
				FollowerCount: userinstance.Follower_count,
			},
		}, err
	}

	return &usermicro.LoginUserResp{
		Status: false,
		User:   &usermicro.UserInfo{},
	}, err
}

// CreateUserMethod implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUserMethod(ctx context.Context, request *usermicro.CreateUserReq) (resp *usermicro.CreateUserResp, err error) {
	// TODO: Your code here...
	u := query.User

	var email string
	if request.Email != nil {
		email = *request.Email
	}

	user1 := &model.User{
		Name:           request.Name,
		Email:          email,
		Password:       MD5(request.Password),
		Follow_count:   0,
		Follower_count: 0,
	}

	err = u.Create(user1)
	if err != nil {
		return &usermicro.CreateUserResp{
			Status: false,
			User:   &usermicro.UserInfo{},
		}, err
	}
	return &usermicro.CreateUserResp{
		Status: true,
		User: &usermicro.UserInfo{
			Id:            int64(user1.ID),
			Name:          user1.Name,
			Email:         &user1.Email,
			FollowCount:   user1.Follow_count,
			FollowerCount: user1.Follower_count,
		},
	}, nil
}

// UpdateUserMethod implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUserMethod(ctx context.Context, request *usermicro.UpdateUserReq) (resp *usermicro.UpdateUserResp, err error) {
	// TODO: Your code here...
	u := query.User

	user, err := service.QueryUserByIdOrName(request.Id, request.Name)
	if err != nil {
		return &usermicro.UpdateUserResp{
			Status: false,
		}, err
	}

	if request.Email != nil {
		user.Update(u.Email, request.Email)
	}
	if request.Password != nil {
		user.Update(u.Password, MD5(*request.Password))
	}

	userinstance, err := user.First()
	if err != nil {
		return &usermicro.UpdateUserResp{
			Status: false,
		}, err
	}

	return &usermicro.UpdateUserResp{
		Status: true,
		User: &usermicro.UserInfo{
			Id:            int64(userinstance.ID),
			Name:          userinstance.Name,
			Email:         &userinstance.Email,
			FollowCount:   userinstance.Follow_count,
			FollowerCount: userinstance.Follower_count,
		},
	}, nil
}

// GetUserSetByIdSetMethod implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserSetByIdSetMethod(ctx context.Context, request *usermicro.GetUserSetByIdSetReq) (resp *usermicro.GetUserSetByIdSetResp, err error) {
	// TODO: Your code here...

	u := query.User

	idSet := request.IdSet

	// 切片互转有内存风险，暂采用最原始的方式转换id格式
	idSetUint := make([]uint, len(idSet))
	for index, id := range idSet {
		idSetUint[index] = uint(id)
	}

	users, err := u.Where(u.ID.In(idSetUint...)).Find()
	if err != nil {
		return &usermicro.GetUserSetByIdSetResp{
			Status: false,
		}, err
	}

	respvideolist := make([]*usermicro.UserInfo, len(users))
	for index, user := range users {
		respvideolist[index] = &usermicro.UserInfo{
			Id:            int64(user.ID),
			Name:          user.Name,
			Email:         &user.Email,
			FollowCount:   user.Follow_count,
			FollowerCount: user.Follower_count,
		}
	}

	return &usermicro.GetUserSetByIdSetResp{
		Status:  true,
		UserSet: respvideolist,
	}, nil
}
