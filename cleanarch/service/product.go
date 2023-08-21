// usecase/product_usecase.go
package service

import (
	"github.com/cocoasterr/cleanarch/domain"
	"github.com/cocoasterr/cleanarch/infra/database/postgres/repository"
)

type ProductUseCase struct {
    repo repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) *ProductUseCase {
    return &ProductUseCase{repo: repo}
}

func (uc *ProductUseCase) CreateProduct(product *domain.Product) error {
    // Call the repository's Create method
    if err := uc.repo.Create(*product); err != nil {
        return err
    }

    return nil
}

func (uc *ProductUseCase) GetAllProduct(product *domain.Product, page, limit int) (map[string]interface{}, error) {
    // Call the repository's Create method
    res, err := uc.repo.GetAll(*product, page, limit)
    if err != nil{
        return nil, err
    }

    return res, nil
}

func (uc *ProductUseCase) GetProductById(product *domain.Product, id int) (map[string]interface{}, error){
    
    res,err := uc.repo.GetById(*product, id)
    if err != nil{
        return nil, err
    }
    return res, nil
}
