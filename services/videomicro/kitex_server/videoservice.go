package kitex_server

import (
	"context"
	"errors"
	"sync"

	"github.com/ClubWeGo/douyin/biz/model/core"
	"github.com/ClubWeGo/simple-tiktok/services/videomicro/kitex_gen/videomicro"
)

// utils
// 将VideoCount集合转为map
func ConvertVideoCountSetToMap(setData []*videomicro.VideoCount) map[int64]videomicro.VideoCount {
	dataMap := make(map[int64]videomicro.VideoCount, len(setData))
	for _, data := range setData {
		dataMap[data.Id] = *data
	}
	return dataMap
}

// 协程接口

// 通过GetVideoCountSetByIdUserSetMethod获取WorkCount的集合，然后转为Map
func GetVideoCountMap(idSet []int64, respVideoCountMap chan map[int64]videomicro.VideoCount, wg *sync.WaitGroup, errChan chan error) {
	defer wg.Done()

	r, err := Videoclient.GetVideoCountSetByIdUserSetMethod(context.Background(), &videomicro.GetVideoCountSetByIdUserSetReq{
		AuthorIdSet: idSet,
	})
	if err != nil {
		respVideoCountMap <- map[int64]videomicro.VideoCount{}
		errChan <- err
		return
	}
	if r.Status {

		rVideoCountMap := ConvertVideoCountSetToMap(r.CountSet)
		respVideoCountMap <- rVideoCountMap
		errChan <- nil
		return // 成功
	}
	respVideoCountMap <- map[int64]videomicro.VideoCount{}
	errChan <- errors.New("userservice GetVideoCountSetByIDUserSet error: 微服务调用成功，但是返回状态显示失败")
}

// 用于获取视频切片中 被喜欢数，评论数，用户是否喜欢 的最新数据；写成协程，是为了feed那块可以并发查询
// 本函数根据查到的信息，构建Video切片，其中只有Video.Id和本函数查询的数据为真，其余数据以视频流查到的为准
func GetVideoLatestMap(idSet []int64, currentUser int64, respVideoMap chan map[int64]core.Video, wg *sync.WaitGroup, errChan chan []error) {
	defer wg.Done() // 外部的wg

	wgVideo := &sync.WaitGroup{} // 本函数子协程的wg

	// 批量查询视频的 被喜欢数 ，传入视频id的切片，返回对应的FavoriteCount的切片（需携带对应视频id） 从Favorite服务
	respVideosFavoriteCountMap := make(chan map[int64]int64, 1)
	defer close(respVideosFavoriteCountMap)
	respVideosFavoriteCountMapError := make(chan error, 1)
	defer close(respVideosFavoriteCountMapError)
	wgVideo.Add(1)
	go GetVideosFavoriteCountMap(idSet, respVideosFavoriteCountMap, wgVideo, respVideosFavoriteCountMapError)

	// 批量查询视频的评论数，传入视频id的切片，返回对应的评论数（需携带对应视频id），从comment服务
	// 在此处类似上边的写法，写评论数的实现
	respCommentCountMap := make(chan map[int64]int64, 1)
	defer close(respCommentCountMap)
	respCommentCountMapError := make(chan error, 1)
	defer close(respCommentCountMapError)
	wgVideo.Add(1)
	go GetCommentCountMap(idSet, respCommentCountMap, wgVideo, respCommentCountMapError)

	// 批量查询 is_favorite, 传入目标视频id切片和currentUser查is_favorite的切片(结果需要携带视频id，douyin里后续需要转成map)：从favorite;
	respIsFavoriteMap := make(chan map[int64]bool, 1)
	defer close(respIsFavoriteMap)
	respIsFavoriteMapError := make(chan error, 1)
	defer close(respIsFavoriteMapError)
	wgVideo.Add(1)
	go GetIsFavoriteMap(idSet, currentUser, respIsFavoriteMap, wgVideo, respIsFavoriteMapError)

	// 等待数据
	wgVideo.Wait()

	var errSlice = []error{}
	VideosFavoriteCountMap := <-respVideosFavoriteCountMap
	err := <-respVideosFavoriteCountMapError
	if err != nil {
		errSlice = append(errSlice, err)
	}

	IsFavoriteMap := <-respIsFavoriteMap
	err = <-respIsFavoriteMapError
	if err != nil {
		errSlice = append(errSlice, err)
	}

	VideosCommentCountMap := <-respCommentCountMap
	err = <-respCommentCountMapError
	if err != nil {
		errSlice = append(errSlice, err)
	}

	errChan <- errSlice // 记录错误的切片，至少应该返回一个空切片，否则chan会阻塞

	// 更新数据
	videoLatestMap := make(map[int64]core.Video, len(idSet)) // 视频切片的id是没有重复的
	for _, id := range idSet {
		videoLatestMap[id] = core.Video{ // 视频id对应的Video存储查到的关键字段
			FavoriteCount: VideosFavoriteCountMap[id], //
			CommentCount:  VideosCommentCountMap[id],  //
			IsFavorite:    IsFavoriteMap[id],          //
		}
	}
	respVideoMap <- videoLatestMap // 返回数据
}

// 业务接口

