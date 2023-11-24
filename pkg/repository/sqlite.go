package repository

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	Dsn    string
	Driver string
}

func NewSqliteDB(config *Config) (*sql.DB, error) {
	db, err := sql.Open(config.Driver, config.Dsn)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
