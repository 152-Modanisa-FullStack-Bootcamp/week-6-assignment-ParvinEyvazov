package data

import "bootcamp/hmw6/model"

var Users model.Users = map[string]float64{
	"parvin": 1,
}

func GetUsers() model.Users {
	return Users
}

func GetUserBalance(username string) (balance float64, ok bool) {
	balance, ok = Users[username]
	return
}
