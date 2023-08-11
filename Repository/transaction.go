package repository

import (
	transactionController "REST-API/Controller/Transaction"
	model "REST-API/Model"
)

type TransactionRepository interface {
	FindById(id string) *model.Transaction
	GetAll() []model.Transaction
	Create(request transactionController.Request) *model.Transaction
}
