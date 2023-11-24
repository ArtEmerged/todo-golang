package service

import "todo/pkg/repository"

type Authorization interface{}

type TodoList interface{}

type TodoItem interface{}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: repos.Authorization,
		TodoList:      repos.TodoList,
		TodoItem:      repos.TodoItem,
	}
}
