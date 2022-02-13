package service_test

import (
	"bank/config"
	mock "bank/mock"
	"bank/model"
	"bank/service"
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
	mockRepository := mock.NewMockIRepository(mockController)

	mockRepository.
		EXPECT().
		GetUsers().
		Return(model.Users{}).
		Times(1)

	service := service.NewService(mockRepository, conf)
	users, err := service.GetUsers()

	assert.Equal(t, model.Users{}, users)
	assert.Nil(t, err)
}

func Test_GetUserBalance(t *testing.T) {
	t.Run("User exist case", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockRepository := mock.NewMockIRepository(mockController)
		mock_username := "username"
		mock_balance := float64(10)

		mockRepository.
			EXPECT().
			GetUserBalance(mock_username).
			Return(mock_balance, true).
			Times(1)

		service := service.NewService(mockRepository, conf)
		balance, err := service.GetUserBalance(mock_username)

		assert.Equal(t, mock_balance, balance)
		assert.Nil(t, err)
	})

	t.Run("User doesn`t exist case", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockRepository := mock.NewMockIRepository(mockController)
		username := "username"
		mock_balance := float64(0)

		mockRepository.
			EXPECT().
			GetUserBalance(username).
			Return(mock_balance, false).
			Times(1)

		service := service.NewService(mockRepository, conf)
		balance, err := service.GetUserBalance(username)

		assert.Equal(t, mock_balance, balance)
		assert.NotNil(t, err)
	})
}

func Test_CreateUser(t *testing.T) {
	mockController := gomock.NewController(t)
	mockRepository := mock.NewMockIRepository(mockController)
	mock_username := "username"
	mock_balance := float64(0)

	t.Run("Creates user successfully", func(t *testing.T) {
		mockRepository.
			EXPECT().
			GetUserBalance(mock_username).
			Return(mock_balance, false).
			Times(1)

		mockRepository.
			EXPECT().
			CreateUser(mock_username, conf.InitialBalance).
			Return(conf.InitialBalance, nil).
			Times(1)

		service := service.NewService(mockRepository, conf)
		balance, err := service.CreateUser(mock_username)

		assert.Equal(t, conf.InitialBalance, balance)
		assert.Nil(t, err)
	})

	t.Run("User already exists", func(t *testing.T) {
		mockRepository.
			EXPECT().
			GetUserBalance(mock_username).
			Return(mock_balance, true).
			Times(1)

		mockRepository.
			EXPECT().
			CreateUser(mock_username, conf.InitialBalance).
			Return(conf.InitialBalance, nil).
			Times(0)

		service := service.NewService(mockRepository, conf)
		balance, err := service.CreateUser(mock_username)

		assert.Equal(t, mock_balance, balance)
		assert.Nil(t, err)
	})
}

func Test_UpdateBalance(t *testing.T) {
	mockController := gomock.NewController(t)
	mockRepository := mock.NewMockIRepository(mockController)

	mock_username := "username"
	mock_updateData := model.UpdateBalanceBody{
		Balance: float64(-101),
	}
	mock_balance := float64(0.0)

	t.Run("Undefined user", func(t *testing.T) {
		mockRepository.
			EXPECT().
			GetUserBalance(mock_username).
			Return(mock_balance, false).
			Times(1)

		service := service.NewService(mockRepository, conf)
		balance, err := service.UpdateBalance(mock_username, mock_updateData)

		assert.Equal(t, mock_balance, balance)
		assert.NotNil(t, err)
	})

	t.Run("Invalid operation", func(t *testing.T) {
		mockRepository.
			EXPECT().
			GetUserBalance(mock_username).
			Return(mock_balance, true).
			Times(1)

		service := service.NewService(mockRepository, conf)
		balance, err := service.UpdateBalance(mock_username, mock_updateData)

		assert.Equal(t, mock_balance, balance)
		assert.NotNil(t, err)
	})

	t.Run("Successfully update balance", func(t *testing.T) {
		mock_balance = 100
		mockRepository.
			EXPECT().
			GetUserBalance(mock_username).
			Return(mock_balance, true).
			Times(1)

		mockRepository.
			EXPECT().
			UpdateBalance(mock_username, mock_balance+mock_updateData.Balance).
			Return(mock_balance+mock_updateData.Balance, nil).
			Times(1)

		service := service.NewService(mockRepository, conf)
		balance, err := service.UpdateBalance(mock_username, mock_updateData)

		assert.Equal(t, mock_balance+mock_updateData.Balance, balance)
		assert.Nil(t, err)
	})
}
