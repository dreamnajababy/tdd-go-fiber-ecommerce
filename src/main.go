package main

import (
	repo "github.com/dreamnajababy/go-ecom/src/repositories"
	router "github.com/dreamnajababy/go-ecom/src/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	saleRepo := &repo.SaleInlineRepository{}
	productRepo := &repo.ProductInlineRepository{}
	receiptRepo := &repo.ReceiptInlineRepository{}

	r := router.Router{app}
	r.SetLoginRoutes()
	r.SetProductRoutes(productRepo)
	r.SetSaleRoutes(saleRepo, receiptRepo)

	app.Listen("localhost:3000")
}
