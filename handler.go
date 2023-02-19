package main

import (
	"context"
	"time"

	"github.com/ClubWeGo/videomicro/dal/model"
	"github.com/ClubWeGo/videomicro/dal/pack"
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
		Favorite_count: 0, // 初始为0
		Comment_count:  0, // 初始为0
	}
	err = v.Create(video)

	if err != nil {
		return &videomicro.CreateVideoResp{
			Status: false,
		}, err
	}

	// 发布视频，更新作者的作品计数
	err = pack.AddCount(request.AuthorId)
	if err != nil {
		return &videomicro.CreateVideoResp{
			Status: true,
		}, err // 携带错误，但是创建成功
	}

	return &videomicro.CreateVideoResp{
		Status: true,
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

	respvideolist := make([]*videomicro.Video, len(videolist))

	for index, video := range videolist {
		respvideolist[index] = &videomicro.Video{
			Id:            int64(video.ID),
			Title:         video.Title,
			AuthorId:      video.Author_id,
			PlayUrl:       video.Play_url,
			CoverUrl:      video.Cover_url,
			FavoriteCount: video.Favorite_count,
			CommentCount:  video.Comment_count,
		}
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
	startTime := time.Unix(0, request.LatestTime)
	videolist, err := v.Where(v.CreatedAt.Lt(startTime)).Limit(int(request.Limit)).Order(v.CreatedAt.Desc()).Find()
	if err != nil {
		return &videomicro.GetVideosFeedResp{
			Status: true,
		}, err
	}

	respvideolist := make([]*videomicro.Video, len(videolist))

	for index, video := range videolist {
		respvideolist[index] = &videomicro.Video{
			Id:            int64(video.ID),
			Title:         video.Title,
			AuthorId:      video.Author_id,
			PlayUrl:       video.Play_url,
			CoverUrl:      video.Cover_url,
			FavoriteCount: video.Favorite_count,
			CommentCount:  video.Comment_count,
		}
	}

	var endTimeUnix = request.LatestTime // 没有条目，则时间还是之前的，交由hertz端处理
	if len(videolist) >= 1 {             // 搜到条目，则更新时间
		endTimeUnix = videolist[len(videolist)-1].CreatedAt.UnixNano()
	}

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
	videoAuthor, err := v.Select(v.Author_id).Where(v.ID.Eq(uint(request.VideoId))).First()
	if err != nil {
		return &videomicro.DeleteVideoResp{
			Status: false,
		}, err
	}
	_, err = v.Where(v.ID.Eq(uint(request.VideoId))).Delete(&model.Video{}) // 软删除
	if err != nil {
		return &videomicro.DeleteVideoResp{
			Status: false,
		}, err
	}

	err = pack.DecCount(videoAuthor.Author_id)
	if err != nil {
		return &videomicro.DeleteVideoResp{
			Status: true,
		}, err // 携带错误，但是删除成功
	}

	return &videomicro.DeleteVideoResp{
		Status: true,
	}, err
}

// UpdateVideoMethod implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) UpdateVideoMethod(ctx context.Context, request *videomicro.UpdateVideoReq) (resp *videomicro.UpdateVideoResp, err error) {
	// TODO: Your code here...

	v := query.Video
	// 只查询要更新的字段，后续增加字段这里需要修改
	video := v.Select(v.Title).Where(v.ID.Eq(uint(request.Id)))

	var title string
	if request.Title != nil {
		title = *request.Title
	}

	// 更新数据
	_, err = video.Updates(model.Video{
		Title: title,
	})
	if err != nil {
		return &videomicro.UpdateVideoResp{
			Status: false, // 更新失败
		}, err
	}

	return &videomicro.UpdateVideoResp{
		Status: true,
	}, err
}

// GetVideoAuthorIdMethod implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideoAuthorIdMethod(ctx context.Context, request *videomicro.GetVideoAuthorIdReq) (resp *videomicro.GetVideoAuthorIdResp, err error) {
	// TODO: Your code here...
	v := query.Video
	// 只查询要的字段，后续增加字段这里需要修改
	video, err := v.Select(v.Author_id).Where(v.ID.Eq(uint(request.Id))).First()
	if err != nil {
		return &videomicro.GetVideoAuthorIdResp{
			Status:   false,
			AuthorId: 0,
		}, err
	}

	return &videomicro.GetVideoAuthorIdResp{
		Status:   true,
		AuthorId: video.Author_id,
	}, nil
}

// GetVideoSetByIdSetMethod implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideoSetByIdSetMethod(ctx context.Context, request *videomicro.GetVideoSetByIdSetReq) (resp *videomicro.GetVideoSetByIdSetResp, err error) {
	// TODO: Your code here...
	v := query.Video

	idSet := request.IdSet

	// 切片互转有内存风险，暂采用最原始的方式转换id格式
	idSetUint := make([]uint, len(idSet))
	for index, id := range idSet {
		idSetUint[index] = uint(id)
	}

	// in 批量查询
	videos, err := v.Where(v.ID.In(idSetUint...)).Find()
	if err != nil {
		return &videomicro.GetVideoSetByIdSetResp{
			Status: false,
		}, err
	}

	// 批量转换格式
	respvideolist := make([]*videomicro.Video, len(videos))
	for index, video := range videos {
		respvideolist[index] = &videomicro.Video{
			Id:            int64(video.ID),
			Title:         video.Title,
			AuthorId:      video.Author_id,
			PlayUrl:       video.Play_url,
			CoverUrl:      video.Cover_url,
			FavoriteCount: video.Favorite_count,
			CommentCount:  video.Comment_count,
		}
	}

	return &videomicro.GetVideoSetByIdSetResp{
		Status:   true,
		VideoSet: respvideolist,
	}, nil
}

// GetVideoCountSetByIdUserSetMethod implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideoCountSetByIdUserSetMethod(ctx context.Context, request *videomicro.GetVideoCountSetByIdUserSetReq) (resp *videomicro.GetVideoCountSetByIdUserSetResp, err error) {
	// TODO: Your code here...
	vc := query.VideoCount

	idSet := request.AuthorIdSet

	// in 批量查询
	videoCounts, err := vc.Select(vc.Work_count).Where(vc.Author_id.In(idSet...)).Find()
	if err != nil {
		return &videomicro.GetVideoCountSetByIdUserSetResp{
			Status:   false,
			CountSet: []*videomicro.VideoCount{}, // 没查到为空
		}, err
	}

	// 批量转换格式
	respvideoCountslist := make([]*videomicro.VideoCount, len(videoCounts))
	for index, videoCount := range videoCounts {
		respvideoCountslist[index] = &videomicro.VideoCount{
			Id:    videoCount.Author_id,
			Count: videoCount.Work_count,
		}
	}

	return &videomicro.GetVideoCountSetByIdUserSetResp{
		Status:   true,
		CountSet: respvideoCountslist,
	}, nil
}
