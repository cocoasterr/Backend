package PGRepository

import "gorm.io/gorm"

type ProductRepository struct {
	Repository
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		Repository: Repository{
			DB: db,
		},
	}
}
