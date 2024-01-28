package productRepo

import "github.com/funukonta/product-restapi/entity"

type ProductRepo interface {
	CreateProduct(newProduct entity.Product) (*entity.Product, error)
	GetProductById(string) (*entity.Product, error)
	GetProduct() (*[]entity.Product, error)
	GetProductSort(string, string) (*[]entity.Product, error)
}
