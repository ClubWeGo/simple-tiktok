package kitex_server

import (
	"context"
	"errors"
	"fmt"
	"github.com/prometheus/common/log"

	"github.com/ClubWeGo/douyin/biz/model/relation"
	relationserver "github.com/ClubWeGo/relationmicro/kitex_gen/relation"
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
func GetFollowList(myUid int64, targetUid int64) ([]*relation.User, error) {
	resp, err := Relationclient.GetFollowListMethod(context.Background(), &relationserver.GetFollowListReq{MyId: &myUid, TargetId: targetUid})
	if err != nil {
		log.Errorf("rpc请求relation服务失败，详情:%s", err)
		return nil, fmt.Errorf("本次请求失败,请稍后重试")
	}
	switch resp.StatusCode {
	case SUCCESS:
		userList := make([]*relation.User, len(resp.UserList))
		for i, user := range resp.GetUserList() {
			userList[i] = &relation.User{
				ID:            user.Id,
				Name:          user.Name,
				FollowCount:   &user.FollowCount,
				FollowerCount: &user.FollowerCount,
				IsFollow:      user.IsFollow,
			}
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
func GetFollowerList(myUid int64, targetUid int64) ([]*relation.User, error) {
	resp, err := Relationclient.GetFollowerListMethod(context.Background(), &relationserver.GetFollowerListReq{MyId: &myUid, TargetId: targetUid})
	if err != nil {
		log.Errorf("rpc请求relation服务失败，详情:%s", err)
		return nil, fmt.Errorf("本次请求失败,请稍后重试")
	}
	switch resp.StatusCode {
	case SUCCESS:
		userList := make([]*relation.User, len(resp.UserList))
		for i, user := range resp.GetUserList() {
			userList[i] = &relation.User{
				ID:            user.Id,
				Name:          user.Name,
				FollowCount:   &user.FollowCount,
				FollowerCount: &user.FollowerCount,
				IsFollow:      user.IsFollow,
			}
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

func GetFriendList() {

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
