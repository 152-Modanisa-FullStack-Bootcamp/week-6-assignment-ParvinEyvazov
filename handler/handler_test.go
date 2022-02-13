package handler_test

import (
	"bank/config"
	"bank/handler"
	"bank/mock"
	"bank/model"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var conf config.Config = config.Config{
	InitialBalance: 0,
	MinumumBalance: -100,
}

func Test_GetUsers(t *testing.T) {

	mockController := gomock.NewController(t)
	mockService := mock.NewMockIService(mockController)

	t.Run("Successfully get users", func(t *testing.T) {
		mockService.
			EXPECT().
			GetUsers().
			Return(model.Users{}, nil).
			Times(1)

		handler := handler.NewHandler(mockService)

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)

		handler.GetUsers(res, req)

		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	})

	t.Run("Error on get users", func(t *testing.T) {
		mockService.
			EXPECT().
			GetUsers().
			Return(model.Users{}, errors.New("any error")).
			Times(1)

		handler := handler.NewHandler(mockService)

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)

		handler.GetUsers(res, req)

		assert.Equal(t, http.StatusNotFound, res.Result().StatusCode)
	})
}

func Test_GetBalance(t *testing.T) {

	mockController := gomock.NewController(t)
	mockService := mock.NewMockIService(mockController)
	mock_username := "username"
	mock_balance := float64(0.0)
	target := "/" + mock_username

	t.Run("Successfully get balance", func(t *testing.T) {
		mockService.
			EXPECT().
			GetUserBalance(mock_username).
			Return(mock_balance, nil).
			Times(1)

		handler := handler.NewHandler(mockService)

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, target, http.NoBody)

		handler.GetBalance(res, req)

		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	})

	t.Run("Getting balance of undefined user", func(t *testing.T) {
		mockService.
			EXPECT().
			GetUserBalance(mock_username).
			Return(mock_balance, errors.New("undefined user")).
			Times(1)

		handler := handler.NewHandler(mockService)

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, target, http.NoBody)

		handler.GetBalance(res, req)

		assert.Equal(t, http.StatusNotFound, res.Result().StatusCode)
	})
}

func Test_CreateUser(t *testing.T) {
	mockController := gomock.NewController(t)
	mockService := mock.NewMockIService(mockController)
	mock_username := "username"
	mock_balance := float64(0.0)
	target := "/" + mock_username

	t.Run("Successfully create user", func(t *testing.T) {
		mockService.
			EXPECT().
			CreateUser(mock_username).
			Return(conf.InitialBalance, nil).
			Times(1)

		handler := handler.NewHandler(mockService)

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPut, target, http.NoBody)

		handler.CreateUser(res, req)

		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	})

	t.Run("Server error", func(t *testing.T) {
		mockService.
			EXPECT().
			CreateUser(mock_username).
			Return(mock_balance, errors.New("server error")).
			Times(1)

		handler := handler.NewHandler(mockService)

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPut, target, http.NoBody)

		handler.CreateUser(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Result().StatusCode)
	})
}

func Test_UpdateUser(t *testing.T) {

	mockController := gomock.NewController(t)
	mockService := mock.NewMockIService(mockController)
	mock_username := "username"
	mock_balance := float64(0.0)
	target := "/" + mock_username

	t.Run("Update user balance successfully", func(t *testing.T) {
		updateBalanceBody := model.UpdateBalanceBody{
			Balance: 100,
		}

		mockService.
			EXPECT().
			UpdateBalance(mock_username, updateBalanceBody).
			Return(mock_balance+updateBalanceBody.Balance, nil).
			Times(1)

		handler := handler.NewHandler(mockService)

		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(updateBalanceBody)

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, target, &buf)

		handler.UpdateUser(res, req)

		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	})

	t.Run("Wrong body format", func(t *testing.T) {
		handler := handler.NewHandler(mockService)

		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode("s")

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, target, &buf)

		handler.UpdateUser(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Result().StatusCode)
	})

	t.Run("Error on updating user balance", func(t *testing.T) {
		updateBalanceBody := model.UpdateBalanceBody{
			Balance: 100,
		}

		mockService.
			EXPECT().
			UpdateBalance(mock_username, updateBalanceBody).
			Return(0.0, errors.New("any error")).
			Times(1)

		handler := handler.NewHandler(mockService)

		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(updateBalanceBody)

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, target, &buf)

		handler.UpdateUser(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Result().StatusCode)
	})
}
