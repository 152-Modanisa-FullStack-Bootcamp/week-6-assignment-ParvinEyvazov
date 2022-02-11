package service

import "bootcamp/hmw6/model"

type IService interface {
	GetUsers() (model.Users, error)
}

type Service struct {
}

func (*Service) GetUsers() (model.Users, error) {
	return model.Users{
		{
			Username: "Parvin",
			Balance:  452,
		},
	}, nil
}

func NewService() IService {
	return &Service{}
}
