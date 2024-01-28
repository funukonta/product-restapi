package product_pg

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/funukonta/product-restapi/entity"
	"github.com/funukonta/product-restapi/repository/productRepo"
)

type productPG struct {
	db *sql.DB
}

func NewProductPG(db *sql.DB) productRepo.ProductRepo {
	return &productPG{
		db: db,
	}
}

func (p *productPG) CreateProduct(newProduct entity.Product) (*entity.Product, error) {
	query := `INSERT INTO products (
		name,
		price,
		description,
		quantity
	) values ($1,$2,$3,$4)
	returning id,name,price,description,quantity,createdAt
	`
	product := entity.Product{}

	err := p.db.QueryRow(query, newProduct.Name, newProduct.Price, newProduct.Desc, newProduct.Qty).
		Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Desc,
			&product.Qty,
			&product.CreatedAt,
		)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *productPG) GetProductById(id string) (*entity.Product, error) {
	query := `select id,name,price,description,quantity,createdat,updatedat from products where id=$1`

	product := entity.Product{}
	err := p.db.QueryRow(query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.Desc,
		&product.Qty,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		if err != nil {
			if strings.Contains(err.Error(), "no rows") {
				return nil, fmt.Errorf("no data found")
			} else {
				return nil, err
			}
		}
	}

	return &product, nil
}

func (p *productPG) GetProduct() (*[]entity.Product, error) {
	query := `select id,name,price,description,quantity,createdat,updatedat from products`
	products := []entity.Product{}
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		prod := entity.Product{}
		err := rows.Scan(
			&prod.ID,
			&prod.Name,
			&prod.Price,
			&prod.Desc,
			&prod.Qty,
			&prod.CreatedAt,
			&prod.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, prod)
	}

	return &products, nil
}

func (p *productPG) GetProductSort(sortby, tipe string) (*[]entity.Product, error) {
	query := `select id,name,price,description,quantity,createdat,updatedat from products 
	order by %s %s`
	query = fmt.Sprintf(query, sortby, tipe)
	products := []entity.Product{}
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		prod := entity.Product{}
		err := rows.Scan(
			&prod.ID,
			&prod.Name,
			&prod.Price,
			&prod.Desc,
			&prod.Qty,
			&prod.CreatedAt,
			&prod.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, prod)
	}

	return &products, nil
}
