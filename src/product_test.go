package main

import (
	"encoding/json"
	"net/http/httptest"
	"reflect"
	"testing"

	models "github.com/dreamnajababy/go-ecom/src/models"
	"github.com/gofiber/fiber/v2"
)

func TestProductUnit(t *testing.T) {
	app := SetupProductTest(fiber.New())

	t.Run("get products and return products as slice of json", func(t *testing.T) {
		var got []models.Product
		want := []models.Product{
			{Id: 1}, {Id: 2}, {Id: 3, Name: "Wonderland"}, {Id: 4}, {Id: 5, Name: "KY"},
		}

		request := httptest.NewRequest("GET", "/products", nil)
		resp, _ := app.Test(request)
		err := json.NewDecoder(resp.Body).Decode(&got)

		assertStatusCode(t, 200, resp.StatusCode)
		assertStruct(t, want, got, err)
	})

	t.Run("get product with id and return product as json", func(t *testing.T) {
		var got models.Product
		want := models.Product{Id: 2}

		request := httptest.NewRequest("GET", "/product/2", nil)
		resp, _ := app.Test(request)
		err := json.NewDecoder(resp.Body).Decode(&got)

		assertStatusCode(t, 200, resp.StatusCode)
		assertStruct(t, want, got, err)
	})

	t.Run("get product by search product with product name and return products as slice of json", func(t *testing.T) {
		var got []models.Product
		want := []models.Product{{Id: 3, Name: "Wonderland"}}

		request := httptest.NewRequest("GET", "/products/search?keyword=Wonderland", nil)
		resp, _ := app.Test(request)
		err := json.NewDecoder(resp.Body).Decode(&got)

		assertStatusCode(t, 200, resp.StatusCode)
		assertStruct(t, want, got, err)
	})
}

func assertStatusCode(t testing.TB, want, got int) {
	t.Helper()
	if got != want {
		t.Errorf("want %v,  but got %v", want, got)
	}
}

func assertStruct(t testing.TB, want, got interface{}, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Unable to parse response from server")
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Want: %v \nGot:%v\n", want, got)
	}
}
