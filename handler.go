package main

import (
	"context"
	"time"

	"github.com/ClubWeGo/videomicro/dal/model"
	"github.com/ClubWeGo/videomicro/dal/query"
	videomicro "github.com/ClubWeGo/videomicro/kitex_gen/videomicro"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// CreateVideoMethod implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CreateVideoMethod(ctx context.Context, request *videomicro.CreateVideoReq) (resp *videomicro.CreateVideoResp, err error) {
	// TODO: Your code here...
	v := query.Video

	video := &model.Video{
		Title:          request.Title,
		Author_id:      request.AuthorId,
		Cover_url:      request.CoverUrl,
		Play_url:       request.PlayUrl,
		Favorite_count: 0,
		Comment_count:  0,
	}
	err = v.Create(video)
	if err != nil {
		return &videomicro.CreateVideoResp{
			Status: false,
		}, err
	}
	return &videomicro.CreateVideoResp{
		Status: true,
	}, nil
}

// GetVideoMethod implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideoMethod(ctx context.Context, request *videomicro.GetVideoReq) (resp *videomicro.GetVideoResp, err error) {
	// TODO: Your code here...

	v := query.Video

	video, err := v.Where(v.ID.Eq(uint(request.Id))).First()
	if err != nil {
		return &videomicro.GetVideoResp{
			Status: false,
		}, err
	}
	return &videomicro.GetVideoResp{
		Status: true,
		Video: &videomicro.Video{
			Id:            int64(video.ID),
			Title:         video.Title,
			AuthorId:      video.Author_id,
			PlayUrl:       video.Play_url,
			CoverUrl:      video.Cover_url,
			FavoriteCount: video.Favorite_count,
			CommentCount:  video.Comment_count,
		},
	}, nil
}

// GetVideosByAuthorIdMethod implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideosByAuthorIdMethod(ctx context.Context, request *videomicro.GetVideosByAuthorIdReq) (resp *videomicro.GetVideosByAuthorIdResp, err error) {
	// TODO: Your code here...

	v := query.Video
	// return all video by this user, desc
	videolist, err := v.Where(v.Author_id.Eq(request.AuthorId)).Order(v.CreatedAt.Desc()).Find()
	if err != nil {
		return &videomicro.GetVideosByAuthorIdResp{
			Status: true,
		}, err
	}

	respvideolist := make([]*videomicro.Video, 0) //此处需要改进，很低效

	// 如果没有视频，直接返回空信息
	if len(videolist) == 0 {
		return &videomicro.GetVideosByAuthorIdResp{
			Status:    true,
			VideoList: respvideolist,
		}, nil
	}

	for _, video := range videolist {
		respvideolist = append(respvideolist, &videomicro.Video{
			Id:            int64(video.ID),
			Title:         video.Title,
			AuthorId:      video.Author_id,
			PlayUrl:       video.Play_url,
			CoverUrl:      video.Cover_url,
			FavoriteCount: video.Favorite_count,
			CommentCount:  video.Comment_count,
		})
	}

	return &videomicro.GetVideosByAuthorIdResp{
		Status:    true,
		VideoList: respvideolist,
	}, nil
}

// GetVideosFeedMethod implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideosFeedMethod(ctx context.Context, request *videomicro.GetVideosFeedReq) (resp *videomicro.GetVideosFeedResp, err error) {
	// TODO: Your code here...

	v := query.Video
	// Desc 倒序查询
	startTime := time.Unix(request.LatestTime, 0)
	videolist, err := v.Where(v.CreatedAt.Lt(startTime)).Limit(int(request.Limit)).Order(v.CreatedAt.Desc()).Find()
	if err != nil {
		return &videomicro.GetVideosFeedResp{
			Status: true,
		}, err
	}

	respvideolist := make([]*videomicro.Video, 0) //此处需要改进，很低效

	// 如果没有视频，则直接返回空信息，且时间为传入的判断时间
	if len(videolist) == 0 {
		return &videomicro.GetVideosFeedResp{
			Status:    true,
			NextTime:  &request.LatestTime,
			VideoList: respvideolist,
		}, nil
	}

	for _, video := range videolist {
		respvideolist = append(respvideolist, &videomicro.Video{
			Id:            int64(video.ID),
			Title:         video.Title,
			AuthorId:      video.Author_id,
			PlayUrl:       video.Play_url,
			CoverUrl:      video.Cover_url,
			FavoriteCount: video.Favorite_count,
			CommentCount:  video.Comment_count,
		})
	}

	endTimeUnix := videolist[len(videolist)-1].CreatedAt.Unix()

	return &videomicro.GetVideosFeedResp{
		Status:    true,
		NextTime:  &endTimeUnix,
		VideoList: respvideolist,
	}, nil
}

// DeleteVideoMethod implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) DeleteVideoMethod(ctx context.Context, request *videomicro.DeleteVideoReq) (resp *videomicro.DeleteVideoResp, err error) {
	// TODO: Your code here...

	v := query.Video

	_, err = v.Where(v.ID.Eq(uint(request.Id))).Delete(&model.Video{}) // 软删除

	if err != nil {
		return &videomicro.DeleteVideoResp{
			Status: false,
		}, err
	}
	return &videomicro.DeleteVideoResp{
		Status: true,
	}, err
}

// UpdateVideoMethod implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) UpdateVideoMethod(ctx context.Context, request *videomicro.UpdateVideoReq) (resp *videomicro.UpdateVideoResp, err error) {
	// TODO: Your code here...

	v := query.Video

	video := v.Where(v.ID.Eq(uint(request.Id)))

	if request.Title != nil {
		_, err := video.Update(v.Title, request.Title)
		if err != nil {
			return &videomicro.UpdateVideoResp{
				Status: false,
			}, err
		}
	}

	return &videomicro.UpdateVideoResp{
		Status: true,
	}, err
}
