package entity

import "time"

type Product struct {
	Id         string    `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Price      float64   `json:"price"`
	ImgUrl     string    `json:"img_url" db:"img"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
	UpdateTime time.Time `json:"update_time" db:"update_time"`
}

type IndexedProduct struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type IndexedProductList []IndexedProduct

func (i *IndexedProductList) GetIds() (ids []string) {
	for _, e := range *i {
		ids = append(ids, (e.Id))
	}
	return
}
