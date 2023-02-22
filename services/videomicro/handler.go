package main

import (
	"context"
<<<<<<< HEAD
	"crypto/md5"
	"encoding/hex"

	"github.com/ClubWeGo/usermicro/dal/model"
	"github.com/ClubWeGo/usermicro/dal/query"
	service "github.com/ClubWeGo/usermicro/dal/service"
	usermicro "github.com/ClubWeGo/usermicro/kitex_gen/usermicro"
)

func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// GetUserMethod implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserMethod(ctx context.Context, request *usermicro.GetUserReq) (resp *usermicro.GetUserResp, err error) {
	// TODO: Your code here...
	user, err := service.QueryUserByIdOrName(request.Id, request.Name)
	if err != nil { // 用户传入的参数有误，查询构造失败
		return &usermicro.GetUserResp{
			Status: false,
			// 查不到，可选字段统一不返回值
		}, err
	}

	// 真正查询用户
	userinstance, err := user.First()
	if err != nil { // 查询失败
		return &usermicro.GetUserResp{
=======
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
>>>>>>> video
			Status: false,
		}, err
	}

<<<<<<< HEAD
	// 查询成功
	return &usermicro.GetUserResp{
		Status: true,
		User: &usermicro.UserInfo{
			Id:              int64(userinstance.ID),
			Name:            userinstance.Name,
			Email:           &userinstance.Email,
			FollowCount:     userinstance.FollowCount,
			FollowerCount:   userinstance.FollowerCount,
			Avatar:          userinstance.Avatar,
			BackgroundImage: userinstance.BackgroundImage,
			Signature:       userinstance.Signature,
			TotalFavorited:  userinstance.TotalFavorited,
			WorkCount:       userinstance.WorkCount,
			FavoriteCount:   userinstance.FavoriteCount,
		},
	}, nil
}

// LoginUserMethod implements the UserServiceImpl interface.
func (s *UserServiceImpl) LoginUserMethod(ctx context.Context, request *usermicro.LoginUserReq) (resp *usermicro.LoginUserResp, err error) {
	// TODO: Your code here...
	u := query.User

	// 用用户名查询用户密码
	user, err := u.Select(u.ID, u.Password).Where(u.Name.Eq(request.Name)).First()
	if err != nil {
		return &usermicro.LoginUserResp{
=======
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
>>>>>>> video
			Status: false,
		}, err
	}

<<<<<<< HEAD
	if MD5(request.Password) == user.Password {
		id := int64(user.ID)             // 这样其实不好
		return &usermicro.LoginUserResp{ // 验证成功
			Status: true,
			UserId: &id,
		}, err
	}

	return &usermicro.LoginUserResp{
		Status: false,
	}, err
}

// CreateUserMethod implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUserMethod(ctx context.Context, request *usermicro.CreateUserReq) (resp *usermicro.CreateUserResp, err error) {
	// TODO: Your code here...
	u := query.User

	newUser := request.Newuser_

	// 先判断是否重名，用name查询；防止重名反复占用索引id（也是攻击的一种）
	user, err := u.Select(u.Name).Where(u.Name.Eq(newUser.Name)).First()
	// log.Println(user) // <nil>
	// log.Println(err)  // record not found

	if user != nil { // 已经存在用户名，注册失败
		return &usermicro.CreateUserResp{Status: false}, err
	}

	// 没有重复，继续创建
	var email, avatar, backgroundImage, signature string // 可选字段
	if newUser.Email != nil {
		email = *newUser.Email
	}
	if newUser.Avatar != nil {
		avatar = *newUser.Avatar
	}
	if newUser.BackgroundImage != nil {
		backgroundImage = *newUser.BackgroundImage
	}
	if newUser.Signature != nil {
		signature = *newUser.Signature
	}

	usermodel := &model.User{
		Name:            newUser.Name, // 必填字段
		Email:           email,
		Password:        MD5(request.Password), // MD5密码
		Avatar:          avatar,
		BackgroundImage: backgroundImage,
		Signature:       signature,
		FollowCount:     0,
		FollowerCount:   0,
		TotalFavorited:  0,
		WorkCount:       0,
		FavoriteCount:   0,
	}

	err = u.Create(usermodel)
	if err != nil {
		return &usermicro.CreateUserResp{
			Status: false,
		}, err
	}

	createdUser, err := u.Select(u.ID).Where(u.Name.Eq(newUser.Name)).First()
	if err != nil {
		return &usermicro.CreateUserResp{
			Status: false,
		}, err
	}

	id := int64(createdUser.ID) // 这样其实不好
	return &usermicro.CreateUserResp{
		Status: true,
		UserId: &id,
	}, nil
}

// UpdateUserMethod implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUserMethod(ctx context.Context, request *usermicro.UpdateUserReq) (resp *usermicro.UpdateUserResp, err error) {
	// TODO: Your code here...

	u := query.User

	// 用用户名查询要更新的字段
	id := uint(request.Id) // 这样其实不好
	// 只查询要更新的字段，后续增加字段这里需要修改
	user := u.Select(u.Email, u.Avatar, u.BackgroundImage, u.Signature, u.Password).Where(u.ID.Eq(id))
	if err != nil {
		return &usermicro.UpdateUserResp{
			Status: false, // 没查到
		}, err
	}

	updateData := request.UpdateData
	// 基础参数
	var email, avatar, backgroundImage, signature, password string
	if updateData.Avatar != nil {
		avatar = *updateData.Avatar
	}
	if updateData.Email != nil {
		email = *updateData.Email
	}
	if updateData.BackgroundImage != nil {
		backgroundImage = *updateData.BackgroundImage
	}
	if updateData.Signature != nil {
		signature = *updateData.Signature
	}

	// 密码
	if updateData.Password != nil {
		password = MD5(*updateData.Password)
	}

	// 更新数据
	_, err = user.Updates(model.User{
		Avatar:          avatar,
		Email:           email,
		BackgroundImage: backgroundImage,
		Signature:       signature,
		Password:        password,
	})
	if err != nil {
		return &usermicro.UpdateUserResp{
=======
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
>>>>>>> video
			Status: false, // 更新失败
		}, err
	}

