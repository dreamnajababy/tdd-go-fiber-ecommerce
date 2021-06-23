package repositories

import (
	models "github.com/dreamnajababy/go-ecom/src/models"
)

type SaleInlineRepository struct {
	Sales []models.Sale
}

func (s *SaleInlineRepository) InitSale() {

}

func (s *SaleInlineRepository) StoreSale(productOrder []models.Product) error {
	counter := 1
	for _, prod := range productOrder {
		s.Sales = append(s.Sales, models.Sale{
			Id:       counter,
			Pid:      prod.Id,
			Price:    prod.Price,
			Quantity: 1,
			Sum:      prod.Price * 1,
		})
	}
	//fmt.Printf("\n%#v", s.Sales)
	return nil
}

func (s SaleInlineRepository) GetSale() ([]models.Sale, error) {
	return s.Sales, nil
}
