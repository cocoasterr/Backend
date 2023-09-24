package service

import PGRepository "github.com/cocoasterr/Backend/go_app/arch/db/pg/gorm/repository"

type ProductService struct {
	ServiceRepo
}

func NewProductService(repo PGRepository.Repository) *ProductService {
	return &ProductService{
		ServiceRepo: ServiceRepo{
			Repo: repo,
		},
	}
}
