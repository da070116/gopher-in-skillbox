package repository

import (
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

type Config struct {
	Name string
}

// NewSqliteDB - Sqlite connection
func NewSqliteDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite", cfg.Name)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
