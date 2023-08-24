package sqlPostgreRepository

import (
	"database/sql"
	"reflect"

	"github.com/cocoaster/go-application/domain"
)



type ProductRepository struct{
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository{
	return &ProductRepository{
		db : db,
	}
}

func (r *ProductRepository) DB() *sql.DB {
    return r.db
}


func (r *ProductRepository) Create(product domain.Product) error {
	listKey := reflect.TypeOf(product)

	return PGGeneralCreate(r, listKey, &product)
}

func (r *ProductRepository) GetAll(product domain.Product, page, limit int) (map [string]interface{}, error) {
	result, err := PGGeneralGetAll(r, page, limit, &product)
	if err != nil{
		return nil, err
	}
	return result, nil
}

func (r *ProductRepository) GetById(product domain.Product, id int) (map[string]interface{}, error) {
	result, err := PGGeneralGetById(r,&product, id )
	if err != nil{
		return nil, err
	}
	return result, nil
}
