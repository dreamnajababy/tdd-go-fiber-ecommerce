package models

type Sale struct {
	Id       int
	Pid      int
	Price    float64
	Quantity int
	Sum      float64
	Rid      int
}

func (s *Sale) Update(quantity int, price float64) error {
	s.Quantity += quantity
	s.Sum += price
	return nil
}
func (s *Sale) UpdateRID(rid int) error {
	s.Rid = rid
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
