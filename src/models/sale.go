package models

type Sale struct {
	Id       int
	Pid      int
	Price    float64
	Quantity int
	Sum      float64
}

func (s *Sale) Update(quantity int, price float64) error {
	s.Quantity += quantity
	s.Sum += price
	return nil
}

var ProductOrder = []Product{
	{Id: 1, Name: "prod 1", Price: 100},
}

var ProductsOrder = []Product{
	{Id: 1, Name: "prod 1", Price: 100},
	{Id: 1, Name: "prod 1", Price: 100},
	{Id: 1, Name: "prod 1", Price: 100},
	{Id: 1, Name: "prod 1", Price: 100},
	{Id: 1, Name: "prod 1", Price: 100},
}
