// Code generated by hertz generator.

package interaction

import (
	"context"

	interaction "github.com/ClubWeGo/douyin/biz/model/interaction"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// FavoriteMethod .
// @router /douyin/favorite/action/ [POST]
func FavoriteMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interaction.FavoriteReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(interaction.FavoriteResp)

	c.JSON(consts.StatusOK, resp)
}