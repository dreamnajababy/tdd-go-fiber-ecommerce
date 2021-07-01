package tests

import (
	repo "github.com/dreamnajababy/go-ecom/src/repositories"
	router "github.com/dreamnajababy/go-ecom/src/router"
	"github.com/gofiber/fiber/v2"
)

func SetupLoginTest() *fiber.App { // injector
	app := fiber.New()
	r := router.Router{app}
	r.SetLoginRoutes()
	return app
}
func SetupProductTest(app *fiber.App) *fiber.App { // injector
	repository := &repo.ProductInlineRepository{}
	repository.InitProduct()

	r := router.Router{app}
	r.SetProductRoutes(repository)
	return app
}

func SetupSaleTest(app *fiber.App, saleRepository repo.SaleRepository, receiptRepository repo.ReceiptRepository) *fiber.App { // injector
	r := router.Router{app}
	r.SetSaleRoutes(saleRepository, receiptRepository)
	return app
}
