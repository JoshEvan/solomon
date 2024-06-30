package entity

type UpsertRequest struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	ImgUrl string  `json:"img_url"`
}

type SearchProductRequest struct {
	SearchText string  `json:"search_text"`
	PriceMin   float64 `json:"price_min"`
	PriceMax   float64 `json:"price_max"`
}

const (
	PageFormValue       = "page"
	LimitFormValue      = "limit"
	SearchTextFormValue = "q"
	SortByFormValue     = "sort"
	IsSortAscFormValue  = "asc"
	PriceMinFormValue   = "pmin"
	PriceMaxFormValue   = "pmax"
)

type SelectQuery struct {
	Page       int
	Limit      int
	SearchText string
	SortBy     string
	IsSortAsc  bool
	PriceMin   float64
	PriceMax   float64
}
