package domain

import (
	"time"
)


type Product struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Quantity int `json:"qty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
func (r *Product)NewProduct(p Product) Product {
	return Product{
		Id:        p.Id,
		Name:      p.Name,
		Quantity:  p.Quantity,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (r *Product)TableName()string{
	return "products"
}
