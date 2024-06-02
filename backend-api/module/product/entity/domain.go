package entity

type Product struct {
	Id     string `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	ImgUrl string `json:"img_url" db:"img_url"`
}
