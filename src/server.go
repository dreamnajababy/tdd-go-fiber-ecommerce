package main

import (
	repo "github.com/dreamnajababy/go-ecom/src/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {

}

func SetupProductTest(app *fiber.App) *fiber.App { // injector
	repository := &repo.ProductInlineRepository{}
	repository.InitProduct()

	r := Router{app}
	r.SetProductRoutes(repository)
	return app
}

func SetupSaleTest(app *fiber.App) *fiber.App { // injector
	repository := &repo.SaleInlineRepository{}
	repository.InitSale()

	r := Router{app}
	r.SetSaleRoutes(repository)
	return app
}
