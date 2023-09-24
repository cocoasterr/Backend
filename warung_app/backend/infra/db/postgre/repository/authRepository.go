package PGRepository

import "gorm.io/gorm"

type AuthRepository struct {
	Repository
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		Repository: Repository{
			DB: db,
		},
	}
}
