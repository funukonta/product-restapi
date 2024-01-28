package service

import (
	"fmt"
	"net/http"

	"github.com/funukonta/product-restapi/dto"
	"github.com/funukonta/product-restapi/entity"
	productrepo "github.com/funukonta/product-restapi/repository/productRepo"
)

type ProductService interface {
	CreateProduct(newProductRequest dto.NewProductRequest) (*dto.ProductResponseJson, error)
	GetProductById(string) (*dto.ProductResponseJson, error)
	GetProduct() (*dto.ProductsResponseJson, error)
	GetProductSort(string, string) (*dto.ProductsResponseJson, error)
}

type productService struct {
	productRepo productrepo.ProductRepo
}

func NewProductService(newproductRepo productrepo.ProductRepo) ProductService {
	return &productService{
		productRepo: newproductRepo,
	}
}

func (p *productService) CreateProduct(newProductRequest dto.NewProductRequest) (*dto.ProductResponseJson, error) {
	productReq := entity.Product{
		Name:  newProductRequest.Name,
		Price: newProductRequest.Price,
		Desc:  newProductRequest.Desc,
		Qty:   newProductRequest.Qty,
	}

	newProduct, err := p.productRepo.CreateProduct(productReq)
	if err != nil {
		return nil, err
	}

	productRes := dto.ProductResponse{
		ID:        newProduct.ID,
		Name:      newProduct.Name,
		Price:     newProduct.Price,
		Desc:      newProduct.Desc,
		Qty:       newProduct.Qty,
		CreatedAt: newProduct.CreatedAt,
	}

	response := dto.ProductResponseJson{
		Message:    "success add product",
		StatusCode: http.StatusCreated,
		Data:       productRes,
	}

	return &response, nil
}

func (p *productService) GetProductById(id string) (*dto.ProductResponseJson, error) {

	productDb, err := p.productRepo.GetProductById(id)
	if err != nil {
		return nil, err
	}

	response := dto.ProductResponseJson{
		Message:    "success get product",
		StatusCode: http.StatusOK,
		Data: dto.ProductResponse{
			ID:        productDb.ID,
			Name:      productDb.Name,
			Price:     productDb.Price,
			Desc:      productDb.Desc,
			Qty:       productDb.Qty,
			CreatedAt: productDb.CreatedAt,
			UpdatedAt: productDb.UpdatedAt,
		},
	}

	return &response, nil
}

func (p *productService) GetProduct() (*dto.ProductsResponseJson, error) {

	productDb, err := p.productRepo.GetProduct()
	if err != nil {
		return nil, err
	}
	pdb := *productDb
	data := []dto.ProductResponse{}
	for i := range pdb {
		data = append(data, dto.ProductResponse(pdb[i]))
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("data not found")
	}

	response := dto.ProductsResponseJson{
		Message:    "success get products",
		StatusCode: http.StatusOK,
		Data:       data,
	}

	return &response, nil
}

func (p *productService) GetProductSort(sortby, tipe string) (*dto.ProductsResponseJson, error) {

	if sortby != "createdat" && sortby != "price" && sortby != "name" {
		return nil, fmt.Errorf("sorting type not allowed")
	}
	if tipe != "asc" && tipe != "desc" {
		return nil, fmt.Errorf("sorting type not allowed")
	}

	productDb, err := p.productRepo.GetProductSort(sortby, tipe)
	if err != nil {
		return nil, err
	}
	pdb := *productDb
	data := []dto.ProductResponse{}
	for i := range pdb {
		data = append(data, dto.ProductResponse(pdb[i]))
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("data not found")
	}

	response := dto.ProductsResponseJson{
		Message:    "success get products",
		StatusCode: http.StatusOK,
		Data:       data,
	}

	return &response, nil
}
