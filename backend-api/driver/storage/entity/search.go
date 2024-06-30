package entity

type Config struct {
	Address string
}

type SearchQuery struct {
	SearchSpace string
	Params      map[string]SearchQueryCriteria
}

type SearchDocument struct {
	SearchSpace string
	Id          string
	Data        interface{}
}

type SearchQueryCriteria struct {
	Val          interface{}
	CriteriaType SearchQueryCriteriaType
}

type SearchQueryCriteriaType int

const (
	SearchContain SearchQueryCriteriaType = iota
	SearchRange
	SearchIncludedInArray
)

type RangeFloatType struct {
	From float64
	To   float64
}

type SearchPagination struct {
	Page  int
	Limit int
}

type SearchSorting struct {
	SortBy    string
	IsSortAsc bool
}

type SearchResult interface {
	GetDataList() []interface{}
}

func (p *SearchPagination) StartFrom() int {
	from := (p.Page - 1) * p.Limit
	if from < 0 {
		from = 0
	}
	return from
}
