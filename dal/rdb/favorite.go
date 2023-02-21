package rdb

import (
	"context"
	"fmt"
	"github.com/ClubWeGo/favoritemicro/dal"
	"github.com/ClubWeGo/favoritemicro/pkg/errno"
	"github.com/redis/go-redis/v9"
	"strconv"
)

// GetAuthorId 从缓存中获取视频作者id
func GetAuthorId(ctx context.Context, vid int64) (int64, error) {
	res, err := dal.Redis.HGet(ctx, "video_author", keyVideo(vid)).Result()
	if err != nil {
		return 0, err
	}
	aid, err := strconv.ParseInt(res, 10, 64)
	if err != nil {
		return 0, err
	}
	return aid, nil
}

// SetAuthorId 缓存视频作者id
func SetAuthorId(ctx context.Context, vid int64, aid int64) {
	_, err := dal.Redis.HSet(ctx, "video_author", keyVideo(vid), keyUser(aid)).Result()
	if err != nil {
		fmt.Println(errno.DBErr.WithMessage(err.Error()))
	}
}

// GetFavoriteRelation 获取点赞状态
func GetFavoriteRelation(ctx context.Context, uid int64, vid int64) (bool, error) {
	result, err := dal.Redis.SIsMember(ctx, keyUser(uid), vid).Result()
	if err != nil {
		return false, err
	}
	if err != nil {
		return false, errno.DBErr.WithMessage(err.Error())
	}
	return result, nil
}

func GetFavoriteRelations(ctx context.Context, uid int64, vids []int64) (map[int64]bool, error) {
	res := make(map[int64]bool)
	for _, vid := range vids {
		status, err := dal.Redis.SIsMember(ctx, keyUser(uid), vid).Result()
		if err != nil {
			return nil, errno.DBErr.WithMessage(err.Error())
		}
		res[vid] = status
	}
	return res, nil
}

func AddFavorite(ctx context.Context, uid int64, vid int64, aid int64) error {
	// 事务函数
	txf := func(tx *redis.Tx) error {
		result, err := dal.Redis.SIsMember(ctx, keyUser(uid), vid).Result()
		if err != nil {
			return errno.DBErr.WithMessage(err.Error())
		}
		if result {
			return errno.RecordAlreadyExistErr.WithMessage("用户已经点赞过了")
		}
		// 事务操作，添加用户点赞记录，作者点赞数+1，视频点赞数+1
		_, err = tx.TxPipelined(
			ctx, func(pipe redis.Pipeliner) error {
				pipe.SAdd(ctx, keyUser(uid), vid)
				pipe.Incr(ctx, keyAuthor(aid))
				pipe.Incr(ctx, keyVideo(vid))
				return nil
			})
		return err
	}
	for i := 0; i < 3; i++ {
		err := dal.Redis.Watch(ctx, txf, "")
		if err == nil {
			return nil
		}
		if err == redis.TxFailedErr {
			continue
		}
		return errno.DBErr.WithMessage(err.Error())
	}
	return errno.DBErr.WithMessage("redis事务执行失败")
}

func DeleteFavorite(ctx context.Context, uid int64, vid int64, aid int64) error {
	txf := func(tx *redis.Tx) error {
		result, err := dal.Redis.SIsMember(ctx, keyUser(uid), vid).Result()
		if err != nil {
			return errno.DBErr.WithMessage(err.Error())
		}
		if !result {
			return errno.RecordAlreadyExistErr.WithMessage("重复取消，请求无效")
		}
		_, err = tx.TxPipelined(
			ctx, func(pipe redis.Pipeliner) error {
				pipe.SRem(ctx, keyUser(uid), vid)
				pipe.Decr(ctx, keyAuthor(aid))
				pipe.Decr(ctx, keyVideo(vid))
				return nil
			})
		return err
	}
	for i := 0; i < 3; i++ {
		err := dal.Redis.Watch(ctx, txf, "")
		if err == nil {
			return nil
		}
		if err == redis.TxFailedErr {
			continue
		}
		return err
	}
	return errno.DBErr.WithMessage("redis事务执行失败")
}

// GetFavoriteList get favorite video list
func GetFavoriteList(ctx context.Context, uid int64) ([]int64, error) {
	result, err := dal.Redis.SMembers(ctx, keyUser(uid)).Result()
	if err != nil {
		return nil, errno.DBErr.WithMessage(err.Error())
	}
	return StringSliceToInt64Slice(result), nil
}

// CountVideoFavorite 计算视频被点赞数
func CountVideoFavorite(ctx context.Context, vid int64) (int64, error) {
	cnt, err := dal.Redis.Get(ctx, keyVideo(vid)).Int64()
	if err != nil {
		return 0, errno.DBErr.WithMessage(err.Error())
	}
	return cnt, nil
}

// CountUserFavorite 计算用户点赞数和被点赞数
func CountUserFavorite(ctx context.Context, uid int64) (int64, int64, error) {
	favoriteCnt, err := dal.Redis.SCard(ctx, keyUser(uid)).Result()
	if err != nil {
		return 0, 0, errno.DBErr.WithMessage(err.Error())
	}
	favoritedCnt, err := dal.Redis.Get(ctx, keyAuthor(uid)).Int64()
	if err != nil {
		return 0, 0, errno.DBErr.WithMessage(err.Error())
	}
	return favoriteCnt, favoritedCnt, nil
}

func CountVideosFavorite(ctx context.Context, vids []int64) (map[int64]int64, error) {
	res := make(map[int64]int64)
	var err error
	for _, v := range vids {
		res[v], err = dal.Redis.Get(ctx, keyVideo(v)).Int64()
		if err != nil {
			return nil, errno.DBErr.WithMessage(err.Error())
		}
	}
	return res, nil
}

func CountUsersFavorite(ctx context.Context, uids []int64) (map[int64][]int64, error) {
	var favoriteCnt int64
	var favoritedCnt int64
	var err error

	res := make(map[int64][]int64)
	for _, u := range uids {
		favoriteCnt, err = dal.Redis.SCard(ctx, keyUser(u)).Result()
		if err != nil {
			return nil, errno.DBErr.WithMessage(err.Error())
		}
		favoritedCnt, err = dal.Redis.Get(ctx, keyAuthor(u)).Int64()
		if err != nil {
			return nil, errno.DBErr.WithMessage(err.Error())
		}
		res[u] = append(res[u], favoriteCnt)
		res[u] = append(res[u], favoritedCnt)
	}
	return res, nil
}

func keyUser(uid int64) string {
	return fmt.Sprintf("u%d", uid)
}

func keyVideo(vid int64) string {
	return fmt.Sprintf("v%d", vid)
}

func keyAuthor(aid int64) string {
	return fmt.Sprintf("a%d", aid)
}

func StringSliceToInt64Slice(strs []string) []int64 {
	var res []int64
	for _, str := range strs {
		i, _ := strconv.ParseInt(str, 10, 64)
		res = append(res, i)
	}
	return res
}
