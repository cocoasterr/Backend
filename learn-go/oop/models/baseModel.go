package models

import "time"

type setTime struct {
	UpdatedAt time.Time `json:"updated_at"`
}

func NewSetTime() *setTime {
	return &setTime{
		UpdatedAt: time.Now(),
	}
}
