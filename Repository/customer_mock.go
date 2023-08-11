package repository

import (
	customerController "REST-API/Controller/Customer"
	model "REST-API/Model"

	"github.com/stretchr/testify/mock"
)

type CustomerRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CustomerRepositoryMock) FindById(id string) *model.Customer {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	} else {
		customer := arguments.Get(0).(model.Customer)
		return &customer
	}
}

func (repository *CustomerRepositoryMock) GetAll() []model.Customer {
	arguments := repository.Mock.Called()
	if arguments.Get(0) == nil {
		return nil
	} else {
		customers := arguments.Get(0).([]model.Customer)
		return customers
	}
}

func (repository *CustomerRepositoryMock) Create(request customerController.Request) *model.Customer {
	arguments := repository.Mock.Called(request)
	if arguments.Get(0) == nil {
		return nil
	} else {
		customers := arguments.Get(0).(model.Customer)
		return &customers
	}
}
