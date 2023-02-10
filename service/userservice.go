package service

import (
	"errors"

	"github.com/ClubWeGo/usermicro/dal/query"
)

func QueryUserByIdOrName(id *int64, name *string) (query.IUserDo, error) {
	// by ID prioritily
	if id == nil && name == nil {
		return nil, errors.New("needs valid id or name")
	}

	u := query.User

	var user query.IUserDo = nil
	// use id to search user prioritily
	if id != nil {
		findeduser := u.Where(u.ID.Eq(uint(*id)))
		user = findeduser
	} else if name != nil {
		findeduser := u.Where(u.Name.Eq(*name))
		user = findeduser // 在u.where那里直接赋值，接收不到
	}
	return user, nil
}

func QueryUserByIdOrNameOrEmail(id *int64, name *string, email *string) (query.IUserDo, error) {
	// by ID prioritily
	if id == nil && name == nil && email == nil {
		return nil, errors.New("needs valid id or name or email")
	}

	u := query.User

	var user query.IUserDo
	// use id to search user prioritily
	if id != nil {
		findeduser := u.Where(u.ID.Eq(uint(*id)))
		user = findeduser
	} else if name != nil {
		findeduser := u.Where(u.Name.Eq(*name))
		user = findeduser
	} else if email != nil {
		findeduser := u.Where(u.Email.Eq(*email))
		user = findeduser
	}
	return user, nil
}
