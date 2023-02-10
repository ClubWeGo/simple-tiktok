package kitex_server

import (
	"context"
	"errors"

	"github.com/ClubWeGo/douyin/biz/model/core"
	"github.com/ClubWeGo/usermicro/kitex_gen/usermicro"
)

func GetUser(userid int64) (*core.User, error) {
	r, err := Userclient.GetUserMethod(context.Background(), &usermicro.GetUserReq{Id: &userid})
	if err != nil {
		return &core.User{}, err
	}
	if r.Status {
		return &core.User{
			ID:            r.User.Id,
			Name:          r.User.Name,
			FollowCount:   r.User.FollowCount,
			FollowerCount: r.User.FollowerCount,
			IsFollow:      false, // 需要后续增加社交接口，才可以实现follow
		}, nil
	}
	return &core.User{}, errors.New("kitex-usermicroserver : error to get user") // return a null user
}

func RegisterUser(username, password string) (userid int64, err error) {
	r, err := Userclient.CreateUserMethod(context.Background(), &usermicro.CreateUserReq{
		Name:     username,
		Password: password, //加密由user微服务进行
	})
	if err != nil {
		return 0, errors.New("kitex-usermicroserver : error to create new user")
	}

	if r.Status {
		return r.User.Id, nil
	}
	return 0, errors.New("kitex-usermicroserver : error to create new user")
}

func LoginUser(username, password string) (userid int64, err error) {
	r, err := Userclient.LoginUserMethod(context.Background(), &usermicro.LoginUserReq{
		Name:     &username,
		Password: password,
	})
	if err != nil {
		return 0, errors.New("kitex-usermicroserver : error to login")
	}
	if r.Status {
		return r.User.Id, nil
	}
	return 0, errors.New("kitex-usermicroserver : login failed")
}
