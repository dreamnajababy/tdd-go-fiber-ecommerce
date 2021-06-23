package repositories

import (
	"errors"

	models "github.com/dreamnajababy/go-ecom/src/models"
)

type SaleInlineRepository struct {
	Sales []models.Sale
}

func (s *SaleInlineRepository) InitSale() {
	s.Sales = make([]models.Sale, 0)
}

var counter = 1

/* search in []sales
if found
return sale
else
return newSale

*/
func (s *SaleInlineRepository) getSaleFromProductID(pid int) (*models.Sale, error) {
	for idx, sale := range s.Sales {
		if sale.Pid == pid {
			return &s.Sales[idx], nil
		}
	}
	return &models.Sale{}, errors.New("sale not found.")
}

func (s *SaleInlineRepository) StoreSale(productsOrder []models.Product) error {
	for _, prod := range productsOrder {
		sale, err := s.getSaleFromProductID(prod.Id)
		if err != nil {
			s.Sales = append(s.Sales, models.Sale{
				Id:       counter,
				Pid:      prod.Id,
				Price:    prod.Price,
				Quantity: 1,
				Sum:      prod.Price,
			})
			continue
		}
		sale.Quantity += 1
		sale.Sum += prod.Price
	}
	return nil
}

func (s SaleInlineRepository) GetSale() ([]models.Sale, error) {
	return s.Sales, nil
}
