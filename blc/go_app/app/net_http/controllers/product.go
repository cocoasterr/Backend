package controllers

import (
	"github.com/cocoasterr/Backend/go_app/service"
)

type ProductConService struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) *ProductConService {
	return &ProductConService{
		ProductService: productService,
	}
}
