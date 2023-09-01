package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cocoasterr/net_http/app/services"
	"github.com/cocoasterr/net_http/models"
)
type ProductController struct{
	ProductService services.ProductService
}

func NewProductController(productService services.ProductService) *ProductController{
	return &ProductController{
		ProductService: productService,
	}
}

func (c *ProductController) CreateProductController(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		http.Error(w, "Method not allowed!", http.StatusInternalServerError)
		return
	}
	var payload models.Product
	if err := json.NewDecoder(r.Body).Decode(&payload);err !=nil{
		http.Error(w, "Invalid Request Body!", http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	if err := c.ProductService.CreateProduct(ctx, &payload); err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp:= map[string]interface{}{"Message": "Success!"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (c *ProductController)IndexProdcuctController(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		http.Error(w, "Method not allowed!", http.StatusInternalServerError)
		return
	}
	page, _ := strconv.Atoi(r.URL.Query()["page"][0])
	limit, _ := strconv.Atoi(r.URL.Query()["limit"][0])
	
	var product models.Product
	ctx :=r.Context()
	res,err :=c.ProductService.IndexProduct(ctx, &product, page, limit)
	if err != nil{
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}