package kitex_server

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"strconv"

	"github.com/ClubWeGo/douyin/biz/model/core"
	"github.com/ClubWeGo/douyin/biz/model/relation"
	relationserver "github.com/ClubWeGo/relationmicro/kitex_gen/relation"
	"github.com/prometheus/common/log"
)

// 响应码
const (
	// 服务器异常
	ERROR = 0
	// 正常响应
	SUCCESS = 1
	// 参数校验不通过
	VERIFY = 2
)

// 关注操作类型
const (
	// 关注
	FOLLOW = 1
	// 取关
	UNFOLLOW = 2
)

func SendMsg(fromUserId, toUserId int64, msg string) error {
	resp, err := Relationclient.SendMessageMethod(context.Background(), &relationserver.SendMessageReq{
		UserId:   fromUserId,
		ToUserId: toUserId,
		Content:  msg,
	})
	if err != nil {
		return err
	}
	if resp.Status {
		return nil
	}
	return errors.New("kitex-relationserver : error to send msg")
}

func GetAllMsg(fromUserId, toUserId int64) ([]*relation.Message, error) {
	resp, err := Relationclient.GetAllMessageMethod(context.Background(), &relationserver.GetAllMessageReq{
		UserId:   fromUserId,
		ToUserId: toUserId,
	})
	// 如果出错，拿到的是空切片，所以videoList也是空的
	messageList := make([]*relation.Message, len(resp.Msg)) // Msg这个命名不太好
	if err != nil {
		return messageList, err
	}
	if resp.Status { // true 表示成功
		for index, message := range resp.Msg {
			messageList[index] = &relation.Message{
				ID:         message.Id,
				ToUserID:   message.ToUserId,
				FromUserID: &message.FromUserId,
				Content:    &message.Content,
				CreateTime: *message.CreateTime, // 返回的是 "2006-01-02" 格式
			}
		}
		return messageList, nil
	}
	return messageList, err
}

// 关注
func Follow(myUid int64, targetUid int64, actionType int32) error {
	if errMsg := VerifyFollowParam(myUid, targetUid, actionType); errMsg != nil {
		return fmt.Errorf(*errMsg)
	}
	resp, err := Relationclient.FollowMethod(context.Background(), &relationserver.FollowReq{MyUid: myUid, TargetUid: targetUid, ActionType: actionType})
	if err != nil {
		log.Errorf("rpc请求relation服务失败，详情:%s", err)
		return fmt.Errorf("本次请求失败,请稍后重试")
	}
	// 根据relation返回状态码做日志记录和其他操作
	switch resp.StatusCode {
	case SUCCESS:
		return nil
	case ERROR:
		log.Errorf("relation服务异常，详情:%s", *resp.Msg)
		break
	case VERIFY:
		log.Errorf("relation服务参数校验异常，详情：%s", *resp.Msg)
		break
	default:
		break
	}
	return fmt.Errorf("本次请求失败，请稍后重试")
}

// 获取关注列表
func GetFollowList(myUid int64, targetUid int64) ([]*core.User, error) {
	resp, err := Relationclient.GetFollowListMethod(context.Background(), &relationserver.GetFollowListReq{MyId: &myUid, TargetId: targetUid})
	if err != nil {
		log.Errorf("GetFollowList rpc请求relation服务失败，详情:%s", err)
		return nil, fmt.Errorf("本次请求失败,请稍后重试")
	}
	switch resp.StatusCode {
	case SUCCESS:
		userList := make([]*core.User, len(resp.UserList))
		for i, user := range resp.GetUserList() {
			userList[i] = ConvertCoreUser(user)
		}
		return userList, nil
	case ERROR:
		log.Errorf("relation服务异常，详情:%s", *resp.Msg)
		break
	case VERIFY:
		log.Errorf("relation服务参数校验异常，详情：%s", *resp.Msg)
		break
	default:
		break
	}
	return nil, fmt.Errorf("本次请求失败，请稍后重试")
}

// 粉丝列表
func GetFollowerList(myUid int64, targetUid int64) ([]*core.User, error) {
	resp, err := Relationclient.GetFollowerListMethod(context.Background(), &relationserver.GetFollowerListReq{MyId: &myUid, TargetId: targetUid})
	if err != nil {
		log.Errorf("GetFollowerList rpc请求relation服务失败，详情:%s", err)
		return nil, fmt.Errorf("本次请求失败,请稍后重试")
	}
	switch resp.StatusCode {
	case SUCCESS:
		userList := make([]*core.User, len(resp.UserList))
		for i, user := range resp.GetUserList() {
			userList[i] = ConvertCoreUser(user)
		}
		return userList, nil
	case ERROR:
		log.Errorf("relation服务异常，详情:%s", *resp.Msg)
		break
	case VERIFY:
		log.Errorf("relation服务参数校验异常，详情：%s", *resp.Msg)
		break
	default:
		break
	}
	return nil, fmt.Errorf("本次请求失败，请稍后重试")
}

// 好友列表
func GetFriendList(myUid int64, targetUid int64) ([]*core.User, error) {
	resp, err := Relationclient.GetFriendListMethod(context.Background(), &relationserver.GetFriendListReq{MyUid: &myUid, TargetUid: targetUid})
	if err != nil {
		log.Errorf("GetFriendList rpc请求relation服务失败，详情:%s", err)
		return nil, fmt.Errorf("本次请求失败,请稍后重试")
	}
	switch resp.StatusCode {
	case SUCCESS:
		userList := make([]*core.User, len(resp.FriendList))
		for i, user := range resp.GetFriendList() {
			userList[i] = ConvertCoreUser(user)
		}
		return userList, nil
	case ERROR:
		log.Errorf("relation服务异常，详情:%s", *resp.Msg)
		break
	case VERIFY:
		log.Errorf("relation服务参数校验异常，详情：%s", *resp.Msg)
		break
	default:
		break
	}
	return nil, fmt.Errorf("本次请求失败，请稍后重试")
}

