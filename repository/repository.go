package repository

import (
	"bootcamp/hmw6/model"
)

var Users model.Users = map[string]float64{
	"parvin": 1,
}

type IRepository interface {
	GetUsers() model.Users
	GetUserBalance(username string) (balance float64, ok bool)
	UpdateUserBalance(username string, new_balance float64) (balance float64, err error)
}

type MemoryRepository struct {
}

func (*MemoryRepository) GetUsers() model.Users {
	return Users
}

func (*MemoryRepository) GetUserBalance(username string) (balance float64, ok bool) {
	balance, ok = Users[username]
	return
}

func (*MemoryRepository) UpdateUserBalance(username string, new_balance float64) (balance float64, err error) {
	Users[username] = new_balance

	return new_balance, nil
}

func NewRepository() IRepository {
	return &MemoryRepository{}
}
