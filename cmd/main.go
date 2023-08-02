package main

import (
	"log"
	"net/http"
	"product-service-go/internal/infrastructure"
	"product-service-go/pkg/api"
	"product-service-go/pkg/application"
	"product-service-go/pkg/factory"
)

func main() {
	repo := infrastructure.NewProductRepositoryImpl()
	factory := factory.NewProductFactory()
	useCase := application.NewCreateProductUseCase(repo, factory)
	controller := api.NewProductController(useCase)

	http.HandleFunc("/product", controller.CreateProduct)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
