package aggregates

import (
	"product-service-go/pkg/domain/valueobjects"
)

type Product struct {
	ID       string
	Category valueobjects.Category
	Image    valueobjects.Image
}
