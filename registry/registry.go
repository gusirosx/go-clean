package registry

import (
	"database/sql"
	"go-clean/adapters"
)

type registry struct {
	db *sql.DB
}

type Registry interface {
	NewAppController() adapters.AppController
}

func NewRegistry(db *sql.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() adapters.AppController {
	return adapters.AppController{
		User: r.NewUserController(),
	}
}
