package dto

import "time"

type NewProductRequest struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Desc  string `json:"desc"`
	Qty   int    `json:"qty"`
}

type ProductResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	Desc      string    `json:"desc"`
	Qty       int       `json:"qty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type ProductResponseJson struct {
	StatusCode int             `json:"statusCode"`
	Message    string          `json:"message"`
	Data       ProductResponse `json:"data"`
}

type ProductsResponseJson struct {
	StatusCode int               `json:"statusCode"`
	Message    string            `json:"message"`
	Data       []ProductResponse `json:"data"`
}

type GetProductRequest struct {
	ID int `json:"id"`
}
