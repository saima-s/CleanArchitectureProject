package store

import "C"
import (
	_ "Clean-Architecture/entities"
	customer "Clean-Architecture/entities"
	"database/sql"
	"fmt"
	"log"
)

type CustomerStorer struct {
	db *sql.DB
}

func New(db *sql.DB) CustomerStorer {
	return CustomerStorer{db: db}
}
func (c CustomerStorer) GetCustomerByName(name string) ([]customer.Customer) {
	query := "SELECT * FROM Customers INNER JOIN Address ON Customers.ID = Address.CustId;"
	if name ==""{
		fmt.Println("empty name")
	}
	var data []interface{}
	if name !=""{
		query = "SELECT * FROM Customers INNER JOIN Address ON Customers.ID = Address.CustId where Customers.Name = ?;"
		data = append(data, name)
	}
	 rows, err := c.db.Query(query, data...)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var result []customer.Customer
	for rows.Next() {
		var cust customer.Customer
		if err := rows.Scan(&cust.ID, &cust.Name, &cust.DOB, &cust.Address.ID, &cust.Address.City, &cust.Address.State, &cust.Address.StreetName, &cust.Address.CustId); err != nil {
			log.Fatal(err)
		}
		result = append(result, cust)
	}
	return result
}
func (c CustomerStorer) GetCustomerById(id int) (customer.Customer) {
	var ids []interface{}
	ids = append(ids, id)
	query := `SELECT * FROM Customers INNER JOIN Address ON Customers.ID = Address.CustId where Customers.ID = ?; `
	rows, err := c.db.Query(query, ids...)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var cust customer.Customer
	for rows.Next() {
		if err := rows.Scan(&cust.ID, &cust.Name, &cust.DOB, &cust.Address.ID, &cust.Address.City, &cust.Address.State, &cust.Address.StreetName, &cust.Address.CustId); err != nil {
			log.Fatal(err)
		}
	}
	return cust
}
func (c CustomerStorer) CreateCustomer(customer1 customer.Customer) (customer.Customer) {
	var customers []interface{}
	customers = append(customers, customer1.Name)
	customers = append(customers, customer1.DOB)
	query := `INSERT INTO Customers(name, DOB) VALUES(?,?);`
	rows, err := c.db.Exec(query, customers...)
	if err != nil {
		panic(err.Error())
	}
	id, _ := rows.LastInsertId()
	var addr []interface{}
	addr = append(addr, customer1.Address.City)
	addr = append(addr, customer1.Address.State)
	addr = append(addr, customer1.Address.StreetName)
	addr = append(addr, id)
	query1 := `INSERT INTO Address(City,State,StreetName,CustId) VALUES(?,?,?,?)`
	row, err1 := c.db.Exec(query1, addr...)
	if err1 != nil {
		panic(err.Error())
	}
	idAddr, _ := row.LastInsertId()
	customer1.ID = int(id)
	customer1.Address.ID = int(idAddr)
	customer1.Address.CustId = int(id)
	return customer1
}
func (c CustomerStorer) EditCustomerDetails(id int, customer2 customer.Customer) (customer.Customer) {
	query1 := `SELECT * from Customers where ID =?`
	var id1 []interface{}
	id1 = append(id1, id)
	row, err := c.db.Query(query1, id1...)
	if err != nil {
		panic(err.Error())
	}
	if !row.Next() {
		return customer.Customer{}

	} else {
		if customer2.Name != "" {
			_, err := c.db.Exec("update Customers set Name=? where ID=?", customer2.Name, id)
			if err != nil {
				panic(err.Error())
				return customer.Customer{}
			}
			var custId []interface{}
			custId = append(custId, id)
			q := `SELECT * FROM Customers INNER JOIN Address where Address.CustID =?`
			r, _ := c.db.Query(q, custId...)
			var cu customer.Customer
			for r.Next() {
				e := r.Scan(&cu.ID, &cu.Name, &cu.DOB, &cu.Address.ID, &cu.Address.City, &cu.Address.State, &cu.Address.StreetName, &cu.Address.CustId)
				if e != nil {
					log.Fatal(e)
				}
			}
			var data []interface{}
			query := "update Address set "
			if customer2.Address.City != "" {
				query += "City = ? ,"
				data = append(data, customer2.Address.City)
			}
			if customer2.Address.State != "" {
				query += "State = ? ,"
				data = append(data, customer2.Address.State)
			}
			if customer2.Address.StreetName != "" {
				query += "StreetName = ? ,"
				data = append(data, customer2.Address.StreetName)
			}

			query = query[:len(query)-1]
			query += "where CustId = ? and ID = ?"
			data = append(data, id)
			data = append(data, cu.Address.ID)
			_, err = c.db.Exec(query, data...)

			if err != nil {
				log.Fatal(err)
			}
			return customer2
		}
		return customer.Customer{}
	}
}
func (c CustomerStorer) DeleteCustomerById(id int) (customer.Customer) {
	var ids []interface{}
	ids = append(ids, id)
	query := `SELECT * FROM Customers INNER JOIN Address ON Customers.ID = Address.CustId where Customers.ID = ?; `
	rows, err := c.db.Query(query, ids...)
	if err != nil {
		panic(err.Error())
	}
	if !rows.Next() {
		return customer.Customer{}
	} else {
		query = `DELETE  FROM Customers where ID =?; `
		_, err1 := c.db.Exec(query, ids...)
		if err1 != nil {
			panic(err.Error())
		}
		defer rows.Close()
		var cust customer.Customer
		for rows.Next() {
			if err := rows.Scan(&cust.ID, &cust.Name, &cust.DOB, &cust.Address.ID, &cust.Address.City, &cust.Address.State, &cust.Address.StreetName, &cust.Address.CustId); err != nil {
				log.Fatal(err)
			}
		}
		return cust
	}
}
