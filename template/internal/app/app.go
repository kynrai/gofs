package app

import "module/placeholder/internal/db"

type App struct {
	db db.DB
}

func New(db db.DB) App {
	return App{db: db}
}
