package main

import (
	"log"
	"net/http"

	"github.com/funukonta/product-restapi/database"
	"github.com/funukonta/product-restapi/handler"
	"github.com/funukonta/product-restapi/repository/productRepo/product_pg"
	"github.com/funukonta/product-restapi/service"
	"github.com/gorilla/mux"
)

func main() {
	db := database.NewPG()

	productRepo := product_pg.NewProductPG(db)

	productService := service.NewProductService(productRepo)

	productHandler := handler.NewProductHandler(productService)

	r := mux.NewRouter()

	// Soal 1
	r.HandleFunc("/product", productHandler.CreateProduct).Methods("POST")
	// Soal 2
	r.HandleFunc("/product/sort/{sortby}/{type}", productHandler.GetProductSort).Methods("GET")

	// Untuk cek data aja
	r.HandleFunc("/product", productHandler.GetProduct).Methods("GET")
	r.HandleFunc("/product/{id}", productHandler.GetProductById).Methods("GET")

	port := ":8080"
	log.Println("Server listening to port", port)
	http.ListenAndServe(port, r)
}
