package models

import (
	"database/sql"
)

type Client struct {
	ID      int
	Name    string
	Address string
	Age     int
}

func (c *Client) TableName() string {
	return "clients"
}

func (c *Client) Scan(rows *sql.Rows) (interface{}, error) {
	err := rows.Scan(&c.ID, &c.Name, &c.Address, &c.Age)
	return c, err
}
