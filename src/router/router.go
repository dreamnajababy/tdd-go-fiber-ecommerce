package router

import (
	handler "github.com/dreamnajababy/go-ecom/src/handlers"
	repo "github.com/dreamnajababy/go-ecom/src/repositories"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	App *fiber.App
}

func (r *Router) SetApp(app *fiber.App) {
	r.App = app
}
func (r *Router) SetProductRoutes(repository repo.ProductRepository) {
	productHandler := &handler.ProductHandler{}
	productHandler.InitHandler(&repository)
	r.App.Get("/product/:id", productHandler.GetProductByID)
	r.App.Get("/products", productHandler.GetProducts)
	r.App.Get("/products/search", productHandler.SearchProduct)
}

func (r *Router) SetSaleRoutes(saleRepository repo.SaleRepository, receiptRepository repo.ReceiptRepository) {
	saleHandler := &handler.SaleHandler{}
	saleHandler.InitHandler(&saleRepository, &receiptRepository)
	r.App.Post("/sales", saleHandler.StoreSale)
}
func (r *Router) SetLoginRoutes() {
	authHandler := &handler.AuthHandler{}
	repo := &repo.UserInlineRepository{}
	//fmt.Printf("Type:%T &repo=%p repo=%v *repo=%v\n", repo, &repo, repo, *repo)
	authHandler.InitHandler(repo)
	r.App.Post("/login", authHandler.Login)
}
