	package models

import (
	"database/sql"
)

type Product struct {
	ID       int
	Name     string
	Quantity int
}

func (p *Product) TableName() string {
	return "products"
}

func (p *Product) Scan(rows *sql.Rows) (interface{}, error) {
	err := rows.Scan(&p.ID, &p.Name, &p.Quantity)
	return p, err
}
