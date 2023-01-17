package usecase

import "go-clean/domain/model"

type UserRepository interface {
	FindAll() ([]*model.User, error)
	FindByID(UID string) (*model.User, error)
	Create(*model.User) (*model.User, error)
	Update(*model.User) (*model.User, error)
	Delete(UID string) (*model.User, error)
}
