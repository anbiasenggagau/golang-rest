package services

import (
	customerController "REST-API/Controller/Customer"
	model "REST-API/Model"
	repository "REST-API/Repository"
	"errors"
)

type CustomerService struct {
	Repository repository.CustomerRepository
}

func (service CustomerService) GetACustomer(id string) (*model.Customer, error) {
	customer := service.Repository.FindById(id)
	if customer == nil {
		return nil, errors.New("Not Found")
	} else {
		return customer, nil
	}
}

func (service CustomerService) GetAllCustomer() ([]model.Customer, error) {
	customers := service.Repository.GetAll()
	if customers == nil {
		return nil, errors.New("Not Found")
	} else {
		return customers, nil
	}
}

func (service CustomerService) CreateACustomer(request customerController.Request) (*model.Customer, error) {
	customer := service.Repository.Create(request)
	if customer == nil {
		return nil, errors.New("Bad Request")
	} else {
		return customer, nil
	}
}
