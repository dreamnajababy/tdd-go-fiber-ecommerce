package models

type Product struct {
	Id          int
	Name        string
	Price       float64
	Remain      int    `json:",omitempty"`
	Description string `json:",omitempty"`
}
