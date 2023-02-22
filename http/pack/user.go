package pack

import (
	"github.com/ClubWeGo/simple-tiktok/biz/model/core"
	"github.com/ClubWeGo/usermicro/kitex_gen/usermicro"
	"strconv"
)

// User 将 usermicro.UserInfo 转换为 core.User，针对每个用户，需要预先把对作者的关注状态传入
func User(user *usermicro.UserInfo, isFollow bool) *core.User {
	return &core.User{
		ID:              user.Id,
		Name:            user.Name,
		FollowCount:     user.FollowCount,
		FollowerCount:   user.FollowerCount,
		IsFollow:        isFollow,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
		Signature:       user.Signature,
		TotalFavourited: strconv.FormatInt(user.TotalFavorited, 10),
		WorkCount:       user.WorkCount,
		FavoriteCount:   user.FavoriteCount,
	}
}
