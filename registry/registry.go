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

//Note the dbHandlers map in every repository – that’s here so every repository can use every other repository without giving up on Dependency Injection – if some of the repositories use a different DbHandler implementation than others, then repositories using other repositories don’t need to know who uses what; it’s kind of a poor man’s Dependency Injection Container.
