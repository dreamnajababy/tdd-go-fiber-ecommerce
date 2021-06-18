package main

import (
	"encoding/json"
	"net/http/httptest"
	"reflect"
	"testing"

	models "github.com/dreamnajababy/go-ecom/src/models"
)

func TestProduct(t *testing.T) {
	app := Setup()
	t.Run("get product with id and get that product returns", func(t *testing.T) {
		var got models.Product
		want := models.Product{Id: 5}
		request := httptest.NewRequest("GET", "/products/5", nil)

		resp, _ := app.Test(request) //resp.Body return io.Reader
		err := json.NewDecoder(resp.Body).Decode(&got)

		assertStatusCode(t, 200, resp.StatusCode)
		assertStruct(t, want, got, err)
	})
	t.Run("get products when enter link products", func(t *testing.T) {
		var got []models.Product
		want := []models.Product{{Id: 1}, {Id: 2}, {Id: 3}, {Id: 4}, {Id: 5}}
		request := httptest.NewRequest("GET", "/products", nil)

		resp, _ := app.Test(request) //resp.Body return io.Reader
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
