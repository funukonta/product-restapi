package handler

import (
	"net/http"

	"github.com/funukonta/product-restapi/dto"
	"github.com/funukonta/product-restapi/pkg"
	"github.com/funukonta/product-restapi/service"
	"github.com/gorilla/mux"
)

type productHandler struct {
	productService service.ProductService
}

func NewProductHandler(service service.ProductService) *productHandler {
	return &productHandler{
		productService: service,
	}
}

func (ph *productHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newproductReq dto.NewProductRequest

	err := pkg.DecodeJsonReq(r, &newproductReq)
	if err != nil {
		pkg.WriteJsonRes(w, 400, err.Error())
		return
	}

	res, err := ph.productService.CreateProduct(newproductReq)
	if err != nil {
		pkg.WriteJsonRes(w, 400, err.Error())
		return
	}

	pkg.WriteJsonRes(w, res.StatusCode, res)
}

func (ph *productHandler) GetProductById(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	res, err := ph.productService.GetProductById(id)
	if err != nil {
		pkg.WriteJsonRes(w, 400, pkg.MessageErr{
			Data:    res,
			Message: err.Error(),
		})
		return
	}

	pkg.WriteJsonRes(w, res.StatusCode, res)
}

func (ph *productHandler) GetProduct(w http.ResponseWriter, r *http.Request) {

	res, err := ph.productService.GetProduct()
	if err != nil {
		pkg.WriteJsonRes(w, 400, pkg.MessageErr{
			Data:    res,
			Message: err.Error(),
		})
		return
	}

	pkg.WriteJsonRes(w, res.StatusCode, res)
}

func (ph *productHandler) GetProductSort(w http.ResponseWriter, r *http.Request) {

	sortby := mux.Vars(r)["sortby"]
	tipe := mux.Vars(r)["type"]
	res, err := ph.productService.GetProductSort(sortby, tipe)
	if err != nil {
		pkg.WriteJsonRes(w, 400, pkg.MessageErr{
			Data:    res,
			Message: err.Error(),
		})
		return
	}

	pkg.WriteJsonRes(w, res.StatusCode, res)
}
