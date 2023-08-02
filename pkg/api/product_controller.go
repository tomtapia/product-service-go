package api

import (
	"encoding/json"
	"net/http"
	"product-service-go/pkg/application"
)

type ProductController struct {
	useCase application.CreateProductUseCase
}

func NewProductController(useCase application.CreateProductUseCase) *ProductController {
	return &ProductController{useCase: useCase}
}

func (c *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "501 Not Implemented.", http.StatusNotImplemented)
		return
	}

	// Aquí se leerían los parámetros del request y se pasarían al caso de uso.
	product, err := c.useCase.CreateProduct("Electronics", "image-url.jpg")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(product)
}
