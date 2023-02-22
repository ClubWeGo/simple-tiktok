package pack

import (
	"github.com/ClubWeGo/simple-tiktok/services/favoritemicro/dal/model"
)

func Favorites(favorite []*model.Favorite) []int64 {
	favorites := make([]int64, 0)
	for _, v := range favorite {
		favorites = append(favorites, int64(v.VideoId))
	}
	return favorites
}
