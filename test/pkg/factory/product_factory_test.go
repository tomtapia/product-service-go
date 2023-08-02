package factory

import (
	"product-service-go/pkg/domain/aggregates"
	"product-service-go/pkg/domain/valueobjects"
	"product-service-go/pkg/factory"
	"reflect"
	"testing"
)

func TestNewProductFactory(t *testing.T) {
	tests := []struct {
		name string
		want factory.ProductFactory
	}{
		{
			name: "Test creation of new product factory",
			want: factory.NewProductFactory(),
		},
	}
	// Test case: Verificar que la creación de una instancia de Factory de productos sea consistente.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factory.NewProductFactory(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProductFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductFactory_CreateProduct(t *testing.T) {
	// Definir objetos de categoría e imagen para la prueba
	category := valueobjects.Category{Name: "Electronics"}
	image := valueobjects.Image{URL: "image.jpg"}

	type args struct {
		category valueobjects.Category
		image    valueobjects.Image
	}

	tests := []struct {
		name string
		f    *factory.ProductFactory
		args args
		want *aggregates.Product
	}{
		{
			name: "Test creation of new product with valid category and image",
			args: args{category: category, image: image},
			want: &aggregates.Product{ID: "unique-id", Category: category, Image: image},
		},
		// Otros casos de prueba pueden incluir combinaciones inválidas de categoría e imagen
	}
	// Test case: Verificar que la creación de un producto con categoría e imagen válidas sea correcta.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &factory.ProductFactory{}
			if got := f.CreateProduct(tt.args.category, tt.args.image); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductFactory.CreateProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductFactory_CreateCategory(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		f    *factory.ProductFactory
		args args
		want valueobjects.Category
	}{
		{
			name: "Test creation of new category with valid name",
			args: args{name: "Electronics"},
			want: valueobjects.Category{Name: "Electronics"},
		},
	}
	// Test case: Verificar que la creación de una categoría con un nombre válido sea correcta.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &factory.ProductFactory{}
			if got := f.CreateCategory(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductFactory.CreateCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductFactory_CreateImage(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		f    *factory.ProductFactory
		args args
		want valueobjects.Image
	}{
		{
			name: "Test creation of new image with valid URL",
			args: args{url: "image.jpg"},
			want: valueobjects.Image{URL: "image.jpg"},
		},
	}
	// Test case: Verificar que la creación de una imagen con una URL válida sea correcta.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &factory.ProductFactory{}
			if got := f.CreateImage(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductFactory.CreateImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
