package repositories

import models "github.com/dreamnajababy/go-ecom/src/models"

type ProductRepository interface {
	GetProducts() ([]models.Product, error)
	GetProductByID(id int) (models.Product, error)
	SearchProduct(keyword string) ([]models.Product, error)
}
