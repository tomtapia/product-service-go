package application

import (
	"product-service-go/pkg/domain/aggregates"
	"product-service-go/pkg/factory"
	"product-service-go/pkg/repository"
)

type CreateProductUseCase struct {
	repo    repository.ProductRepository
	factory factory.ProductFactory
}

func NewCreateProductUseCase(repo repository.ProductRepository, factory factory.ProductFactory) CreateProductUseCase {
	return CreateProductUseCase{repo: repo, factory: factory}
}

func (u *CreateProductUseCase) CreateProduct(categoryName string, imageURL string) (*aggregates.Product, error) {
	category := u.factory.CreateCategory(categoryName)
	image := u.factory.CreateImage(imageURL)
	product := u.factory.CreateProduct(category, image)
	err := u.repo.Save(product)
	return product, err
}
