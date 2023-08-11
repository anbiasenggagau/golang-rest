package services

import (
	customerAddressController "REST-API/Controller/CustomerAddress"
	model "REST-API/Model"
	repository "REST-API/Repository"
	"errors"
)

type CustomerAddressService struct {
	Repository repository.CustomerAddressRepository
}

func (service CustomerAddressService) GetACustomerAddress(id string) (*model.CustomerAddress, error) {
	customerAddress := service.Repository.FindById(id)
	if customerAddress == nil {
		return nil, errors.New("Not Found")
	} else {
		return customerAddress, nil
	}
}

func (service CustomerAddressService) GetAllCustomerAddress() ([]model.CustomerAddress, error) {
	customerAddress := service.Repository.GetAll()
	if customerAddress == nil {
		return nil, errors.New("Not Found")
	} else {
		return customerAddress, nil
	}
}

func (service CustomerAddressService) CreateACustomerAddress(request customerAddressController.Request) (*model.CustomerAddress, error) {
	customerAddress := service.Repository.Create(request)
	if customerAddress == nil {
		return nil, errors.New("Bad Request")
	} else {
		return customerAddress, nil
	}
}
