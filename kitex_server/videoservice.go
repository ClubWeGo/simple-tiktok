package kitex_server

import (
	"context"
	"errors"
	"sync"

	"github.com/ClubWeGo/douyin/biz/model/core"
	"github.com/ClubWeGo/videomicro/kitex_gen/videomicro"
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

	// 批量查询视频的 被喜欢数 ，Favorite 从Favorite服务
	// 批量查询 favorite_count, total_favourited 从favorite服务: kitex_server.FavoriteClient.UserFavoriteCountMethod()
	// TODO: 结果要 videoId与对应的数据 map

	// 批量查询 is_follow, 从relation服务; 传入目标userID和currentUser

	// 批量查询 follow_count， follower_cout 从relation服务

	// 等待数据
	wgVideo.Wait()

	// // 处理协程错误
	var errSlice = []error{} // 防止外部设置的chan缓存不够造成阻塞，要求外部设置长度为1的error切片类型
	// err := <-respAuthorMapError
	// if err != nil {
	// 	errSlice = append(errSlice, err)
	// }

	// // TODO: 其他协程的错误处理

	errChan <- errSlice // 记录错误的切片，至少应该返回一个空切片，否则chan会阻塞

	// 更新数据
	videoLatestMap := make(map[int64]core.Video, len(idSet)) // 视频切片的id是没有重复的
	for _, id := range idSet {
		videoLatestMap[id] = core.Video{ // 视频id对应的Video存储查到的关键字段
			FavoriteCount: 0,     // TODO:从拿到的MAP数据更新
			CommentCount:  0,     // TODO:从拿到的MAP数据更新
			IsFavorite:    false, // TODO:从拿到的MAP数据更新
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
		go GetVideoLatestMap(authorIdSet, currentUserId, respLatestVideoMap, wg, respLatestVideoMapError)

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
		return resultList, *r.NextTime, nil
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
		go GetVideoLatestMap(authorIdSet, id, respLatestVideoMap, wg, respLatestVideoMapError)

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
