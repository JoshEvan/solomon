package entity

import "github.com/google/uuid"

type Product struct {
	Id     uuid.UUID `json:"id" db:"id"`
	Name   string    `json:"name" db:"name"`
	Price  float64   `json:"price"`
	ImgUrl string    `json:"img_url" db:"img"`
}
