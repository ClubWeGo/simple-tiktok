package pack

import (
	"errors"
	"github.com/ClubWeGo/favoritemicro/kitex_gen/favorite"
	"github.com/ClubWeGo/favoritemicro/pkg/errno"
)

func BuildBaseResp(err error) (resp *favorite.BaseResp) {
	e := errno.ErrNo{}
	if err == nil {
		e = errno.Success
	} else if errors.As(err, &e) {
		e = errno.ConvertErr(err)
	} else {
		e = errno.ServiceErr.WithMessage(err.Error())
	}
	resp = &favorite.BaseResp{}
	resp.StatusCode = e.ErrCode
	resp.StatusMsg = &e.ErrMsg
	return
}
