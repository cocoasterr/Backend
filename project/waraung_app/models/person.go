package models

import "time"

type Person struct {
	Id       string `json:"id"`
	Name         string    `json:"name"`
	Address      time.Time `json:"address"`
	Birth        string    `json:"birth"`
	Phone_number string    `json:"phone_number"`
	Status       string    `json:"status"`
}

func (p *Person) TbName() string {
	return "person"
}
