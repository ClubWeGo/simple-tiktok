package kitex_server

import (
	"context"
	"errors"
	"strconv"
	"sync"

	"github.com/ClubWeGo/douyin/biz/model/core"
	"github.com/ClubWeGo/usermicro/kitex_gen/usermicro"
	"github.com/ClubWeGo/videomicro/kitex_gen/videomicro"
	"github.com/prometheus/common/log"
)

// 工具函数

func ConvertUserInfoSetToMap(setData []*usermicro.UserInfo) map[int64]usermicro.UserInfo {
	dataMap := make(map[int64]usermicro.UserInfo, len(setData))
	for _, data := range setData {
		dataMap[data.Id] = *data
	}
	return dataMap
}

func ConvertUserInfoMapToCoreUserMap(mapData map[int64]usermicro.UserInfo) map[int64]core.User {
	newMap := make(map[int64]core.User, len(mapData))
	for key, value := range mapData {
		newMap[key] = core.User{
			ID:              value.Id,
			Name:            value.Name,
			FollowCount:     value.FollowCount,
			FollowerCount:   value.FollowerCount,
			IsFollow:        false,
			Avatar:          value.Avatar,
			BackgroundImage: value.BackgroundImage,
			Signature:       value.Signature,
			TotalFavourited: strconv.FormatInt(value.TotalFavorited, 10),
			WorkCount:       value.WorkCount,
			FavoriteCount:   value.FavoriteCount,
		}
	}
	return newMap
}

// 协程接口

// 通过GetUserSetByIdSet获取用户集合，然后转为Map, 此处获取的是usermicro中的user，交互和社交字段是不准确的
func GetUserMap(idSet []int64, respUserMap chan map[int64]core.User, wg *sync.WaitGroup, errChan chan error) {
	defer wg.Done()

	r, err := Userclient.GetUserSetByIdSetMethod(context.Background(), &usermicro.GetUserSetByIdSetReq{
		IdSet: idSet,
	})
	if err != nil {
		respUserMap <- map[int64]core.User{} // 出错返回一个空的，否则会阻塞<-chan
		errChan <- err
		return
	}
	if r.Status {
		rUserMap := ConvertUserInfoSetToMap(r.UserSet)
		respUserMap <- ConvertUserInfoMapToCoreUserMap(rUserMap)
		errChan <- nil
		return // 成功
	}
	respUserMap <- map[int64]core.User{}
	errChan <- errors.New("userservice GetUserSetByIDSet error: 微服务调用成功，但是返回状态显示失败")
}

