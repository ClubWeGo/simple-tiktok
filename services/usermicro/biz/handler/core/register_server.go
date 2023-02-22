// Code generated by hertz generator.

package core

import (
	"context"
	"log"

	core "github.com/ClubWeGo/douyin/biz/model/core"
	"github.com/ClubWeGo/douyin/kitex_server"
	"github.com/ClubWeGo/douyin/minio_server"
	"github.com/ClubWeGo/douyin/tools"
	"github.com/ClubWeGo/douyin/tools/safe"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// RegisterMethod .
// @router /douyin/user/register/ [POST]
func RegisterMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.RegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(core.RegisterResp)

	// TODO : 缓存记录ip地址和注册api调用次数，限制统一设备短时间太多的注册，预防黑灰产。同feed中要求

	// 注册字段过滤 : 注入，敏感词检查
	// 虽然gorm提供了sql防注入，但是这里在业务层再加一次字段过滤可以尽早阻止非法请求
	err = safe.SqlInjectCheck(req.Username)
	if err != nil {
		msg := "非法字段"
		resp.StatusCode = 1
		resp.StatusMsg = &msg
		c.JSON(consts.StatusOK, resp)
		return
	}

	// 题目要求的基础注册功能
	// userid, err := kitex_server.RegisterUser(req.Username, *req.Password)
	// 附带个人设置的注册功能
	log.Println(111)
	var testBackgroundImage = "http://" + minio_server.GlobalConfig.Endpoint + "/douyin/" + "backgroud.jpg"
	var testAvatar = "http://" + minio_server.GlobalConfig.Endpoint + "/douyin/" + "0019534761_20.jpg"
	userid, err := kitex_server.RegisterUserALL(req.Username, *req.Password, nil, nil, &testBackgroundImage, &testAvatar)
	log.Println(122)
	if err != nil {
		msgFailed := "注册失败"
		resp.StatusCode = 1
		resp.StatusMsg = &msgFailed
		c.JSON(consts.StatusOK, resp)
		return
	}

	msgsucceed := "注册成功"
	resp.StatusMsg = &msgsucceed
	resp.Token = tools.GenerateToken(userid)
	resp.UserID = userid

	c.JSON(consts.StatusOK, resp)
}