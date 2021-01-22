package main

import (
	delivery "Clean-Architecture/Delivery/customer"
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"Clean-Architecture/Service"
	"Clean-Architecture/Store/customer"
	_ "github.com/go-sql-driver/mysql"
)

func main()  {
	var db, err = sql.Open("mysql", "root:saima@123Sult@/CustomerDB")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	r := mux.NewRouter()
	datastore := store.New(db)
	service := service.New(datastore)
	handler := delivery.New(service)

	r.HandleFunc("/customer", handler.GetCustomerByName).Methods(http.MethodGet)
	r.HandleFunc("/customer/{id}", handler.GetCustomerById).Methods(http.MethodGet)
	r.HandleFunc("/customer", handler.CreateCustomer).Methods(http.MethodPost)
	r.HandleFunc("/customer/{id}", handler.EditCustomerDetails).Methods(http.MethodPut)
	r.HandleFunc("/customer/{id}", handler.DeleteCustomerById).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":8080", r))
}