// 用于获取用户的最新数据；写成协程，是为了feed那块可以和视频查询协程并发查询
// 根据idSet查core.User的map
func GetUserLatestMap(idSet []int64, currentUser int64, respUserMap chan map[int64]core.User, wg *sync.WaitGroup, errChan chan []error) {
	defer wg.Done() // 外部的wg

	wgUser := &sync.WaitGroup{} // 本函数子协程的wg

	// 批量查询author: Userclient.GetUserSetByIdSetMethod()，获取不准确的user信息
	// 传切片，切片扩容会导致地址改变，故还是chan通信可靠
	respAuthorMap := make(chan map[int64]core.User, 1)
	defer close(respAuthorMap)
	respAuthorMapError := make(chan error, 1)
	defer close(respAuthorMapError)
	wgUser.Add(1)
	go GetUserMap(idSet, respAuthorMap, wgUser, respAuthorMapError)

	// 批量查询author作品数，Work_count 从video服务
	// Videoclient.GetVideoCountSetByIdUserSetMethod(context.Background(), &videomicro.GetVideoCountSetByIdUserSetReq{})
	respVideoCountMap := make(chan map[int64]videomicro.VideoCount, 1)
	defer close(respVideoCountMap)
	respVideoCountMapError := make(chan error, 1)
	defer close(respVideoCountMapError)
	wgUser.Add(1)
	go GetVideoCountMap(idSet, respVideoCountMap, wgUser, respVideoCountMapError)

	// 批量查询FollowCount, FollowerCount，Is_follow 从relation服务
	// Videoclient.GetVideoCountSetByIdUserSetMethod(context.Background(), &videomicro.GetVideoCountSetByIdUserSetReq{})
	respRelationMap := make(chan map[int64]FollowInfoWithId, 1)
	defer close(respRelationMap)
	respRelationMapError := make(chan error, 1)
	defer close(respRelationMapError)
	wgUser.Add(1)
	go GetRelationMap(idSet, currentUser, respRelationMap, wgUser, respRelationMapError)

	// TODO : TotalFavourited, FavoriteCount，传入查询的userId切片，查对应这两个字段的切片，（结果需要携带UserId）：从favorite服务

	// 等待数据
	wgUser.Wait()

	// 处理协程错误
	AuthorMap := <-respAuthorMap
	var errSlice = []error{} // 防止外部设置的chan缓存不够造成阻塞，干脆要求外部设置长度为1的error切片类型
	err := <-respAuthorMapError
	if err != nil {
		errSlice = append(errSlice, err)
	}

	VideoCountMap := <-respVideoCountMap
	err = <-respVideoCountMapError
	if err != nil {
		errSlice = append(errSlice, err)
	}

	RelationMap := <-respRelationMap
	err = <-respRelationMapError
	if err != nil {
		errSlice = append(errSlice, err)
	}
	// TODO: 其他协程的错误处理

	errChan <- errSlice // 错误切片

	// 更新数据
	for id, user := range AuthorMap {
		AuthorMap[id] = core.User{
			ID:              user.ID,
			Name:            user.Name,
			FollowCount:     RelationMap[id].followInfo.FollowCount,   // Relation服务 最新的followCount
			FollowerCount:   RelationMap[id].followInfo.FollowerCount, // Relation服务 最新的followerCount
			IsFollow:        RelationMap[id].followInfo.IsFollow,      // Relation服务 最新的isFollow
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			TotalFavourited: "",                      // TODO: 从获取的数据中拿
			WorkCount:       VideoCountMap[id].Count, // 最新的count数据
			FavoriteCount:   0,                       // TODO: 从获取的数据中拿
		}

	}
	respUserMap <- AuthorMap // 返回数据
}

// 业务接口

// 基础注册：用户名和密码
func RegisterUser(username, password string) (userid int64, err error) {
	newUser := usermicro.CreateUserInfo{
		Name: username,
	}
	r, err := Userclient.CreateUserMethod(context.Background(), &usermicro.CreateUserReq{
		Newuser_: &newUser,
		Password: password, // 此处传输明文，加密由user微服务进行
	})
	if err != nil {
		log.Errorf("注册error：%s", err)
		return 0, err
	}

	if r.Status {
		return *r.UserId, nil
	}
	return 0, errors.New("kitex-usermicroserver : error to create new user")
}

// 完善用户信息的注册
func RegisterUserALL(username, password string, email, signature, backgroundImage, avatar *string) (userid int64, err error) {
	newUser := usermicro.CreateUserInfo{
		Name:            username,
		Email:           email,
		Signature:       signature,
		BackgroundImage: backgroundImage,
		Avatar:          avatar,
	}
	r, err := Userclient.CreateUserMethod(context.Background(), &usermicro.CreateUserReq{
		Newuser_: &newUser,
		Password: password, // 此处传输明文，加密由user微服务进行
	})
	if err != nil {
		return 0, err
	}

	if r.Status {
		return *r.UserId, nil
	}
	return 0, errors.New("kitex-usermicroserver : error to create new user")
}

// 登录校验，目前只支持用户名密码
func LoginUser(username, password string) (userid int64, err error) {
	r, err := Userclient.LoginUserMethod(context.Background(), &usermicro.LoginUserReq{
		Name:     username,
		Password: password,
	})
	if err != nil {
		return 0, err
	}
	if r.Status {
		return *r.UserId, nil
	}
	return 0, errors.New("kitex-usermicroserver : login failed")
}
