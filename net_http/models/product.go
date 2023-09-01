package models

import "time"

type Product struct{
	Id string `json:"id"`
	Name string `json:"name"`
	Stock int `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy time.Time `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy time.Time `json:"updated_by"`
}

func (p *Product)TbName()string{
	return "product"
}