package database

import (
	"database/sql"
	"errors"
)

var (
	ErrNoRowsAffected = errors.New("no rows affected")
)

type Database struct {
	db *sql.DB
}

func New(db *sql.DB) *Database {
	return &Database{
		db: db,
	}
}
