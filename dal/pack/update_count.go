package pack

import (
	"github.com/ClubWeGo/videomicro/dal/model"
	"github.com/ClubWeGo/videomicro/dal/query"
)

func AddCount(authorId int64) error {
	vc := query.VideoCount

	// 先判断条目是否存在
	item, _ := vc.Select(vc.Author_id).Where(vc.Author_id.Eq(authorId)).First()
	// log.Println(item) // <nil>
	// log.Println(err)  // record not found
	if item == nil {
		err := vc.Create(&model.VideoCount{
			Author_id:  authorId,
			Work_count: 1,
		})
		if err != nil {
			return err
		}
	} else {
		_, err := vc.Where(vc.Author_id.Eq(authorId)).UpdateSimple(vc.Work_count.Add(1))
		if err != nil {
			return err
		}
	}
	return nil
}

func DecCount(authorId int64) error {
	vc := query.VideoCount

	// 条目肯定存在
	_, err := vc.Where(vc.Author_id.Eq(authorId)).UpdateSimple(vc.Work_count.Add(-1))
	if err != nil {
		return err
	}
	return nil
}
