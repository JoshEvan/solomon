package search

import (
	"context"
	"encoding/json"
	"log"

	"github.com/JoshEvan/solomon/driver/config"
	"github.com/JoshEvan/solomon/driver/storage"
	entitySearch "github.com/JoshEvan/solomon/driver/storage/entity"
	"github.com/JoshEvan/solomon/module/product/entity"
)

type SearchEngine interface {
	Search(ctx context.Context, searchQuery entity.SearchProductRequest, page, limit int, sortBy string, isSortAsc bool) (entity.IndexedProductList, error)
	Insert(context.Context, entity.IndexedProduct) error
	Update(context.Context, entity.IndexedProduct) error
}

type searchEngineImpl struct {
	se storage.SearchEngine
}

func GetSearchEngine(se storage.SearchEngine) *searchEngineImpl {
	return &searchEngineImpl{
		se: se,
	}
}

func (s *searchEngineImpl) Search(ctx context.Context, searchQuery entity.SearchProductRequest, page int, limit int, sortBy string, isSortAsc bool) (entity.IndexedProductList, error) {
	params := make(map[string]entitySearch.SearchQueryCriteria)
	if searchQuery.SearchText != "" {
		params["name"] = entitySearch.SearchQueryCriteria{
			Val:          searchQuery.SearchText,
			CriteriaType: entitySearch.SearchContain,
		}
	}
	if searchQuery.PriceMax != 0 && searchQuery.PriceMin != 0 {
		params["price"] = entitySearch.SearchQueryCriteria{
			Val: entitySearch.RangeFloatType{
				From: searchQuery.PriceMin,
				To:   searchQuery.PriceMax,
			},
			CriteriaType: entitySearch.SearchRange,
		}
	}

	jsonS, _ := json.Marshal(params)
	log.Println(string(jsonS))
	res, err := s.se.Search(ctx, entitySearch.SearchQuery{
		SearchSpace: config.IndexES,
		Params:      params,
	}, entitySearch.SearchPagination{
		Page:  page,
		Limit: limit,
	}, entitySearch.SearchSorting{
		SortBy:    sortBy,
		IsSortAsc: isSortAsc,
	})

	if err != nil {
		return nil, err
	}

	ret := []entity.IndexedProduct{}
	for _, e := range res.GetDataList() {
		indexed := entity.IndexedProduct{}
		if rawJSON, ok := e.(json.RawMessage); ok {
			json.Unmarshal(rawJSON, &indexed)
		}
		if len(indexed.Id) > 0 {
			ret = append(ret, indexed)
		}
	}

	return ret, nil
}

func (s *searchEngineImpl) Insert(ctx context.Context, req entity.IndexedProduct) error {
	return s.se.Insert(ctx, entitySearch.SearchDocument{
		SearchSpace: config.IndexES,
		Id:          req.Id,
		Data:        req,
	})
}

func (s *searchEngineImpl) Update(ctx context.Context, req entity.IndexedProduct) error {
	return s.se.Update(ctx, entitySearch.SearchDocument{
		SearchSpace: config.IndexES,
		Id:          req.Id,
		Data:        req,
	})
}
