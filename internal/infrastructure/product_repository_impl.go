package infrastructure

import "product-service-go/pkg/domain/aggregates"

type ProductRepositoryImpl struct {
	// Aquí puedes agregar tu implementación de la base de datos.
}

func NewProductRepositoryImpl() *ProductRepositoryImpl {
	return &ProductRepositoryImpl{}
}

func (r *ProductRepositoryImpl) Save(product *aggregates.Product) error {
	// Aquí guardarías el producto en tu base de datos.
	return nil
}
