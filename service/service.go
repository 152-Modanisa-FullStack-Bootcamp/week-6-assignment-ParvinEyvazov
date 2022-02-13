package service

import (
	"bank/config"
	"bank/model"
	"bank/repository"
	"errors"
)

type IService interface {
	GetUsers() (model.Users, error)
	GetUserBalance(username string) (float64, error)
	CreateUser(username string) (float64, error)
	UpdateBalance(username string, updateData model.UpdateBalanceBody) (float64, error)
}

type Service struct {
	repository repository.IRepository
	config     config.Config
}

func (s *Service) GetUsers() (model.Users, error) {
	return s.repository.GetUsers(), nil
}

func (s *Service) GetUserBalance(username string) (float64, error) {
	balance, ok := s.repository.GetUserBalance(username)
	if !ok {
		return balance, errors.New("undefined user")
	}

	return balance, nil
}

func (s *Service) CreateUser(username string) (float64, error) {
	balance, ok := s.repository.GetUserBalance(username)
	if !ok {
		return s.repository.CreateUser(username, s.config.InitialBalance)
	}

	return balance, nil
}

func (s *Service) UpdateBalance(username string, updateData model.UpdateBalanceBody) (float64, error) {
	balance, ok := s.repository.GetUserBalance(username)
	if !ok {
		return balance, errors.New("undefined user")
	}

	if (balance + updateData.Balance) < s.config.MinumumBalance {
		return balance, errors.New("invalid operation")
	}

	// good to go
	balance, err := s.repository.UpdateBalance(username, (balance + updateData.Balance))
	if err != nil {
		return balance, err
	}

	return balance, nil
}

func NewService(repository repository.IRepository, config config.Config) IService {
	return &Service{repository: repository, config: config}
}
