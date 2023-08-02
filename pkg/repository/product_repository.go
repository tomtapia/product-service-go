package repository

import "product-service-go/pkg/domain/aggregates"

type ProductRepository interface {
	Save(product *aggregates.Product) error
}
