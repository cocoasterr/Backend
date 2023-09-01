package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/cocoasterr/gonofw/app/services"
	"github.com/cocoasterr/gonofw/models"
)

type ProductControllers struct{
	productService services.ProductService
}

func NewProductController(productService services.ProductService) *ProductControllers{
	return &ProductControllers{
		productService: productService,
	}
}


func (c *ProductControllers) CreateProduct(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
	
	var payload models.Product

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil{
		http.Error(w, "Invalid Request Body!", http.StatusBadRequest)
		return
	}
	ctx := r.Context()

	if err := c.productService.CreateProduct(ctx, &payload); err!= nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response:= map[string]interface{}{"message": "success!"}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}