package service

import (
	"bootcamp/hmw6/data"
	"bootcamp/hmw6/model"
	"errors"
)

type IService interface {
	Users() (model.Users, error)
	UserBalance(username string) (float64, error)
}

type Service struct {
}

func (*Service) Users() (model.Users, error) {
	return data.GetUsers(), nil
}

func (*Service) UserBalance(username string) (float64, error) {
	balance, ok := data.GetUserBalance(username)
	if !ok {
		return balance, errors.New("undefined user")
	}

	return balance, nil
}

func NewService() IService {
	return &Service{}
}
