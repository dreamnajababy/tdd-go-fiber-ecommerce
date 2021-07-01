package handlers

import (
	models "github.com/dreamnajababy/go-ecom/src/models"
	repo "github.com/dreamnajababy/go-ecom/src/repositories"
	"github.com/gofiber/fiber/v2"
)

type SaleHandler struct {
	saleRepository    *repo.SaleRepository
	receiptRepository *repo.ReceiptRepository
}

func (s *SaleHandler) InitHandler(saleRepository *repo.SaleRepository, receiptRepository *repo.ReceiptRepository) {
	s.saleRepository = saleRepository
	s.receiptRepository = receiptRepository
}

func (s *SaleHandler) StoreSale(c *fiber.Ctx) error {
	var productOrders []models.Product
	err := c.BodyParser(&productOrders)

	if err != nil {
		return fiber.NewError(500, "cannot parse json to products struct.")
	}

	err = (*s.saleRepository).StoreSale(productOrders)

	if err != nil {
		return fiber.NewError(500, "cannot store orders.")
	}

	sale, err := (*s.saleRepository).GetMutateSale()

	if err != nil {
		return fiber.NewError(500, "cannot get sale.")
	}

	err = (*s.receiptRepository).CreateReceiptFromSale(sale)

	if err != nil {
		return fiber.NewError(500, "cannot create receipt from sale.")
	}

	response := models.Response{
		StatusCode: 201,
		Msg:        "created sale and receipt successfully.",
	}

	return c.Status(201).JSON(response)
}
