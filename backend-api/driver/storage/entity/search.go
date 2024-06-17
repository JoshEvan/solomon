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
	DataType    string
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
	GetData() interface{}
}

func (p *SearchPagination) StartFrom() int {
	return (p.Page - 1) * p.Limit
}
