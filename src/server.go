package main

import (
	repo "github.com/dreamnajababy/go-ecom/src/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {

}

func SetupProductTest() *fiber.App { // injector
	app := fiber.New()
	r := Router{app}
	repository := (&repo.ProductInlineRepository{}).InitProduct()
	r.SetProductRoutes(repository)
	return app
}
