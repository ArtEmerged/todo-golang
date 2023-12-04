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
	GetListById(userId, listId int) (todo.TodoList, error)
	Delete(listId int) error
	Update(listId int, input todo.UpdateList) error
}

type TodoItem interface {
	Create(listId int, input todo.TodoItem) (int, error)
	GetAll(listId int) ([]todo.TodoItem, error)
	GetItemById(userId, itemId int) (todo.TodoItem, error)
	Update(itemId int, input todo.UpdateItem) error
	Delete(itemId int) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthSqlite(db),
		TodoList:      NewTodoListSqlite(db),
		TodoItem:      NewTodoItemSqlite(db),
	}
}