// 异步调用各种微服务获取feed流以及最新的信息
func GetFeed(latestTime int64, currentUserId int64, limit int32) (resultList []*core.Video, nextTime int64, err error) {
	// currentUserId 用于登录用户刷视频的时候，看是否关注过视频作者

	r, err := Videoclient.GetVideosFeedMethod(context.Background(), &videomicro.GetVideosFeedReq{LatestTime: latestTime, Limit: limit})
	if err != nil {
		return []*core.Video{}, 0, err
	}

	if r.Status {
		authorIdSet := make([]int64, len(r.VideoList))
		for index, video := range r.VideoList {
			authorIdSet[index] = video.AuthorId
		}
		videoIdSet := make([]int64, len(r.VideoList))
		for index, video := range r.VideoList {
			videoIdSet[index] = video.Id
		}

		wg := &sync.WaitGroup{}

		// 各种chan千万别写错，否则会引起各种读或者写的阻塞（发生阻塞一般都是这种情况）

		// 获取最新的用户信息
		respLatestAuthorMap := make(chan map[int64]core.User, 1)
		defer close(respLatestAuthorMap)
		respLatestAuthorMapError := make(chan []error, 1)
		defer close(respLatestAuthorMapError)
		wg.Add(1)
		go GetUserLatestMap(authorIdSet, currentUserId, respLatestAuthorMap, wg, respLatestAuthorMapError)

		// 获取视频的最新信息
		respLatestVideoMap := make(chan map[int64]core.Video, 1)
		defer close(respLatestVideoMap)
		respLatestVideoMapError := make(chan []error, 1)
		defer close(respLatestVideoMapError)
		wg.Add(1)
		go GetVideoLatestMap(videoIdSet, currentUserId, respLatestVideoMap, wg, respLatestVideoMapError)

		// 等待数据
		wg.Wait()

		// 处理协程错误
		AuthorMap := <-respLatestAuthorMap
		errSlice := <-respLatestAuthorMapError
		for _, errItem := range errSlice {
			if errItem != nil {
				return []*core.Video{}, 0, errItem
			}
		}

		videoMap := <-respLatestVideoMap
		errSlice = <-respLatestVideoMapError
		for _, errItem := range errSlice {
			if errItem != nil {
				return []*core.Video{}, 0, errItem
			}
		}

		// 拼接结果
		resultList = make([]*core.Video, len(r.VideoList))
		for index, video := range r.VideoList {
			// TODO:没有查询到的错误处理
			author := AuthorMap[video.AuthorId]

			// TODO:设置机制，慢速同步其他服务的最新数据到user服务的主表，video的主表

			resultList[index] = &core.Video{
				ID:            video.Id,
				Author:        &author,
				PlayURL:       video.PlayUrl,
				CoverURL:      video.CoverUrl,
				FavoriteCount: videoMap[video.Id].FavoriteCount,
				CommentCount:  videoMap[video.Id].CommentCount,
				IsFavorite:    videoMap[video.Id].IsFavorite,
				Title:         video.Title,
			}
		}
		return resultList, *r.NextTime, nil // 成功
	}
	return []*core.Video{}, 0, errors.New("向kitex请求feed失败")
}

// 通过userid获取用户的发布列表
func GetVideosByAuthorId(id int64) (resultList []*core.Video, err error) {
	r, err := Videoclient.GetVideosByAuthorIdMethod(context.Background(), &videomicro.GetVideosByAuthorIdReq{
		AuthorId: id,
	})
	if err != nil {
		return []*core.Video{}, err
	}

	if r.Status {
		authorIdSet := []int64{id} // 只有作者本人
		videoIdSet := make([]int64, len(r.VideoList))
		for index, video := range r.VideoList {
			videoIdSet[index] = video.Id
		}

		wg := &sync.WaitGroup{}

		// 获取最新的用户信息
		respLatestAuthorMap := make(chan map[int64]core.User, 1)
		defer close(respLatestAuthorMap)
		respLatestAuthorMapError := make(chan []error, 1)
		defer close(respLatestAuthorMapError)
		wg.Add(1)
		go GetUserLatestMap(authorIdSet, id, respLatestAuthorMap, wg, respLatestAuthorMapError)

		// 获取视频的最新信息
		respLatestVideoMap := make(chan map[int64]core.Video, 1)
		defer close(respLatestVideoMap)
		respLatestVideoMapError := make(chan []error, 1)
		defer close(respLatestVideoMapError)
		wg.Add(1)
		go GetVideoLatestMap(videoIdSet, id, respLatestVideoMap, wg, respLatestVideoMapError)

		// 等待数据
		wg.Wait()

		// 处理协程错误
		AuthorMap := <-respLatestAuthorMap
		// log.Println(AuthorMap)
		errSlice := <-respLatestAuthorMapError
		for _, errItem := range errSlice {
			if errItem != nil {
				return []*core.Video{}, errItem
			}
		}

		videoMap := <-respLatestVideoMap
		// log.Println(VideoMap)
		errSlice = <-respLatestVideoMapError
		for _, errItem := range errSlice {
			if errItem != nil {
				return []*core.Video{}, errItem
			}
		}

		// 拼接结果
		resultList = make([]*core.Video, len(r.VideoList))
		for index, video := range r.VideoList {
			// TODO:没有查询到的错误处理
			author := AuthorMap[video.AuthorId] // 只有作者本人，查map其实无所谓
			// TODO:设置机制，慢速同步其他服务的最新数据到user服务的主表，video的主表

			resultList[index] = &core.Video{
				ID:            video.Id,
				Author:        &author,
				PlayURL:       video.PlayUrl,
				CoverURL:      video.CoverUrl,
				FavoriteCount: videoMap[video.Id].FavoriteCount,
				CommentCount:  videoMap[video.Id].CommentCount,
				IsFavorite:    videoMap[video.Id].IsFavorite,
				Title:         video.Title,
			}
		}

		return resultList, nil
	}
	return []*core.Video{}, errors.New("向kitex请求作者的发布信息失败")
}

// 发布视频接口
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
