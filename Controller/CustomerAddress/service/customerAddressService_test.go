package services

import (
	customerAddressController "REST-API/Controller/CustomerAddress"
	model "REST-API/Model"
	repository "REST-API/Repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var customerAddressRepository = &repository.CustomerAddressRepositoryMock{Mock: mock.Mock{}}
var customerAddressService = CustomerAddressService{Repository: customerAddressRepository}

func TestCustomerAddress_Found(t *testing.T) {
	var customerAddress = model.CustomerAddress{
		Id:         1,
		Address:    "Semarang",
		CustomerId: 1,
	}

	customerAddressRepository.Mock.On("FindById", "1").Return(customerAddress)

	result, err := customerAddressService.GetACustomerAddress("1")

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, customerAddress.Id, result.Id)
	assert.Equal(t, customerAddress.Address, result.Address)
}

func TestCustomerAddress_NotFound(t *testing.T) {

	customerAddressRepository.Mock.On("FindById", "0").Return(nil)

	result, err := customerAddressService.GetACustomerAddress("0")

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func TestCustomerAddress_ErrorInput(t *testing.T) {
	customerAddressRepository.Mock.On("FindById", "cd").Return(nil)

	result, err := customerAddressService.GetACustomerAddress("cd")
	assert.NotNil(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "Not Found", err.Error())
}

func TestCustomerAddress_GetAll(t *testing.T) {

	var customerAddress = []model.CustomerAddress{
		{
			Id:      1,
			Address: "Semarang",
		},
		{
			Id:      2,
			Address: "Jakarta",
		},
	}

	customerAddressRepository.Mock.On("GetAll").Return(customerAddress)

	result, err := customerAddressService.GetAllCustomerAddress()

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, len(customerAddress), len(result))
}

func TestCustomerAddress_CreateToCustomerIdThatAlreadyHas(t *testing.T) {
	var request = customerAddressController.Request{
		Address:    "Bandung",
		CustomerId: 1,
	}

	customerAddressRepository.Mock.On("Create", request).Return(nil)

	result, err := customerAddressService.CreateACustomerAddress(request)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "Bad Request", err.Error())
}

func TestCustomerAddress_CreateWithNonExistingCustomerId(t *testing.T) {
	var request = customerAddressController.Request{
		Address:    "Bandung",
		CustomerId: 99,
	}

	customerAddressRepository.Mock.On("Create", request).Return(nil)

	result, err := customerAddressService.CreateACustomerAddress(request)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "Bad Request", err.Error())
}

func TestCustomerAddress_CreateWithExistingCustomerId(t *testing.T) {
	var request = customerAddressController.Request{
		Address:    "Bandung",
		CustomerId: 3,
	}

	var response = model.CustomerAddress{
		Id:         3,
		Address:    "Bandung",
		CustomerId: 3,
	}

	customerAddressRepository.Mock.On("Create", request).Return(response)

	result, err := customerAddressService.CreateACustomerAddress(request)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}
