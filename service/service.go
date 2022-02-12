package service

import (
	"bootcamp/hmw6/model"
	"bootcamp/hmw6/repository"
	"errors"
)

type IService interface {
	Users() (model.Users, error)
	UserBalance(username string) (float64, error)
}

type Service struct {
	repository repository.IRepository
}

func (s *Service) Users() (model.Users, error) {
	return s.repository.GetUsers(), nil
}

func (s *Service) UserBalance(username string) (float64, error) {
	balance, ok := s.repository.GetUserBalance(username)
	if !ok {
		return balance, errors.New("undefined user")
	}

	return balance, nil
}

func NewService(repository repository.IRepository) IService {
	return &Service{repository: repository}
}
