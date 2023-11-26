package repository

import (
	"database/sql"
	"todo"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}

type TodoList interface{}

type TodoItem interface{}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthSqlite(db),
	}
}
