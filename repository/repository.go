package repository

import (
	"bank/model"
)

var Users model.Users = map[string]float64{
	"parvin": 1,
}

type IRepository interface {
	GetUsers() model.Users
	GetUserBalance(username string) (balance float64, ok bool)
	CreateUser(username string, new_balance float64) (balance float64, err error)
	UpdateBalance(username string, new_balance float64) (balance float64, err error)
}

type MemoryRepository struct {
}

// GetUsers is returning all users from database
func (*MemoryRepository) GetUsers() model.Users {
	return Users
}

// GetUserBalance is finds users in the database and returns their balance
func (*MemoryRepository) GetUserBalance(username string) (balance float64, ok bool) {
	balance, ok = Users[username]
	return
}

// CreateUser is creating user in the database
func (*MemoryRepository) CreateUser(username string, new_balance float64) (balance float64, err error) {
	Users[username] = new_balance

	return new_balance, nil
}

// UpdateBalance is updating user`s balance in the database
func (*MemoryRepository) UpdateBalance(username string, new_balance float64) (balance float64, err error) {
	Users[username] = new_balance

	return new_balance, nil
}

// NewRepository is creating new repository
func NewRepository() IRepository {
	return &MemoryRepository{}
}
