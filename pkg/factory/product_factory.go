package factory

import (
	"product-service-go/pkg/domain/aggregates"
	"product-service-go/pkg/domain/valueobjects"
)

type ProductFactory struct{}

func NewProductFactory() ProductFactory {
	return ProductFactory{}
}

func (f *ProductFactory) CreateProduct(category valueobjects.Category, image valueobjects.Image) *aggregates.Product {
	// Aquí podrías generar un ID único para el producto.
	return &aggregates.Product{ID: "unique-id", Category: category, Image: image}
}

func (f *ProductFactory) CreateCategory(name string) valueobjects.Category {
	return valueobjects.Category{Name: name}
}

func (f *ProductFactory) CreateImage(url string) valueobjects.Image {
	return valueobjects.Image{URL: url}
}
