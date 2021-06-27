package repositories

import models "github.com/dreamnajababy/go-ecom/src/models"

type SaleRepository interface {
	StoreSale([]models.Product) error
	GetSale() ([]models.Sale, error)
	GetReadyMutateSale() (*[]models.Sale, error)
}
