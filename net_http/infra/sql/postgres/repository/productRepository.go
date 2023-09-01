package PGRepository

import (
	"context"
	"database/sql"
	"reflect"

	"github.com/cocoasterr/net_http/models"
)

type ProductRepoInterface interface{
	Create(ctx context.Context, product *models.Product)error
	Index(ctx context.Context, product *models.Product, page, limit int) (map[string]interface{}, error)
	GetById(ctx context.Context, id string)([]*models.Product, error)
	Update(ctx context.Context, product *models.Product, id string)error
	Delete(ctx context.Context, id string)error
}


type ProductRepository struct{
	Db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository{
	return &ProductRepository{
		Db: db,
	}
}

func (p *ProductRepository) Create(ctx context.Context, product *models.Product)error{
	listKey := reflect.TypeOf(*product)
	query,values := GeneralCreate(ctx, listKey, product.TbName(), product)
	trx, err:= p.Db.Begin()
	if err != nil{
		return err
	}
	_, err = trx.ExecContext(ctx, query, values...)
	if err != nil{
		trx.Rollback()
		return err
	}
	err = trx.Commit()
	if err != nil {
		trx.Rollback()
		return err
	}
	return nil
}
func (p *ProductRepository) Index(ctx context.Context, product *models.Product, page, limit int) (map[string]interface{}, error){
	return GeneralIndex(p.Db, ctx, page, limit, product.TbName())
}
func (p *ProductRepository) GetById(ctx context.Context, id string)(*models.Product, error){
	return nil, nil
}
func (p *ProductRepository) Update(ctx context.Context, product *models.Product, id string)error{
	return nil
}
func (p *ProductRepository) Delete(ctx context.Context, id string)error{
	return nil
}