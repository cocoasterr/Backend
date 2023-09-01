package PGRepositories

import (
	"context"
	"database/sql"
	"reflect"

	"github.com/cocoasterr/gonofw/models"
)

type ProductRepository interface{
	Create(ctx context.Context, product *models.Product)error
	Index(ctx context.Context)([]*models.Product, error)
	FindById(ctx context.Context, id string)(*models.Product, error)
	Update(ctx context.Context, product *models.Product, id string) error
	Delete(ctx context.Context, id string)error
}


type productRepository struct{
	Db *sql.DB
}

func NewProductRepository(db *sql.DB) *productRepository{
	return &productRepository{
		Db: db,
	}
}

func (r *productRepository)Create(ctx context.Context, product *models.Product)error{
	listKey := reflect.TypeOf(*product)
	query,values := GeneralCreate(ctx, listKey, product.TbName(), product)

	trx,err:= r.Db.Begin()
	if err != nil{
		return err
	}
	_, err = trx.ExecContext(ctx, query, values...)
	if err != nil{
		trx.Rollback()
		return err
	}
	err = trx.Commit()
	if err !=nil{
		return err
	}
	return nil
}

func (r *productRepository)Index(ctx context.Context)([]*models.Product, error){
	return nil, nil
}

func (r *productRepository)FindById(ctx context.Context, id string)(*models.Product, error){
	return nil, nil
}

func (r *productRepository)Update(ctx context.Context, product *models.Product, id string) error{
	return nil
}

func (r *productRepository)Delete(ctx context.Context, id string)error{
	return nil
}
