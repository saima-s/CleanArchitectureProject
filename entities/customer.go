package entities
type Customer struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	DOB     string  `json:"dob"`
	Address Address `json:"address"`
}
type Address struct {
	ID         int    `json:"id"`
	City       string `json:"city"`
	State      string `json:"state"`
	StreetName string `json:"streetName"`
	CustId     int    `json:"custId"`
}
