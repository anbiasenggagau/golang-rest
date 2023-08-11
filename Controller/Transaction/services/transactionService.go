package services

import (
	transactionController "REST-API/Controller/Transaction"
	model "REST-API/Model"
	repository "REST-API/Repository"
	"errors"
)

type TransactionService struct {
	Repository repository.TransactionRepository
}

func (service TransactionService) GetATransaction(id string) (*model.Transaction, error) {
	transaction := service.Repository.FindById(id)
	if transaction == nil {
		return nil, errors.New("Not Found")
	} else {
		return transaction, nil
	}
}

func (service TransactionService) GetAllTransaction() ([]model.Transaction, error) {
	transactions := service.Repository.GetAll()
	if transactions == nil {
		return nil, errors.New("Not Found")
	} else {
		return transactions, nil
	}
}

func (service TransactionService) CreateATransaction(request transactionController.Request) (*model.Transaction, error) {
	transaction := service.Repository.Create(request)
	if transaction == nil {
		return nil, errors.New("Bad Request")
	} else {
		return transaction, nil
	}
}
