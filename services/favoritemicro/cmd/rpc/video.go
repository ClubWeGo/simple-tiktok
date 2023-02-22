package rpc

import (
	"context"
	"github.com/ClubWeGo/simple-tiktok/services/favoritemicro/pkg/errno"
	"github.com/ClubWeGo/videomicro/kitex_gen/videomicro"
)

func GetAuthorId(ctx context.Context, vid int64) (uid int64, err error) {
	res, err := VideoClient.GetVideoAuthorIdMethod(ctx, &videomicro.GetVideoAuthorIdReq{Id: vid})
	if err != nil {
		return 0, errno.RPCErr.WithMessage(err.Error())
	}
	if !res.Status {
		return 0, errno.RPCErr.WithMessage("GetAuthorId failed")
	}
	return res.AuthorId, nil
}