<<<<<<< HEAD
	return &usermicro.UpdateUserResp{
		Status: true,
	}, nil
}

// GetUserSetByIdSetMethod implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserSetByIdSetMethod(ctx context.Context, request *usermicro.GetUserSetByIdSetReq) (resp *usermicro.GetUserSetByIdSetResp, err error) {
	// TODO: Your code here...

	u := query.User
=======
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
>>>>>>> video

	idSet := request.IdSet

	// 切片互转有内存风险，暂采用最原始的方式转换id格式
	idSetUint := make([]uint, len(idSet))
	for index, id := range idSet {
		idSetUint[index] = uint(id)
	}

	// in 批量查询
<<<<<<< HEAD
	users, err := u.Where(u.ID.In(idSetUint...)).Find()
	if err != nil {
		return &usermicro.GetUserSetByIdSetResp{
=======
	videos, err := v.Where(v.ID.In(idSetUint...)).Find()
	if err != nil {
		return &videomicro.GetVideoSetByIdSetResp{
>>>>>>> video
			Status: false,
		}, err
	}

	// 批量转换格式
<<<<<<< HEAD
	respvideolist := make([]*usermicro.UserInfo, len(users))
	for index, userinstance := range users {
		respvideolist[index] = &usermicro.UserInfo{
			Id:              int64(userinstance.ID),
			Name:            userinstance.Name,
			Email:           &userinstance.Email,
			FollowCount:     userinstance.FollowCount,
			FollowerCount:   userinstance.FollowerCount,
			Avatar:          userinstance.Avatar,
			BackgroundImage: userinstance.BackgroundImage,
			Signature:       userinstance.Signature,
			TotalFavorited:  userinstance.TotalFavorited,
			WorkCount:       userinstance.WorkCount,
			FavoriteCount:   userinstance.FavoriteCount,
		}
	}

	return &usermicro.GetUserSetByIdSetResp{
		Status:  true,
		UserSet: respvideolist,
	}, nil
}

// 缓存更新接口
// UpdateRelationMethod implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateRelationMethod(ctx context.Context, request *usermicro.UpdateRelationCacheReq) (resp *usermicro.UpdateRelationCacheResp, err error) {
	// TODO: Your code here...

	u := query.User

	// 用用户名查询要更新的字段
	id := uint(request.Id) // 这样其实不好
	// 只查询要更新的字段，后续增加字段这里需要修改
	user := u.Select(u.FollowCount, u.FollowerCount).Where(u.ID.Eq(id))
	if err != nil {
		return &usermicro.UpdateRelationCacheResp{
			Status: false,
		}, err
	}

	updateData := request.NewData_
	// 均为必填字段，可以直接更新
	user.Updates(model.User{
		FollowCount:   updateData.FollowCount,
		FollowerCount: updateData.FollowerCount,
	})

	return &usermicro.UpdateRelationCacheResp{
		Status: true,
	}, nil
}

// UpdateInteractionMethod implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateInteractionMethod(ctx context.Context, request *usermicro.UpdateInteractionCacheReq) (resp *usermicro.UpdateInteractionCacheResp, err error) {
	// TODO: Your code here...
	u := query.User

	// 用用户名查询要更新的字段
	id := uint(request.Id) // 这样其实不好
	// 只查询要更新的字段，后续增加字段这里需要修改
	user := u.Select(u.FavoriteCount, u.TotalFavorited).Where(u.ID.Eq(id))
	if err != nil {
		return &usermicro.UpdateInteractionCacheResp{
			Status: false,
		}, err
	}

	updateData := request.NewData_
	// 均为必填字段，可以直接更新
	user.Updates(model.User{
		FavoriteCount:  updateData.FavoriteCount,
		TotalFavorited: updateData.TotalFavorited,
	})

	return &usermicro.UpdateInteractionCacheResp{
		Status: true,
	}, nil
}

// UpdateWorkMethod implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateWorkMethod(ctx context.Context, request *usermicro.UpdateWorkCacheReq) (resp *usermicro.UpdateWorkCacheResp, err error) {
	// TODO: Your code here...
	u := query.User

	// 用用户名查询要更新的字段
	id := uint(request.Id) // 这样其实不好
	// 只查询要更新的字段，后续增加字段这里需要修改
	user := u.Select(u.WorkCount).Where(u.ID.Eq(id))
	if err != nil {
		return &usermicro.UpdateWorkCacheResp{
			Status: false,
		}, err
	}

	updateData := request.NewData_
	// 均为必填字段，可以直接更新
	user.Updates(model.User{
		WorkCount: updateData.WorkCount,
	})

	return &usermicro.UpdateWorkCacheResp{
		Status: true,
	}, nil
}

// TODO: select 限制字段
=======
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
	videoCounts, err := vc.Select(vc.Author_id, vc.Work_count).Where(vc.Author_id.In(idSet...)).Find()
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
>>>>>>> video
