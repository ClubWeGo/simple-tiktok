package service

import (
	"errors"

	"github.com/ClubWeGo/usermicro/dal/query"
)

// 为什么返回查询对象？更新用户信息的场景，需要用这个对象直接去更新，而不是数据结果
func QueryUserByIdOrName(id *int64, name *string) (query.IUserDo, error) {
	// 两个均不存在，无法查询
	if id == nil && name == nil {
		return nil, errors.New("needs valid id or name")
	}

	u := query.User

	var user query.IUserDo = nil
	// 优先使用id进行用户查询
	if id != nil {
		user = u.Where(u.ID.Eq(uint(*id)))
	} else if name != nil {
		user = u.Where(u.Name.Eq(*name))
	}
	return user, nil
}
