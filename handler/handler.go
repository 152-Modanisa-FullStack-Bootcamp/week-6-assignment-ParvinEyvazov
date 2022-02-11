package handler

import (
	"bootcamp/hmw6/service"
	"encoding/json"
	"net/http"
)

type IHandler interface {
	Users(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	service service.IService
}

func (h *Handler) Users(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetUsers()
	if err != nil {
		w.Write([]byte("Error"))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func NewHandler(s service.IService) IHandler {
	return &Handler{service: s}
}
