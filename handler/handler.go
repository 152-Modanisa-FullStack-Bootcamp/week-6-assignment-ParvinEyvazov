package handler

import (
	"bank/model"
	"bank/service"
	"encoding/json"
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
	Err    string `json:"error"`
	Status int    `json:"status"`
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
		h.UpdateUser(w, r)
	default:
		sendError(&w, ErrorResponse{Err: "Undefined endpoint", Status: http.StatusBadRequest})
	}
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetUsers()
	if err != nil {
		sendError(&w, ErrorResponse{Err: "Error", Status: http.StatusNotFound})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	username := getUsername(r.URL.Path)

	balance, err := h.service.GetUserBalance(username)
	if err != nil {
		sendError(&w, ErrorResponse{Err: "Not found", Status: http.StatusNotFound})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(balance)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	username := getUsername(r.URL.Path)

	balance, err := h.service.CreateUser(username)
	if err != nil {
		sendError(&w, ErrorResponse{Err: "Server error", Status: http.StatusInternalServerError})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(balance)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	username := getUsername(r.URL.Path)
	_ = username

	bs := make([]byte, r.ContentLength)
	r.Body.Read(bs)

	var body model.UpdateBalanceBody
	if err := json.Unmarshal(bs, &body); err != nil {
		sendError(&w, ErrorResponse{Err: "Invalid body", Status: http.StatusBadRequest})
		return
	}

	balance, err := h.service.UpdateBalance(username, body)
	if err != nil {
		sendError(&w, ErrorResponse{Err: err.Error(), Status: http.StatusInternalServerError})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(balance)
}

func NewHandler(s service.IService) IHandler {
	return &Handler{service: s}
}

func sendError(w *http.ResponseWriter, err ErrorResponse) {
	(*w).WriteHeader(err.Status)
	(*w).Header().Set("Content-Type", "application/json")
	json.NewEncoder(*w).Encode(err)
}

func getUsername(path string) string {
	parameters := strings.Split(path, "/")
	return parameters[1]
}
