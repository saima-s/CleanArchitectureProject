package service

import customer "Clean-Architecture/entities"

type Customers interface {
	GetCustomerByName(name string) ([]customer.Customer,error)
	GetCustomerById(id int) (customer.Customer,error)
	CreateCustomer(c customer.Customer) (customer.Customer, error)
	EditCustomerDetails(id int, c customer.Customer) (customer.Customer, error)
	DeleteCustomerById(id int) (customer.Customer,error)
}