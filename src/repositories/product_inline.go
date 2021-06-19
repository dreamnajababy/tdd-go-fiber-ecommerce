package repositories

import (
	"errors"

	models "github.com/dreamnajababy/go-ecom/src/models"
)

var Products = []models.Product{{Id: 1}, {Id: 2}, {Id: 3}, {Id: 4}, {Id: 5}}

var (
	errNotFound = errors.New("product not found")
)

type ProductInlineRepository struct {
	Products []models.Product
}

func (p *ProductInlineRepository) InitProduct() *ProductInlineRepository {
	p.Products = Products
	return p
}

func (p ProductInlineRepository) GetProducts() ([]models.Product, error) {
	return p.Products, nil
}

func (p ProductInlineRepository) GetProductByID(id int) (models.Product, error) {
	var result models.Product
	if id == 0 {
		return result, errNotFound
	}
	for _, val := range p.Products {
		if val.Id == id {
			result = val
			break
		}
	}
	return result, nil
}
