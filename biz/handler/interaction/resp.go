package interaction

import (
	"github.com/ClubWeGo/douyin/tools/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type Response struct {
	StatusCode    int32  `json:"status_code"`
	StatusMessage string `json:"status_message"`
}

// SendResponse pack response
func SendResponse(c *app.RequestContext, err error, data interface{}) {
	if err != nil {
		Err := errno.ConvertErr(err)
		c.JSON(consts.StatusOK, Response{
			StatusCode:    Err.ErrCode,
			StatusMessage: Err.ErrMsg,
		})
		return
	}
	c.JSON(consts.StatusOK, data)
}
