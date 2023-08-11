package repository

import (
	transactionController "REST-API/Controller/Transaction"
	model "REST-API/Model"

	"github.com/stretchr/testify/mock"
)

type TransactionRepositoryMock struct {
	Mock mock.Mock
}

func (repository *TransactionRepositoryMock) FindById(id string) *model.Transaction {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	} else {
		transaction := arguments.Get(0).(model.Transaction)
		return &transaction
	}
}

func (repository *TransactionRepositoryMock) GetAll() []model.Transaction {
	arguments := repository.Mock.Called()
	if arguments.Get(0) == nil {
		return nil
	} else {
		transactions := arguments.Get(0).([]model.Transaction)
		return transactions
	}
}

func (repository *TransactionRepositoryMock) Create(request transactionController.Request) *model.Transaction {
	arguments := repository.Mock.Called(request)
	if arguments.Get(0) == nil {
		return nil
	} else {
		transactions := arguments.Get(0).(model.Transaction)
		return &transactions
	}
}
