package service

import (
	"todo"
	"todo/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetListById(userId, listId int) (todo.TodoList, error)
	Update(userId, listId int, input todo.UpdateList) error
	Delete(userId, listId int) error
}

type TodoItem interface {
	Create(userId, listId int, input todo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetItemById(userId, itemId int) (todo.TodoItem, error)
	Update(userId, itemId int, input todo.UpdateItem) error
	Delete(userId, itemId int) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListsService(repos.TodoList),
		TodoItem:      NewTodoItemsService(repos.TodoItem, repos.TodoList),
	}
}
