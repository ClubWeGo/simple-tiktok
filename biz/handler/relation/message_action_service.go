// Code generated by hertz generator.

package relation

import (
	"context"

	relation "github.com/ClubWeGo/simple-tiktok/biz/model/relation"
	"github.com/ClubWeGo/simple-tiktok/kitex_server"
	"github.com/ClubWeGo/simple-tiktok/tools"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// MessageActionMethod .
// @router /douyin/message/action/ [POST]
func MessageActionMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req relation.MessageActionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(relation.MessageActionResp)

	ifvalid, userId, err := tools.ValidateToken(req.Token)
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

	if req.ActionType == 1 { // send action
		err = kitex_server.SendMsg(userId, req.ToUserID, req.Content)
		if err != nil {
			msgFailed := err.Error()
			resp.StatusCode = 1
			resp.StatusMsg = &msgFailed
			c.JSON(consts.StatusOK, resp)
			return
		}
	} else {
		msgFailed := "非法动作"
		resp.StatusCode = 1
		resp.StatusMsg = &msgFailed
		c.JSON(consts.StatusOK, resp)
		return
	}

	c.JSON(consts.StatusOK, resp)
}
