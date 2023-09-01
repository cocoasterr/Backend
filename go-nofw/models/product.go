package models

import "time"

type Product struct{
	Id string `json:"id"`
	Name string `json:"name"`
	Stock int `json:"stock"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
	UpdatedAt time.Time `json:"updatedAt"`
	UpdatedBy string `json:"updatedBy"`
}

func (p *Product)TbName()string{
	return "product"
}