// 校验关注参数
func VerifyFollowParam(myUid int64, targetUid int64, actionType int32) *string {
	var errMsg *string

	if myUid == targetUid {
		*errMsg = "您不能关注自己"
		return errMsg
	}
	if !(actionType == FOLLOW || actionType == UNFOLLOW) {
		*errMsg = "检查到您的数据异常，请求终止"
		return errMsg
	}
	return nil
}

// TODO : .GetIsFollowSetByUserIdSet
// 根据userIds 批量获取用户关注状态
func GetIsFollowSetByUserIdSet(myUid int64, idSet []int64) (map[int64]bool, error) {
	resp, err := Relationclient.GetIsFollowsMethod(context.Background(), &relationserver.GetIsFollowsReq{MyUid: myUid, UserIds: idSet})
	if err != nil {
		log.Errorf("GetIsFollowSetByUserIdSet rpc请求relation服务失败，详情:%s", err)
		return nil, fmt.Errorf("本次请求失败,请稍后重试")
	}
	switch resp.StatusCode {
	case SUCCESS:
		return resp.IsFollowMap, nil
	case ERROR:
		log.Errorf("relation服务异常，详情:%s", *resp.Msg)
		break
	case VERIFY:
		log.Errorf("relation服务参数校验异常，详情：%s", *resp.Msg)
		break
	default:
		break
	}
	return nil, fmt.Errorf("本次请求失败，请稍后重试")
}

// TODO : .GetIsFollowMapByUserIdSet
func GetIsFollowMapByUserIdSet(uid int64, idSet []int64) (isFollowMap map[int64]bool, err error) {
	resp, err := Relationclient.GetIsFollowsMethod(context.Background(), &relationserver.GetIsFollowsReq{MyUid: uid, UserIds: idSet})
	if err != nil {
		log.Errorf("GetIsFollowSetByUserIdSet rpc请求relation服务失败，详情:%s", err)
		return nil, fmt.Errorf("本次请求失败,请稍后重试")
	}
	switch resp.StatusCode {
	case SUCCESS:
		return resp.IsFollowMap, nil
	case ERROR:
		log.Errorf("relation服务异常，详情:%s", *resp.Msg)
		break
	case VERIFY:
		log.Errorf("relation服务参数校验异常，详情：%s", *resp.Msg)
		break
	default:
		break
	}
	return nil, fmt.Errorf("本次请求失败，请稍后重试")
}

// 协程接口

type FollowInfoWithId struct {
	Id         int64
	followInfo relationserver.FollowInfo
}

// 通过GetFollowInfoMethod批量获取FollowInfo，包装为Map
func GetRelationMap(idSet []int64, currentUser int64, respRelationMap chan map[int64]FollowInfoWithId, wg *sync.WaitGroup, errChan chan error) {
	defer wg.Done()

	wgRelation := &sync.WaitGroup{}
	insideResultChan := make(chan FollowInfoWithId, len(idSet))
	insideErrChan := make(chan error, len(idSet))
	for _, id := range idSet {
		wgRelation.Add(1)
		go func(userid int64) {
			defer wgRelation.Done()
			r, err := Relationclient.GetFollowInfoMethod(context.Background(), &relationserver.GetFollowInfoReq{
				MyUid:     &currentUser,
				TargetUid: userid,
			})
			if err != nil {
				insideResultChan <- FollowInfoWithId{} // 出错，传回一个空值
				insideErrChan <- err
				return
			}
			if r.StatusCode == 1 { // 0 error, 1 success, 2 参数不通过
				followInfoWithId := FollowInfoWithId{
					Id:         userid,
					followInfo: *r.FollowInfo,
				}
				insideResultChan <- followInfoWithId
				insideErrChan <- nil
				return // 成功
			}
			insideResultChan <- FollowInfoWithId{} //没有显式出错，但是没有值
			insideErrChan <- nil
		}(id)
	}
	wgRelation.Wait()
	for range idSet { // 检查协程是否出错，出错直接按请求失败处理
		err := <-insideErrChan
		if err != nil {
			respRelationMap <- map[int64]FollowInfoWithId{}
			errChan <- err
			return // 有协程查数据错误，直接报错返回
		}
	}
	// 封装数据
	result := make(map[int64]FollowInfoWithId, len(idSet))
	for range idSet {
		data := <-insideResultChan
		result[data.Id] = data
	}

	respRelationMap <- result // 返回查询结构
	errChan <- nil
}

// kitex relationserver 数据传输 user -> kitex 回显 core.User
func ConvertCoreUser(user *relationserver.User) *core.User {
	totalFavourited := strconv.FormatInt(*user.TotalFavorited, 10)
	return &core.User{
		ID:              user.Id,
		Name:            user.Name,
		FollowCount:     *user.FollowCount,
		FollowerCount:   *user.FollowerCount,
		IsFollow:        user.IsFollow,
		Avatar:          *user.Avatar,
		BackgroundImage: *user.BackgroundImage,
		Signature:       *user.Signature,
		TotalFavourited: totalFavourited,
		WorkCount:       *user.WorkCount,
		FavoriteCount:   *user.FavoriteCount,
	}
}
