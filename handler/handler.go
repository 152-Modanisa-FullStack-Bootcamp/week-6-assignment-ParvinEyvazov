package handler

import (
	"bootcamp/hmw6/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type IHandler interface {
	Gateway(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	service service.IService
}

type ErrorResponse struct {
	Err string `json:"error"`
}

func (h *Handler) Gateway(w http.ResponseWriter, r *http.Request) {
	username := getUsername(r.URL.Path)

	switch {
	case username == "" && r.Method == "GET":
		h.GetUsers(w, r)
	case r.Method == "GET":
		h.GetBalance(w, r)
	case r.Method == "PUT":
		fmt.Println("PUT balance")
	case r.Method == "POST":
		fmt.Println("POST balance")
	default:
		undefinedEndpointError(&w)
	}
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.Users()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Error"))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	username := getUsername(r.URL.Path)

	balance, err := h.service.UserBalance(username)
	if err != nil {
		notFoundError(&w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(balance)
}

func NewHandler(s service.IService) IHandler {
	return &Handler{service: s}
}

func undefinedEndpointError(w *http.ResponseWriter) {
	(*w).WriteHeader(http.StatusBadGateway)
	(*w).Header().Set("Content-Type", "application/json")
	json.NewEncoder(*w).Encode(ErrorResponse{
		Err: "Undefined endpoint",
	})
}

func notFoundError(w *http.ResponseWriter) {
	(*w).WriteHeader(http.StatusNotFound)
	(*w).Header().Set("Content-Type", "application/json")
	json.NewEncoder(*w).Encode(ErrorResponse{
		Err: "Not found",
	})
}

func getUsername(path string) string {
	parameters := strings.Split(path, "/")
	return parameters[1]
}
