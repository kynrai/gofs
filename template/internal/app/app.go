package app

import (
	"module/placeholder/internal/dao"
)

type App struct {
	dao dao.Dao
}

func New(dao dao.Dao) App {
	return App{dao: dao}
}
