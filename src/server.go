package main

import (
	repo "github.com/dreamnajababy/go-ecom/src/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {

}

func Setup() *fiber.App { // injector
	app := fiber.New()
	productRepo := repo.ProductInlineRepository{}
	productRepo.InitProduct()

	r := Router{}
	r.InitRoutes(app)
	r.SetProductRoutes(productRepo)
	return app
}
