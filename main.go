package main

import (
	_ "bootcamp/hmw6/config"
	"bootcamp/hmw6/handler"
	"bootcamp/hmw6/repository"
	"bootcamp/hmw6/service"
	"fmt"
	"net/http"
)

func main() {

	repository := repository.NewRepository()
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	http.HandleFunc("/", handler.Gateway)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ERROR ON SERVING", err)
	}
}
