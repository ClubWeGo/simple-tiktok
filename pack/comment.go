package pack

import (
	"github.com/ClubWeGo/simple-tiktok/biz/model/core"
	"github.com/ClubWeGo/simple-tiktok/biz/model/interaction"
	"github.com/ClubWeGo/simple-tiktok/services/commentmicro/kitex_gen/comment"
)

func Commentlist(cs []*comment.Comment) ([]*interaction.Comment, error) {
	comments := make([]*interaction.Comment, 0)
	for _, v := range cs {
		comments = append(comments, &interaction.Comment{
			ID: v.Id,
			User: &core.User{
				ID:            v.Id,
				Name:          v.User.Name,
				FollowCount:   *v.User.FollowCount,
				FollowerCount: *v.User.FollowerCount,
				IsFollow:      v.User.IsFollow,
			},
			Content:    v.Content,
			CreateDate: v.CreateDate,
		})
	}
	return comments, nil
}
