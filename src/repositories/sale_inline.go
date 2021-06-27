package repositories

import (
	models "github.com/dreamnajababy/go-ecom/src/models"
)

const (
	ErrSaleNotFound = SaleErr("sale not found")
)

type SaleInlineRepository struct {
	Sales []models.Sale
}

func (s *SaleInlineRepository) InitSale() {
	s.Sales = make([]models.Sale, 0)
}

var counter = 1

func (s *SaleInlineRepository) getSaleFromProductID(pid int) (*models.Sale, error) {
	for idx, sale := range s.Sales {
		if sale.Pid == pid {
			return &s.Sales[idx], nil
		}
	}
	return &models.Sale{}, ErrSaleNotFound
}

func (s *SaleInlineRepository) addNewSaleFromProduct(id int, p models.Product) error {
	s.Sales = append(s.Sales, models.Sale{
		Id:       id,
		Pid:      p.Id,
		Price:    p.Price,
		Quantity: 1,
		Sum:      p.Price,
	})
	return nil
}

func (s *SaleInlineRepository) StoreSale(productsOrder []models.Product) error {
	for _, prod := range productsOrder {
		sale, err := s.getSaleFromProductID(prod.Id)

		if err == ErrSaleNotFound {
			s.addNewSaleFromProduct(counter, prod)
			continue
		}

		sale.Update(1, prod.Price)
	}
	return nil
}

func (s SaleInlineRepository) GetSale() ([]models.Sale, error) {
	return s.Sales, nil
}

func (s SaleInlineRepository) GetReadyMutateSale() (*[]models.Sale, error) {
	return &s.Sales, nil
}

type SaleErr string

func (e SaleErr) Error() string {
	return string(e)
}
