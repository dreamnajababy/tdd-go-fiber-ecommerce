package handlers

import (
	models "github.com/dreamnajababy/go-ecom/src/models"
	repo "github.com/dreamnajababy/go-ecom/src/repositories"
	"github.com/gofiber/fiber/v2"
)

type SaleHandler struct {
	repository *repo.SaleRepository
}

func (s *SaleHandler) InitHandler(repository *repo.SaleRepository) {
	s.repository = repository
}

func (s *SaleHandler) StoreSale(c *fiber.Ctx) error {
	var productOrders []models.Product
	err := c.BodyParser(&productOrders)

	if err != nil {
		return fiber.NewError(500, "cannot parse json to products struct.")
	}

	err = (*s.repository).StoreSale(productOrders)

	if err != nil {
		return fiber.NewError(500, "cannot store orders.")
	}

	want := models.Response{
		StatusCode: 201,
		Msg:        "created sales successfully.",
	}

	return c.Status(201).JSON(want)
}
