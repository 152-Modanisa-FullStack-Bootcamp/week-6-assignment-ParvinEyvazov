package repository

import "bootcamp/hmw6/model"

var Users model.Users = map[string]float64{
	"parvin": 1,
}

type IRepository interface {
	GetUsers() model.Users
	GetUserBalance(username string) (balance float64, ok bool)
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

func NewRepository() IRepository {
	return &MemoryRepository{}
}
