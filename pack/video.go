package pack

import (
	"github.com/ClubWeGo/simple-tiktok/biz/model/core"
	"github.com/ClubWeGo/usermicro/kitex_gen/usermicro"
	"github.com/ClubWeGo/videomicro/kitex_gen/videomicro"
)

// Videos 将 videomicro.Video 列表转换为 core.Video 列表，针对每个用户，需要预先把用户对每个视频id的点赞状态和对作者id的关注状态传入
func Videos(videos []*videomicro.Video, authors []*usermicro.UserInfo, isFavorites map[int64]bool, isFollows map[int64]bool) []*core.Video {
	// 一个用户 id 对应一个用户信息，方便后续查询根据视频的作者 id 查询作者信息
	authorMap := make(map[int64]*usermicro.UserInfo, len(authors))
	for _, a := range authors {
		authorMap[a.Id] = a
	}

	res := make([]*core.Video, 0, len(videos))
	for _, v := range videos {
		res = append(res, Video(v, authorMap[v.AuthorId], isFavorites[v.Id], isFollows[v.AuthorId]))
	}
	return res
}

// Video 将 videomicro.Video 转换为 core.Video，针对每个用户，需要预先把用户对视频的点赞状态和对作者的关注状态传入
func Video(video *videomicro.Video, author *usermicro.UserInfo, isFavorite bool, isFollow bool) *core.Video {
	return &core.Video{
		ID:            video.Id,
		Author:        User(author, isFollow),
		PlayURL:       video.PlayUrl,
		CoverURL:      video.CoverUrl,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		IsFavorite:    isFavorite,
		Title:         video.Title,
	}
}
