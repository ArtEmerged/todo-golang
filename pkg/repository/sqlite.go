package repository

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	userTable       = "users"
	todoListTable   = "todo_lists"
	usersListsTable = "users_lists"
	todoItemsTable  = "todo_items"
	listsItemsTable = "lists_items"
)

type Config struct {
	Driver string
	Dsn    string
}

func NewSqliteDB(config *Config) (*sql.DB, error) {
	db, err := sql.Open(config.Driver, config.Dsn)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("PRAGMA foreign_keys=ON;")
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
