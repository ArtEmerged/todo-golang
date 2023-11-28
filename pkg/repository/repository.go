package repository

import (
	"database/sql"

	"todo"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
}

type TodoItem interface{}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthSqlite(db),
		TodoList:      NewTodoListSqlite(db),
	}
}
