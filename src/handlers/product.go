package handlers

import (
	models "github.com/dreamnajababy/go-ecom/src/models"
	repo "github.com/dreamnajababy/go-ecom/src/repositories"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	repository *repo.ProductRepository
}

func (p *ProductHandler) InitHandler(repository *repo.ProductRepository) {
	p.repository = repository
}

func (p ProductHandler) GetProducts(c *fiber.Ctx) error {
	products, err := (*p.repository).GetProducts()
	if err != nil {
		panic(err)
	}
	return c.JSON(products)
}

func (p ProductHandler) GetProductByID(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	product, err := (*p.repository).GetProductByID(id)
	if err != nil {
		panic(err)
	}
	return c.JSON(product)
}

func (p ProductHandler) SearchProduct(c *fiber.Ctx) error {
	keyword := c.Query("keyword")
	if keyword == "" {
		return fiber.NewError(404, "please insert keyword before search.")
	}
	products, err := (*p.repository).SearchProduct(keyword)
	if err != nil {
		want := models.HttpResponse{Status: "success", Description: err.Error(), Code: 200, Data: products}
		return c.JSON(want)
	}
	return c.JSON(products)
}
