package repository

import (
	"bootcamp/hmw6/config"
	"bootcamp/hmw6/model"
)

var Users model.Users = map[string]float64{
	"parvin": 1,
}

type IRepository interface {
	GetUsers() model.Users
	GetUserBalance(username string) (balance float64, ok bool)
	UpdateUserBalance(username string) (balance float64, err error)
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

func (*MemoryRepository) UpdateUserBalance(username string) (balance float64, err error) {

	balance, ok := Users[username]

	if !ok {
		Users[username] = config.C.InitialBalance
		return Users[username], nil
	}

	return balance, nil
}

func NewRepository() IRepository {
	return &MemoryRepository{}
}
