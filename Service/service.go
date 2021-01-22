package service

import (
	"Clean-Architecture/Store"
	customer "Clean-Architecture/entities"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
	age "github.com/bearbin/go-age"
)
type CustomerService struct {
	store Store.Customers
}

func New(customer Store.Customers) CustomerService {
	return CustomerService{store: customer}
}
func getAge(year, month, day int) time.Time {
	dob := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return dob
}
func (c CustomerService)GetCustomerByName(name string) ([]customer.Customer,error){
	if len(name) == 0{
		fmt.Println("empty name in service::",name)
		resp := c.store.GetCustomerByName("")
		return resp,nil
	}else{
		resp := c.store.GetCustomerByName(name)
		return resp,nil
	}
}
func (c CustomerService)GetCustomerById(id int) (customer.Customer,error){
    if id == 0 {
		return customer.Customer{},errors.New("No data to update")
	}else{
		resp := c.store.GetCustomerById(id)
		return resp,nil
	}
}
func (c CustomerService)CreateCustomer(customer1 customer.Customer) (customer.Customer, error){
	var cust []interface{}
	cust = append(cust, customer1.Name)
	cust = append(cust, customer1.DOB)
	dob := customer1.DOB
	dob1 := strings.Split(dob, "/")
	y, _ := strconv.Atoi(dob1[2])
	m, _ := strconv.Atoi(dob1[1])
	d, _ := strconv.Atoi(dob1[0])
	getAge := getAge(y, m, d)
	if age.Age(getAge) >= 18 {
		resp := c.store.CreateCustomer(customer1)
		return resp,nil
	}else{
		return customer.Customer{},errors.New("cannot create customer")
	}

}
func (c CustomerService)EditCustomerDetails(id int ,customer1 customer.Customer) (customer.Customer, error){
	if id == 0 {
		return customer.Customer{},errors.New("No data to update")
	}else{
		resp := c.store.EditCustomerDetails(id,customer1)
		if reflect.ValueOf(resp) == reflect.ValueOf(customer.Customer{}){
			return resp, errors.New("No data to update")
		}
		return resp,nil
	}
}
func (c CustomerService)DeleteCustomerById(id int) (customer.Customer,error){
	if id == 0 {
		return customer.Customer{},errors.New("No data to delete")
	}else{
		resp := c.store.DeleteCustomerById(id)
		if reflect.ValueOf(resp) == reflect.ValueOf(customer.Customer{}){
			return resp, errors.New("No data to delete")
		}
		return resp,nil
	}
}
