package dao

import "database/sql"

type Dao struct {
	db *sql.DB
}

func New(db *sql.DB) Dao {
	return Dao{db: db}
}
