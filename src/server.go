package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {

}

var Products = []Product{{Id: 1}, {Id: 2}, {Id: 3}, {Id: 4}, {Id: 5}}

func Setup() *fiber.App {
	app := fiber.New()
	app.Get("/products", func(c *fiber.Ctx) error {
		return c.JSON(Products)
	})
	app.Get("/products/:id", func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		prod := SearchProdByID(id)
		return c.JSON(prod)
	})
	return app
}

func SearchProdByID(id int) Product {
	var result Product
	if id == 0 {
		fiber.NewError(404, "404 Product Not Found")
	}
	for _, val := range Products {
		if val.Id == id {
			result = val
			break
		}
	}
	return result
}
