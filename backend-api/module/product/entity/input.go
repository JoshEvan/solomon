package entity

import "github.com/google/uuid"

type UpsertRequest struct {
	Id     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	ImgUrl string    `json:"img_url"`
}
