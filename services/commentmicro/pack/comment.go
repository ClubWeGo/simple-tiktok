package pack

import (
	"errors"
	"github.com/ClubWeGo/commentmicro/cmd/rpc"
	"github.com/ClubWeGo/commentmicro/dal/model"
	"github.com/ClubWeGo/commentmicro/kitex_gen/comment"
	"gorm.io/gorm"
)

//func Comments(favorite []*model.Favorite) []int64 {
//	favorites := make([]int64, 0)
//	for _, v := range favorite {
//		favorites = append(favorites, int64(v.VideoId))
//	}
//	return favorites
//}

// Comment pack Comments info.
func Comments(vs []*model.Comment) ([]*comment.Comment, error) {
	comments := make([]*comment.Comment, 0)
	for _, v := range vs {
		user, err := rpc.GetUserByID(v.UserID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		comments = append(comments, &comment.Comment{
			Id: int64(v.ID),
			User: &comment.User{
				Id:            int64(user.ID),
				Name:          user.Name,
				FollowCount:   &user.FollowerCount,
				FollowerCount: &user.FollowerCount,
				IsFollow:      false,
			},
			Content:    v.Content,
			CreateDate: v.CreatedAt.Format("01-02"),
		})
	}
	return comments, nil
}
