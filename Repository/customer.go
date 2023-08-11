package repository

import (
	customerController "REST-API/Controller/Customer"
	model "REST-API/Model"
)

type CustomerRepository interface {
	FindById(id string) *model.Customer
	GetAll() []model.Customer
	Create(request customerController.Request) *model.Customer
}
