package services

import (
	"context"

	PGRepositories "github.com/cocoasterr/gonofw/infra/postgre/repositories"
	"github.com/cocoasterr/gonofw/models"
)

type ProductService interface{
	CreateProduct(ctx context.Context, product *models.Product)error
	IndexProduct(ctx context.Context)([]*models.Product, error)
	FindProductById(ctx context.Context, id string)(*models.Product, error)
	UpdateProduct(ctx context.Context, product *models.Product, id string) error
	DeleteProduct(ctx context.Context, id string)error
}

type productService struct{
	productRepository PGRepositories.ProductRepository
}

func NewProductSevice(repo PGRepositories.ProductRepository) *productService{
	return &productService{
		productRepository: repo,
	}
}

func (r *productService)CreateProduct(ctx context.Context, product *models.Product)error{
	return r.productRepository.Create(ctx, product)
}

func (r *productService)IndexProduct(ctx context.Context)([]*models.Product, error){
	return nil, nil
}

func (r *productService)FindProductById(ctx context.Context, id string)(*models.Product, error){
	return nil, nil
}

func (r *productService)UpdateProduct(ctx context.Context, product *models.Product, id string) error{
	return nil
}

func (r *productService)DeleteProduct(ctx context.Context, id string)error{
	return nil
}

