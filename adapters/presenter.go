package adapters

import (
	"go-clean/domain/model"
	"go-clean/usecase"
)

type userPresenter struct{}

func NewUserPresenter() usecase.UserPresenter {
	return &userPresenter{}
}

func (up *userPresenter) ResponseUsers(us []*model.User) []*model.User {
	for _, u := range us {
		u.Name = "Mr." + u.Name
	}
	return us
}

func (up *userPresenter) ResponseUser(user *model.User) *model.User {
	user.Name = "Mr." + user.Name
	return user
}
