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
	case username == "" && r.Method == http.MethodGet:
		h.GetUsers(w, r)
	case r.Method == http.MethodGet:
		h.GetBalance(w, r)
	case r.Method == http.MethodPut:
		h.CreateUser(w, r)
	case r.Method == http.MethodPost:
		fmt.Println("POST balance")
	default:
		undefinedEndpointError(&w)
	}
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Error"))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	username := getUsername(r.URL.Path)

	balance, err := h.service.GetUserBalance(username)
	if err != nil {
		notFoundError(&w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(balance)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	username := getUsername(r.URL.Path)

	balance, err := h.service.CreateUser(username)
	if err != nil {
		serverError(&w)
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

func serverError(w *http.ResponseWriter) {
	(*w).WriteHeader(http.StatusInternalServerError)
	(*w).Header().Set("Content-Type", "application/json")
	json.NewEncoder(*w).Encode(ErrorResponse{
		Err: "Server error",
	})
}

func getUsername(path string) string {
	parameters := strings.Split(path, "/")
	return parameters[1]
}
