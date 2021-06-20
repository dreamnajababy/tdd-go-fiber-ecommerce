package main

import (
	repo "github.com/dreamnajababy/go-ecom/src/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {

}

func SetupProductTest() *fiber.App { // injector
	app := fiber.New()
	repository := &repo.ProductInlineRepository{}
	repository.InitProduct()

	r := Router{app}
	r.SetProductRoutes(repository)
	return app
}
