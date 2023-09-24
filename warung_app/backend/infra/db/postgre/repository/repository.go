package PGRepository

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type RepositoryInterface interface {
	Create(domain interface{}) error
	Update(domain interface{}, id string) error
	GetAll(domain interface{}) interface{}
	FindBy(domain interface{}, key string, value interface{}) interface{}
	Delete(domain interface{}, id string) error
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) Create(domain interface{}) error {
	trx := r.DB.Begin()
	defer trx.Commit()

	if err := trx.Create(domain).Error; err != nil {
		trx.Rollback()
		return err
	}

	return nil
}

func (r *Repository) Update(domain interface{}, id string) error {
	trx := r.DB.Begin()
	defer trx.Commit()

	if err := trx.Save(domain).Where("id = ?", id).Error; err != nil {
		trx.Rollback()
		return err
	}
	return nil
}

func (r *Repository) GetAll(domain interface{}) interface{} {
	result := r.DB.Find(&domain)
	return result
}

func (r *Repository) FindBy(domain interface{}, key string, value interface{}) interface{} {
	cond := fmt.Sprintf("%s = ?", key)
	result := r.DB.First(&domain, cond, value)
	return result
}

// tb_name, model, condition
func (r *Repository) Delete(domain interface{}, id string) error {
	if err := r.DB.Delete(&domain).Where("id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) CustomQuery(domain interface{}, query string, condition string, value ...interface{}) interface{} {
	res := r.DB.Raw(query+condition, value...).Scan(domain)
	if res.RowsAffected > 0 {
		return nil
	}
	return res
}
