// Code generated by hertz generator.

package core

import (
	"context"

	core "github.com/ClubWeGo/simple-tiktok/biz/model/core"
	"github.com/ClubWeGo/simple-tiktok/kitex_server"
	"github.com/ClubWeGo/simple-tiktok/tools"
	"github.com/ClubWeGo/simple-tiktok/tools/safe"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// LoginMethod .
// @router /douyin/user/login/ [POST]
func LoginMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.LoginReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(core.LoginResp)

	// 虽然gorm提供了sql防注入，但是这里在业务层再加一次字段过滤可以尽早阻止非法请求
	err = safe.SqlInjectCheck(req.Username)
	if err != nil {
		msg := "非法字段"
		resp.StatusCode = 1
		resp.StatusMsg = &msg
		c.JSON(consts.StatusOK, resp)
		return
	}

	msgsucceed := "登录成功"
	msgFailed := "登录失败"

	userid, err := kitex_server.LoginUser(req.Username, req.Password)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = &msgFailed
		c.JSON(consts.StatusOK, resp)
		return
	}

	resp.StatusMsg = &msgsucceed
	resp.UserID = userid
	resp.Token = tools.GenerateToken(userid)

	c.JSON(consts.StatusOK, resp)
}
