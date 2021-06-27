package main

import (
	handler "github.com/dreamnajababy/go-ecom/src/handlers"
	repo "github.com/dreamnajababy/go-ecom/src/repositories"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	app *fiber.App
}

func (r *Router) SetApp(app *fiber.App) {
	r.app = app
}
func (r *Router) SetProductRoutes(repository repo.ProductRepository) {
	productHandler := &handler.ProductHandler{}
	productHandler.InitHandler(&repository)

	r.app.Get("/product/:id", productHandler.GetProductByID)
	r.app.Get("/products", productHandler.GetProducts)
	r.app.Get("/products/search", productHandler.SearchProduct)
}

func (r *Router) SetSaleRoutes(saleRepository repo.SaleRepository, receiptRepository repo.ReceiptRepository) {
	saleHandler := &handler.SaleHandler{}
	saleHandler.InitHandler(&saleRepository, &receiptRepository)
	r.app.Post("/sales", saleHandler.StoreSale)
}
