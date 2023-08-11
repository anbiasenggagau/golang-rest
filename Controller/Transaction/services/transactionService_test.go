package services

import (
	transactionController "REST-API/Controller/Transaction"
	model "REST-API/Model"
	repository "REST-API/Repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var transactionRepository = &repository.TransactionRepositoryMock{Mock: mock.Mock{}}
var transactionService = TransactionService{Repository: transactionRepository}

func TestTransaction_Found(t *testing.T) {
	var transaction = model.Transaction{
		Id:            1,
		CustomerName:  "Anbia",
		ProductName:   "Sabun,Shampoo",
		PaymentMethod: "Cash",
		TotalPrice:    22000,
		CreatedAt:     time.Now(),
	}

	transactionRepository.Mock.On("FindById", "1").Return(transaction)

	result, err := transactionService.GetATransaction("1")

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, transaction.Id, result.Id)
	assert.Equal(t, transaction.CustomerName, result.CustomerName)
	assert.Equal(t, transaction.ProductName, result.ProductName)
	assert.Equal(t, transaction.PaymentMethod, result.PaymentMethod)
	assert.Equal(t, transaction.TotalPrice, result.TotalPrice)
	assert.Equal(t, transaction.CreatedAt, result.CreatedAt)
}

func TestTransaction_NotFound(t *testing.T) {

	transactionRepository.Mock.On("FindById", "0").Return(nil)

	result, err := transactionService.GetATransaction("0")

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func TestTransaction_ErrorInput(t *testing.T) {
	transactionRepository.Mock.On("FindById", "cd").Return(nil)

	result, err := transactionService.GetATransaction("cd")
	assert.NotNil(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "Not Found", err.Error())
}

func TestTransaction_GetAll(t *testing.T) {

	var transactions = []model.Transaction{
		{
			Id:            1,
			CustomerName:  "Anbia",
			ProductName:   "Sabun,Shampoo",
			PaymentMethod: "Cash",
			TotalPrice:    22000,
			CreatedAt:     time.Now(),
		},
		{
			Id:            2,
			CustomerName:  "Senggagau",
			ProductName:   "Sabun",
			PaymentMethod: "Cash",
			TotalPrice:    10000,
			CreatedAt:     time.Now(),
		},
	}

	transactionRepository.Mock.On("GetAll").Return(transactions)

	result, err := transactionService.GetAllTransaction()

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, len(transactions), len(result))
}

func TestTransaction_CreateNew(t *testing.T) {
	var request = transactionController.Request{
		CustomerName:  "Senggagau",
		ProductName:   "Sabun,Shampoo",
		PaymentMethod: "Cash",
	}

	var response = model.Transaction{
		Id:            3,
		CustomerName:  "Senggagau",
		ProductName:   "Sabun,Shampoo",
		PaymentMethod: "Cash",
		TotalPrice:    22000,
		CreatedAt:     time.Now(),
	}
	transactionRepository.Mock.On("Create", request).Return(response)

	result, err := transactionService.CreateATransaction(request)
	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, response.Id, result.Id)
	assert.Equal(t, response.CustomerName, result.CustomerName)
	assert.Equal(t, response.ProductName, result.ProductName)
	assert.Equal(t, response.PaymentMethod, result.PaymentMethod)
	assert.Equal(t, response.TotalPrice, result.TotalPrice)
	assert.Equal(t, response.CreatedAt, result.CreatedAt)
}

func TestTransaction_NoProduct(t *testing.T) {
	var request = transactionController.Request{
		CustomerName:  "Senggagau",
		PaymentMethod: "Cash",
	}

	transactionRepository.Mock.On("Create", request).Return(nil)

	result, err := transactionService.CreateATransaction(request)
	assert.NotNil(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "Bad Request", err.Error())
}

func TestTransaction_InactivePayment(t *testing.T) {
	var request = transactionController.Request{
		CustomerName:  "Senggagau",
		ProductName:   "Sabun",
		PaymentMethod: "Online",
	}

	transactionRepository.Mock.On("Create", request).Return(nil)

	result, err := transactionService.CreateATransaction(request)
	assert.NotNil(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "Bad Request", err.Error())
}
