package main

import (
	"encoding/json"
	"net/http"
	"testing"

	models "github.com/dreamnajababy/go-ecom/src/models"
	"github.com/gofiber/fiber/v2"
)

func TestProductIntegration(t *testing.T) {
	setup()
	t.Run("get product from product i", func(t *testing.T) {
		var got []models.Product
		want := []models.Product{
			{Id: 1}, {Id: 2}, {Id: 3, Name: "Wonderland"}, {Id: 4}, {Id: 5, Name: "KY"},
		}

		resp, err := http.Get("http://localhost:3000/products")
		handleError(t, err)
		err = json.NewDecoder(resp.Body).Decode(&got)

		assertStatusCode(t, 200, resp.StatusCode)
		assertStruct(t, want, got, err)
	})
}

func setup() {
	config := fiber.Config{DisableStartupMessage: true}
	newApp := fiber.New(config)
	app := SetupProductTest(newApp)
	go app.Listen("localhost:3000")
}

func handleError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Error(err)
	}
}
