package delivery

import (
	service "Clean-Architecture/Service"
	_ "Clean-Architecture/Service"
	"Clean-Architecture/entities"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	service service.Customers
}
func New(serv service.Customers) CustomerHandler {
	return CustomerHandler{service: serv}
}
func (a CustomerHandler) GetCustomerByName(w http.ResponseWriter, r *http.Request) {
	name, ok := r.URL.Query()["name"]
		if !ok{
			resp, err := a.service.GetCustomerByName("")
			if err != nil {
				_, _ = w.Write([]byte("could not get customer information"))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			body, _ := json.Marshal(resp)
			_, _ = w.Write(body)
		}else {
			fmt.Println("empty name in service::", name)
			resp, err := a.service.GetCustomerByName(name[0])
			if err != nil {
				_, _ = w.Write([]byte("could not get customer information"))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			body, _ := json.Marshal(resp)
			_, _ = w.Write(body)
		}

}
func (a CustomerHandler) GetCustomerById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id, err := strconv.Atoi(param["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, err := a.service.GetCustomerById(id)
	if err != nil {
		_, _ = w.Write([]byte("could not retrieve customer"))
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
	body, _ := json.Marshal(resp)
	_, _ = w.Write(body)
}
func (a CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer entities.Customer
	body, _ := ioutil.ReadAll(r.Body)
	if len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}else {

		err := json.Unmarshal(body, &customer)
		if err != nil {
			fmt.Println(err)
			_, _ = w.Write([]byte("invalid body"))
			w.WriteHeader(http.StatusBadRequest)

			return
		}
		resp, err := a.service.CreateCustomer(customer)
		if err != nil {
			_, _ = w.Write([]byte("could not create customer"))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		body, _ = json.Marshal(resp)
		_, _ = w.Write(body)
	}
}
func (a CustomerHandler) EditCustomerDetails(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id, err := strconv.Atoi(param["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var customer entities.Customer
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &customer)
	if err != nil {
		fmt.Println(err)
		_, _ = w.Write([]byte("invalid body"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, err := a.service.EditCustomerDetails(id, customer)
	if err != nil {
		_, _ = w.Write([]byte("age is less than 18"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	body, _ = json.Marshal(resp)
	_, _ = w.Write(body)
}
func (a CustomerHandler) DeleteCustomerById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id, err := strconv.Atoi(param["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, err := a.service.DeleteCustomerById(id)
	if err != nil {
		_, _ = w.Write([]byte("No data to delete customer"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}else {
		w.WriteHeader(http.StatusNoContent)
		body, _ := json.Marshal(resp)
		_, _ = w.Write(body)
	}
}

