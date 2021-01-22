package Store

import (
	customer "Clean-Architecture/entities"
)

type Customers interface {
	GetCustomerByName(name string) []customer.Customer
	GetCustomerById(id int) customer.Customer
	CreateCustomer(c customer.Customer) customer.Customer
	EditCustomerDetails(id int, c customer.Customer) customer.Customer
	DeleteCustomerById(id int) customer.Customer
}