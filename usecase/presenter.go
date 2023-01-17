package usecase

import "go-clean/domain/model"

type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
	ResponseUser(u *model.User) *model.User
}
