package main

import (
	handler "github.com/dreamnajababy/go-ecom/src/handlers"
	repo "github.com/dreamnajababy/go-ecom/src/repositories"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	app *fiber.App
}

func (r *Router) InitRoutes(app *fiber.App) {
	r.app = app
}
func (r *Router) SetProductRoutes(repository repo.ProductRepository) {
	productHandler := handler.ProductHandler{}
	productHandler.InitHandler(repository)

	r.app.Get("/products", productHandler.GetProducts)
	r.app.Get("/products/:id", productHandler.GetProductByID)
}
