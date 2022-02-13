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
	GetUsers(w http.ResponseWriter, r *http.Request)
	GetBalance(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	service service.IService
}

type ErrorResponse struct {
	Err    string `json:"error"`
	Status int    `json:"status"`
}

// Initial gateway to separate endpoints by their methods
func (h *Handler) Gateway(w http.ResponseWriter, r *http.Request) {
	username := getUsername(r.URL.Path)

	switch {
	// GET /
	case username == "" && r.Method == http.MethodGet:
		h.GetUsers(w, r)

	// GET /:username
	case username != "" && r.Method == http.MethodGet:
		h.GetBalance(w, r)

	// PUT /:username
	case username != "" && r.Method == http.MethodPut:
		h.CreateUser(w, r)

	// POST /:username
	case username != "" && r.Method == http.MethodPost:
		h.UpdateUser(w, r)

	// Undefined endpoint
	default:
		sendError(&w, ErrorResponse{Err: "Undefined endpoint", Status: http.StatusBadRequest})
	}
}

// GetUsers is fetching all users
func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetUsers()

	// any error on fetching users
	if err != nil {
		sendError(&w, ErrorResponse{Err: "Error", Status: http.StatusNotFound})
		return
	}

	// send response as json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetBalance is fetching balance information by username
func (h *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {

	// get username from the path
	username := getUsername(r.URL.Path)

	balance, err := h.service.GetUserBalance(username)

	// if cannot find user
	if err != nil {
		sendError(&w, ErrorResponse{Err: "Not found", Status: http.StatusNotFound})
		return
	}

	// send response as json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(balance)
}

// CreateUser is creating user with initial balance amount
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	// get username from the path
	username := getUsername(r.URL.Path)

	balance, err := h.service.CreateUser(username)

	// server problem on getting OR creating user
	if err != nil {
		sendError(&w, ErrorResponse{Err: "Server error", Status: http.StatusInternalServerError})
		return
	}

	// send response as json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(balance)
}

// UpdateUser is updating the balance of the user
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	// get username from the path
	username := getUsername(r.URL.Path)

	// read body of the request
	bs := make([]byte, r.ContentLength)
	r.Body.Read(bs)
	var body model.UpdateBalanceBody
	if err := json.Unmarshal(bs, &body); err != nil {
		sendError(&w, ErrorResponse{Err: "Invalid body", Status: http.StatusBadRequest})
		return
	}

	balance, err := h.service.UpdateBalance(username, body)

	// any problem on updating balance (with specific error message)
	if err != nil {
		sendError(&w, ErrorResponse{Err: err.Error(), Status: http.StatusInternalServerError})
		return
	}

	// send response as json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(balance)
}

// NewHandler is returning new handler
func NewHandler(s service.IService) IHandler {
	return &Handler{service: s}
}

// sendError is sending error as a response
func sendError(w *http.ResponseWriter, err ErrorResponse) {
	(*w).WriteHeader(err.Status)
	(*w).Header().Set("Content-Type", "application/json")
	json.NewEncoder(*w).Encode(err)
}

// parse username from the path
func getUsername(path string) string {
	parameters := strings.Split(path, "/")
	return parameters[1]
}
