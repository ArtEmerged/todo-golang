package service

import (
	"todo"
	"todo/pkg/repository"
)

type TodoItemsService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemsService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemsService {
	return &TodoItemsService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemsService) Create(userId, listId int, input todo.TodoItem) (int, error) {
	_, err := s.listRepo.GetListById(userId, listId)
	if err != nil {
		return 0, err
	}
	return s.repo.Create(listId, input)
}

func (s *TodoItemsService) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	_, err := s.listRepo.GetListById(userId, listId)
	if err != nil {
		return nil, err
	}
	return s.repo.GetAll(listId)
}

func (s *TodoItemsService) GetItemById(userId, itemId int) (todo.TodoItem, error) {
	return s.repo.GetItemById(userId, itemId)
}

func (s *TodoItemsService) Delete(userId, itemId int) error {
	_, err := s.repo.GetItemById(userId, itemId)
	if err != nil {
		return err
	}
	return s.repo.Delete(itemId)
}

func (s *TodoItemsService) Update(userId, itemId int, input todo.UpdateItem) error {
	_, err := s.repo.GetItemById(userId, itemId)
	if err != nil {
		return err
	}
	return s.repo.Update(itemId, input)
}
