package storage

import (
	"context"

	"github.com/JoshEvan/solomon/driver/storage/elastic"
	"github.com/JoshEvan/solomon/driver/storage/entity"
)

type SearchEngine interface {
	Search(context.Context, entity.SearchQuery, entity.SearchPagination, entity.SearchSorting) (entity.SearchResult, error)
	Insert(context.Context, entity.SearchDocument) error
	Update(context.Context, entity.SearchDocument) error
}

func NewSearchEngine(cfg entity.Config) SearchEngine {
	return elastic.New(cfg)
}
