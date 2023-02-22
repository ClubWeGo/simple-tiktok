// Code generated by hertz generator.

package interaction

import (
	"context"

	interaction "github.com/ClubWeGo/simple-tiktok/biz/model/interaction"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// CommentListMethod .
// @router /douyin/comment/list/ [GET]
func CommentListMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interaction.CommentListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(interaction.CommentListResp)

	c.JSON(consts.StatusOK, resp)
}
