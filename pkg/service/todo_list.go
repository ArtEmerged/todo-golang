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

func (s *TodoListsService) GetListById(userId, listId int) (todo.TodoList, error) {
	return s.repo.GetListById(userId, listId)
}

func (s *TodoListsService) Delete(userId, listId int) error {
	_, err := s.repo.GetListById(userId, listId)
	if err != nil {
		return err
	}
	return s.repo.Delete(listId)
}

func (s *TodoListsService) Update(userId, listId int, input todo.UpdateList) error {
	_, err := s.repo.GetListById(userId, listId)
	if err != nil {
		return err
	}
	if err = input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(listId, input)
}
