package entity

import "github.com/google/uuid"

type UpsertRequest struct {
	Id     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Price  float64   `json:"price"`
	ImgUrl string    `json:"img_url"`
}
