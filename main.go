package main

import (
	"bank/config"
	"bank/handler"
	"bank/repository"
	"bank/service"
	"fmt"
	"net/http"
)

func main() {

	repository := repository.NewRepository()
	conf := config.Get()

	service := service.NewService(repository, *conf)
	handler := handler.NewHandler(service)

	http.HandleFunc("/", handler.Gateway)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ERROR ON SERVING", err)
	}
}
