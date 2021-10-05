package service

import (
	"errors"

	"github.com/arnoldtherigan15/final_project_golang/domain"
)

type service struct {
	repository domain.TodoRepository
}

func NewService(repository domain.TodoRepository) *service {
	return &service{repository}
}

func (s *service) Create(todo *domain.Todo) (*domain.Todo, error) {
	todo, err := s.repository.Create(todo)
	if err != nil {
		return &domain.Todo{}, err
	}
	return todo, nil
}

func (s *service) FindAll() ([]*domain.Todo, error) {
	todos, err := s.repository.FindAll()
	if err != nil {
		return todos, err
	}
	return todos, nil
}

func (s *service) Update(ID int, todo *domain.Todo) (bool, error) {
	todoDB, err := s.repository.FindByID(ID)
	if err != nil {
		return false, err
	}
	if todoDB.ID == 0 {
		return false, errors.New("todo not found")
	}
	todoDB.Title = todo.Title
	todoDB.Description = todo.Description
	todoDB.DueDate = todo.DueDate
	todoDB.PersonInCharge = todo.PersonInCharge
	todoDB.Status = todo.Status
	updatedTodo, err := s.repository.Update(todoDB)
	if err != nil {
		return false, err
	}
	return updatedTodo, nil
}

func (s *service) FindByID(ID int) (*domain.Todo, error) {
	todo, err := s.repository.FindByID(ID)
	if err != nil {
		return &domain.Todo{}, err
	}
	if todo.ID == 0 {
		return &domain.Todo{}, errors.New("todo not found")
	}
	return todo, nil
}

func (s *service) Delete(ID int) (bool, error) {
	todo, err := s.FindByID(ID)
	if err != nil {
		return false, err
	}
	if todo.ID == 0 {
		return false, errors.New("todo not found")
	}
	isDeleted, err := s.repository.Delete(todo)
	if err != nil {
		return false, err
	}
	return isDeleted, nil
}
