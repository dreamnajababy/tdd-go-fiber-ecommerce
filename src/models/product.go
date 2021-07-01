package models

type Product struct {
	Id          int
	Name        string
	Price       float64
	Remain      int    `json:",omitempty"`
	Description string `json:",omitempty"`
}

var Products = []Product{
	{Id: 1, Name: "first", Price: 100.00, Remain: 10, Description: "first product"},
	{Id: 2, Name: "second", Price: 200.00, Remain: 20, Description: "second product"},
	{Id: 3, Name: "third", Price: 300.00, Remain: 30, Description: "third product"},
	{Id: 4, Name: "forth", Price: 400.00, Remain: 40, Description: "forth product"},
	{Id: 5, Name: "fifth", Price: 500.00, Remain: 50, Description: "last product"},
}
