package repositories

import (
	"time"

	"github.com/dreamnajababy/go-ecom/src/models"
)

type ReceiptInlineRepository struct {
	Receipt models.Receipt
}

var receiptCounter = 1

func (r *ReceiptInlineRepository) GetReceipt() (models.Receipt, error) {
	return r.Receipt, nil
}

func (r *ReceiptInlineRepository) InitReceipt() {
	r.Receipt = models.Receipt{}
}
func (r *ReceiptInlineRepository) CreateReceiptFromSale(sales *[]models.Sale) error {

	var total float64 = 0

	present := time.Date(2021, 06, 30, 12, 0, 0, 0, time.UTC)

	for _, sale := range *sales {
		total += sale.Sum
	}

	r.Receipt = models.Receipt{
		Id:        receiptCounter,
		Total:     total,
		CreatedAt: present,
	}

	for i := 0; i < len(*sales); i++ {
		(*sales)[i].Rid = r.Receipt.Id
	}

	return nil
}
