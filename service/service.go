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

// GetUsers fetches user from repository and returs to handler
func (s *Service) GetUsers() (model.Users, error) {
	return s.repository.GetUsers(), nil
}

// GetUserBalance fetches user balance, if user doesn't exist, return error
func (s *Service) GetUserBalance(username string) (float64, error) {
	balance, ok := s.repository.GetUserBalance(username)

	// if user doesn`t exist, send error
	if !ok {
		return balance, errors.New("undefined user")
	}

	return balance, nil
}

// CreateUser creates user, if already exist, returns its balance
func (s *Service) CreateUser(username string) (float64, error) {
	balance, ok := s.repository.GetUserBalance(username)

	// if user doesn`t exist, creates user
	if !ok {
		return s.repository.CreateUser(username, s.config.InitialBalance)
	}

	// if exists, returns balance of the user
	return balance, nil
}

// UpdateBalance is updating user balance if it can
func (s *Service) UpdateBalance(username string, updateData model.UpdateBalanceBody) (float64, error) {
	balance, ok := s.repository.GetUserBalance(username)

	// if user doesn`t exist, send error
	if !ok {
		return balance, errors.New("undefined user")
	}

	// if updated balance is smaller that minimum balance, sends error
	if (balance + updateData.Balance) < s.config.MinumumBalance {
		return balance, errors.New("invalid operation")
	}

	// if it can update balance, it updates it
	balance, err := s.repository.UpdateBalance(username, (balance + updateData.Balance))

	// if any repository error on updating, sends error
	if err != nil {
		return balance, err
	}

	// returns updated balance
	return balance, nil
}

// NewService is returning new service
func NewService(repository repository.IRepository, config config.Config) IService {
	return &Service{repository: repository, config: config}
}
