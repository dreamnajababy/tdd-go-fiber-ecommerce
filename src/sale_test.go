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
	repository.InitSale()
	app := SetupSaleTest(fiber.New(), repository)

	t.Run("create sales from a product and insert to DB", func(t *testing.T) {
		var got Response
		want := createResponse(201, "created sales successfully.")
		wantSale := createExpectedSales(1, 1, 1, 100)

		bytesData, _ := json.Marshal(models.ProductOrder)
		reader := bytes.NewReader(bytesData)

		request := httptest.NewRequest("POST", "/sales", reader)
		request.Header.Set("Content-Type", "application/json") // need to set header for using json body parser
		resp, _ := app.Test(request)

		err := json.NewDecoder(resp.Body).Decode(&got)
		gotSale, _ := repository.GetSale()

		assertStatusCode(t, 200, resp.StatusCode)
		assertStruct(t, want, got, err)
		assertStruct(t, wantSale, gotSale, err) // assert Mock DB

	})
}

func createResponse(statusCode int, Msg string) Response {
	return Response{
		StatusCode: statusCode,
		Msg:        Msg,
	}
}
func createExpectedSales(id, pid, quantity int, price float64) []models.Sale {
	result := []models.Sale{
		models.Sale{
			Id:       id,
			Pid:      pid,
			Quantity: quantity,
			Price:    price,
			Sum:      float64(quantity) * price,
		},
	}
	return result
}
