package repository

import (
	customerAddressController "REST-API/Controller/CustomerAddress"
	model "REST-API/Model"

	"github.com/stretchr/testify/mock"
)

type CustomerAddressRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CustomerAddressRepositoryMock) FindById(id string) *model.CustomerAddress {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	} else {
		customerAddress := arguments.Get(0).(model.CustomerAddress)
		return &customerAddress
	}
}

func (repository *CustomerAddressRepositoryMock) GetAll() []model.CustomerAddress {
	arguments := repository.Mock.Called()
	if arguments.Get(0) == nil {
		return nil
	} else {
		customerAddress := arguments.Get(0).([]model.CustomerAddress)
		return customerAddress
	}
}

func (repository *CustomerAddressRepositoryMock) Create(request customerAddressController.Request) *model.CustomerAddress {
	arguments := repository.Mock.Called(request)
	if arguments.Get(0) == nil {
		return nil
	} else {
		customerAddress := arguments.Get(0).(model.CustomerAddress)
		return &customerAddress
	}
}
