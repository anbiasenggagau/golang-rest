package services

import (
	customerController "REST-API/Controller/Customer"
	model "REST-API/Model"
	repository "REST-API/Repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var customerRepository = &repository.CustomerRepositoryMock{Mock: mock.Mock{}}
var customerService = CustomerService{Repository: customerRepository}

func TestCustomer_Found(t *testing.T) {
	var customer = model.Customer{
		Id:           1,
		CustomerName: "Anbia",
	}

	customerRepository.Mock.On("FindById", "1").Return(customer)

	result, err := customerService.GetACustomer("1")

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, customer.Id, result.Id)
	assert.Equal(t, customer.CustomerName, result.CustomerName)
}

func TestCustomer_NotFound(t *testing.T) {

	customerRepository.Mock.On("FindById", "0").Return(nil)

	result, err := customerService.GetACustomer("0")

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func TestCustomer_ErrorInput(t *testing.T) {
	customerRepository.Mock.On("FindById", "cd").Return(nil)

	result, err := customerService.GetACustomer("cd")
	assert.NotNil(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "Not Found", err.Error())
}

func TestCustomer_GetAll(t *testing.T) {

	var customers = []model.Customer{
		{
			Id:           1,
			CustomerName: "Anbia",
		},
		{
			Id:           2,
			CustomerName: "Senggagau",
		},
	}

	customerRepository.Mock.On("GetAll").Return(customers)

	result, err := customerService.GetAllCustomer()

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, len(customers), len(result))
}

func TestCustomer_CreateNew(t *testing.T) {
	var request = customerController.Request{
		CustomerName: "Nadia",
	}

	var response = model.Customer{
		Id:           3,
		CustomerName: "Nadia",
	}
	customerRepository.Mock.On("Create", request).Return(response)

	result, err := customerService.CreateACustomer(request)
	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, response.Id, result.Id)
	assert.Equal(t, response.CustomerName, result.CustomerName)
}

func TestCustomer_NoName(t *testing.T) {
	var request = customerController.Request{
		CustomerName: "",
	}

	customerRepository.Mock.On("Create", request).Return(nil)

	result, err := customerService.CreateACustomer(request)
	assert.NotNil(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "Bad Request", err.Error())
}
