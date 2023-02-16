package kitex_server

import (
	"context"
	"errors"

	"github.com/ClubWeGo/douyin/biz/model/core"
	"github.com/ClubWeGo/videomicro/kitex_gen/videomicro"
)

func GetVideosByAuthorId(id int64) ([]*core.Video, error) {
	r, err := Videoclient.GetVideosByAuthorIdMethod(context.Background(), &videomicro.GetVideosByAuthorIdReq{
		AuthorId: id,
	})
	// 如果出错，拿到的是空切片，所以videoList也是空的
	videoList := make([]*core.Video, len(r.VideoList))
	if err != nil {
		return videoList, err
	}
	if r.Status {
		author, _ := GetUser(id) // 避免不必要的检索
		for index, video := range r.VideoList {
			// 暂时不做处理，错误返回空对象即可
			videoList[index] = &core.Video{
				ID:            video.Id,
				Author:        author,
				PlayURL:       video.PlayUrl,
				CoverURL:      video.CoverUrl,
				FavoriteCount: video.FavoriteCount,
				CommentCount:  video.CommentCount,
				IsFavorite:    false, // 需要增加喜欢配置
				Title:         video.Title,
			}
		}
		return videoList, nil
	}
	return videoList, errors.New("kitex-videomicroserver : error to get video by authorod") // return a null video list
}

func CreateVideo(title, playUrl, coverUrl string, authorId int64) error {
	r, err := Videoclient.CreateVideoMethod(context.Background(), &videomicro.CreateVideoReq{
		Title:    title,
		AuthorId: authorId,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
	})

	if err != nil {
		return err
	}

	if r.Status {
		return nil
	}
	return errors.New("kitex-videomicroserver : create video failed")
}
