package handlers

import (
	"github.com/dreamnajababy/go-ecom/src/models"
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
		return c.JSON(MakeHttpResponse("error", "please insert keyword before search.", 404, nil))
		//		return fiber.NewError(404, "please insert keyword before search.")
	}
	products, err := (*p.repository).SearchProduct(keyword)
	if err != nil {
		return c.JSON(MakeHttpResponse("success", err.Error(), 200, products))
	}
	return c.JSON(products)
}
func MakeHttpResponse(status string, description string, code int, data interface{}) models.HttpResponse {
	return models.HttpResponse{Status: status, Description: description, Code: code, Data: data}
}
