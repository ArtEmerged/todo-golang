package service

import (
	"todo"
	"todo/pkg/repository"
)

type TodoListsService struct {
	repo repository.TodoList
}

func NewTodoListsService(repo repository.TodoList) *TodoListsService {
	return &TodoListsService{repo: repo}
}

func (s *TodoListsService) Create(userId int, list todo.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListsService) GetAll(userId int) ([]todo.TodoList, error) {
	return s.repo.GetAll(userId)
}
