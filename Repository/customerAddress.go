package repository

import (
	customerAddressController "REST-API/Controller/CustomerAddress"
	model "REST-API/Model"
)

type CustomerAddressRepository interface {
	FindById(id string) *model.CustomerAddress
	GetAll() []model.CustomerAddress
	Create(request customerAddressController.Request) *model.CustomerAddress
}
