package kitex_server

import (
	"context"
	"errors"

	"github.com/ClubWeGo/douyin/biz/model/relation"
	relationserver "github.com/ClubWeGo/relationmicro/kitex_gen/relation"
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
