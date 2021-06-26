package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	models "github.com/dreamnajababy/go-ecom/src/models"
	repo "github.com/dreamnajababy/go-ecom/src/repositories"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	StatusCode int
	Msg        string
}

func TestSaleUnit(t *testing.T) {
	repository := &repo.SaleInlineRepository{}
	app := SetupSaleTest(fiber.New(), repository)

	tests := []struct {
		name               string
		url                string
		payload            []models.Product
		expectedHTTPstatus int
		expectedMsg        string
		expectedSale       []models.Sale
	}{
		{
			"create a sale from a product and insert to DB",
			"/sales", models.ProductOrder, 201, "created sales successfully.",
			createExpectedSales(1, 1, 1, 100),
		},
		{
			"create sale from products and get response 201 with msg with sale in repository",
			"/sales", models.ProductsOrder, 201, "created sales successfully.",
			createExpectedSales(1, 1, 5, 100),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repository.InitSale()

			var got Response
			want := createResponse(tc.expectedHTTPstatus, tc.expectedMsg)

			bytesData, _ := json.Marshal(tc.payload)
			payload := bytes.NewReader(bytesData)

			request := httptest.NewRequest("POST", "/sales", payload)
			request.Header.Set("Content-Type", "application/json") // need to set header for using json body parser
			resp, _ := app.Test(request)

			err := json.NewDecoder(resp.Body).Decode(&got)

			gotSale, _ := repository.GetSale()
			assertStatusCode(t, tc.expectedHTTPstatus, resp.StatusCode) // assert HTTP Response Status Code
			assertStruct(t, want, got, err)                             // assert HTTP Response
			assertStruct(t, tc.expectedSale, gotSale, err)              // assert Mock DB
		})
	}

}

func createResponse(statusCode int, Msg string) Response {
	return Response{
		StatusCode: statusCode,
		Msg:        Msg,
	}
}
func createExpectedSales(id, pid, quantity int, price float64) []models.Sale {
	result := []models.Sale{
		{
			Id:       id,
			Pid:      pid,
			Quantity: quantity,
			Price:    price,
			Sum:      float64(quantity) * price,
		},
	}
	return result
}
