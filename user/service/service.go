package service

import (
	"errors"

	"github.com/arnoldtherigan15/final_project_golang/domain"
)

type service struct {
	repository domain.UserRepository
}

func NewService(repository domain.UserRepository) *service {
	return &service{repository}
}

func (s *service) Create(user *domain.User) (*domain.User, error) {
	user, err := s.repository.Create(user)
	if err != nil {
		return &domain.User{}, err
	}
	return user, nil
}

func (s *service) Update(ID int, user *domain.User) (bool, error) {
	userDB, err := s.repository.FindByID(ID)
	if err != nil {
		return false, err
	}
	if userDB.UserId == 0 {
		return false, errors.New("user not found")
	}
	userDB.Name = user.Name
	updatedUser, err := s.repository.Update(userDB)
	if err != nil {
		return false, err
	}
	return updatedUser, nil
}

func (s *service) FindByID(ID int) (*domain.User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return &domain.User{}, err
	}
	if user.UserId == 0 {
		return &domain.User{}, errors.New("user not found")
	}
	return user, nil
}

func (s *service) Delete(ID int) (bool, error) {
	user, err := s.FindByID(ID)
	if err != nil {
		return false, err
	}
	if user.UserId == 0 {
		return false, errors.New("user not found")
	}
	isDeleted, err := s.repository.Delete(user)
	if err != nil {
		return false, err
	}
	return isDeleted, nil
}
