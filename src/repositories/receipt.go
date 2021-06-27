package repositories

import models "github.com/dreamnajababy/go-ecom/src/models"

type ReceiptRepository interface {
	CreateReceiptFromSale([]models.Sale) error
}
