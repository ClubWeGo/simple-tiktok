package pack

import (
	"errors"
	"github.com/ClubWeGo/commentmicro/kitex_gen/comment"
	"github.com/ClubWeGo/commentmicro/pkg/errno"
)

func BuildBaseResp(err error) (resp *comment.BaseResp) {
	e := errno.ErrNo{}
	if err == nil {
		e = errno.Success
	} else if errors.As(err, &e) {
		e = errno.ConvertErr(err)
	} else {
		e = errno.ServiceErr.WithMessage(err.Error())
	}
	resp = &comment.BaseResp{}
	resp.StatusCode = e.ErrCode
	resp.StatusMsg = e.ErrMsg
	return
}
