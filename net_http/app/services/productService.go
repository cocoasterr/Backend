package services

import (
	"context"

	PGRepository "github.com/cocoasterr/net_http/infra/sql/postgres/repository"
	"github.com/cocoasterr/net_http/models"
)

type ProductInterface interface{
	CreateProduct(ctx context.Context, product *models.Product) error
	IndexProduct(ctx context.Context,page, limit int)(map[string]interface{}, error)
	GetProductById(ctx context.Context, id string)(*models.Product, error)
	UpdateProduct(ctx context.Context,product *models.Product, id string) error
	DeleteProduct(ctx context.Context, id string) error
}

type ProductService struct{
	ProductRepo PGRepository.ProductRepository
}

func NewProductService(productRepo PGRepository.ProductRepository) *ProductService{
	return &ProductService{
		ProductRepo: productRepo,
	}
}

func(s *ProductService)CreateProduct(ctx context.Context, product *models.Product) error{
	return s.ProductRepo.Create(ctx, product)
}

func(s *ProductService)IndexProduct(ctx context.Context,product *models.Product, page, limit int)(map[string]interface{}, error){
	return s.ProductRepo.Index(ctx, product, page, limit)
}
func(s *ProductService)GetProductById(ctx context.Context, id string)(*models.Product, error){
	return s.ProductRepo.GetById(ctx, id)
}
func(s *ProductService)UpdateProduct(ctx context.Context,product *models.Product, id string) error{
	return s.ProductRepo.Update(ctx, product, id)
}
func(s *ProductService)DeleteProduct(ctx context.Context, id string) error{
	return s.ProductRepo.Delete(ctx, id)
}