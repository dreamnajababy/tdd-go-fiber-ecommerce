package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"time"

	models "github.com/dreamnajababy/go-ecom/src/models"
	repo "github.com/dreamnajababy/go-ecom/src/repositories"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	StatusCode int
	Msg        string
}

func TestSaleReceiptFeature(t *testing.T) {
	saleRepository := &repo.SaleInlineRepository{}
	receiptRepository := &repo.ReceiptInlineRepository{}
	app := SetupSaleTest(fiber.New(), saleRepository, receiptRepository)

	t.Run("create sale from products and create receipt from sale and get resp 201 with msg with sale in repo and receipt in repo", func(t *testing.T) {
		saleRepository.InitSale()

		var got Response
		want := createResponse(201, "created sale and receipt successfully.")

		bytesData, _ := json.Marshal(models.ProductsOrder)
		payload := bytes.NewReader(bytesData)

		request := httptest.NewRequest("POST", "/sales", payload)
		request.Header.Set("Content-Type", "application/json") // need to set header for using json body parser
		resp, _ := app.Test(request)

		err := json.NewDecoder(resp.Body).Decode(&got)

		gotSale, _ := saleRepository.GetSale()
		assertStatusCode(t, 201, resp.StatusCode)                                                                    // assert HTTP Response Status Code
		assertStruct(t, want, got, err)                                                                              // assert HTTP Response
		assertStruct(t, createExpectedSales(1, 1, 5, 1, 100), gotSale, err)                                          // assert Mock DB
		assertStruct(t, createExpectedReceipt(1, 500, time.Date(2021, 06, 30, 12, 0, 0, 0, time.UTC)), gotSale, err) // assert Mock DB
	})

}

func createExpectedReceipt(rid int, total float64, createdAt time.Time) models.Receipt {
	return models.Receipt{
		Id:        rid,
		Total:     total,
		CreatedAt: createdAt,
	}
}

func TestSaleUnit(t *testing.T) {
	saleRepository := &repo.SaleInlineRepository{}
	receiptRepository := &repo.ReceiptInlineRepository{}
	app := SetupSaleTest(fiber.New(), saleRepository, receiptRepository)

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
			"/sales", models.ProductOrder, 201, "created sale and receipt successfully.",
			createExpectedSales(1, 1, 1, 0, 100),
		},
		{
			"create sale from products and get response 201 with msg with sale in repository",
			"/sales", models.ProductsOrder, 201, "created sale and receipt successfully.",
			createExpectedSales(1, 1, 5, 0, 100),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			saleRepository.InitSale()

			var got Response
			want := createResponse(tc.expectedHTTPstatus, tc.expectedMsg)

			bytesData, _ := json.Marshal(tc.payload)
			payload := bytes.NewReader(bytesData)

			request := httptest.NewRequest("POST", "/sales", payload)
			request.Header.Set("Content-Type", "application/json") // need to set header for using json body parser
			resp, _ := app.Test(request)

			err := json.NewDecoder(resp.Body).Decode(&got)

			gotSale, _ := saleRepository.GetSale()
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
func createExpectedSales(id, pid, quantity, rid int, price float64) []models.Sale {
	result := []models.Sale{
		{
			Id:       id,
			Pid:      pid,
			Quantity: quantity,
			Price:    price,
			Sum:      float64(quantity) * price,
			Rid:      rid,
		},
	}
	return result
}
