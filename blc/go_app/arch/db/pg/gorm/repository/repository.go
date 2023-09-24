package PGRepository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type RepositoryInterface interface {
	Create(data interface{}) error
	Update(data interface{}, id string) error
	Index(data interface{}, limit, offset int) interface{}
	FindBy(data interface{}, key string, value interface{}) (interface{}, error)
	Delete(data interface{}, id string) error
	FindCustomQuery(data interface{}, query, condition string, value ...interface{}) interface{}
}

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) Create(data interface{}) error {
	trx := r.DB.Begin()
	defer trx.Commit()
	if err := trx.Create(data).Error; err != nil {
		trx.Rollback()
		return err
	}
	return nil
}

func (r *Repository) Update(data interface{}, id string) error {
	getData := r.DB.First(data, "id=?", id)
	if getData.Error != nil {
		if getData.Error == gorm.ErrRecordNotFound {
			return errors.New("data not found")
		}
		return getData.Error
	}
	trx := r.DB.Begin()
	defer trx.Commit()
	if err := trx.Save(data).Error; err != nil {
		trx.Rollback()
		return err
	}
	return nil
}

func (r *Repository) Index(data interface{}, limit, offset int) interface{} {
	r.DB.Find(data).Limit(limit).Offset(offset)
	return data
}

func (r *Repository) FindBy(data interface{}, key string, value interface{}) (interface{}, error) {
	cond := fmt.Sprintf("%s = ?", key)
	if err := r.DB.First(data, cond, value).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("data not found")
		}
		return nil, err
	}

	return data, nil
}

func (r *Repository) Delete(data interface{}, id string) error {
	getData := r.DB.First(data, "id=?", id)
	if getData.Error != nil {
		if getData.Error == gorm.ErrRecordNotFound {
			return errors.New("data not found")
		}
		return getData.Error
	}
	if err := r.DB.Delete(data, "id=?", id).Error; err != nil {
		return err
	}
	return nil
}
func (r *Repository) FindCustomQuery(data interface{}, query, condition string, value ...interface{}) interface{} {
	res := r.DB.Raw(query+condition, value...).Scan(data)
	return res

}
