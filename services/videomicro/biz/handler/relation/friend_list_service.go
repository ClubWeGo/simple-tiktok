// Code generated by hertz generator.

package relation

import (
	"context"
	"github.com/ClubWeGo/douyin/kitex_server"
	"github.com/ClubWeGo/douyin/tools"

	relation "github.com/ClubWeGo/douyin/biz/model/relation"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// FriendListMethod .
// @router /douyin/relation/friend/list/ [GET]
func FriendListMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req relation.FriendListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp := new(relation.FollowerListResp)

	// 鉴权
	ifvalid, myUid, err := tools.ValidateToken(req.Token)
	if err != nil {
		msgFailed := "非法token"
		resp.StatusCode = 1
		resp.StatusMsg = &msgFailed
		c.JSON(consts.StatusOK, resp)
		return
	}
	if !ifvalid {
		msgFailed := "token无效"
		resp.StatusCode = 1
		resp.StatusMsg = &msgFailed
		c.JSON(consts.StatusOK, resp)
		return
	}
	// 获取好友列表
	friendList, err := kitex_server.GetFriendList(myUid, req.UserID)
	if err != nil {
		errMsg := err.Error()
		resp.StatusCode = 1
		resp.StatusMsg = &errMsg
		c.JSON(consts.StatusOK, resp)
		return
	}

	resp.UserList = friendList
	c.JSON(consts.StatusOK, resp)
}
