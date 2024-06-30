package elastic

import (
	"context"
	"log"

	"github.com/JoshEvan/solomon/driver/storage/entity"
	"github.com/olivere/elastic/v7"
)

const (
	defaultTimeoutMS = 5000
)

type ElasticClient struct {
	client          *elastic.Client
	timeoutInMillis int
}

type ElasticResult struct {
	*elastic.SearchResult
}

func New(cfg entity.Config) *ElasticClient {
	client, err := elastic.NewClient(
		elastic.SetURL(cfg.Address),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}

	return &ElasticClient{
		client:          client,
		timeoutInMillis: defaultTimeoutMS,
	}
}

func (e *ElasticClient) Search(ctx context.Context, query entity.SearchQuery, pagination entity.SearchPagination, sorting entity.SearchSorting) (result entity.SearchResult, err error) {
	searchResult, err := e.client.Search().
		Index(query.SearchSpace).
		Query(toESQuery(query.Params)).
		// Sort(sorting.SortBy, sorting.IsSortAsc).
		From(pagination.StartFrom()).
		Size(pagination.Limit).
		TimeoutInMillis(e.timeoutInMillis).
		Do(ctx)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &ElasticResult{searchResult}, nil
}

func (e *ElasticClient) Insert(ctx context.Context, doc entity.SearchDocument) error {
	_, err := e.client.Index().Index(doc.SearchSpace).Id(doc.Id).BodyJson(doc.Data).Do(ctx)
	return err
}

func (e *ElasticClient) Update(ctx context.Context, doc entity.SearchDocument) (err error) {
	_, err = e.client.Update().Index(doc.SearchSpace).Id(doc.Id).Doc(doc.Data).DocAsUpsert(true).DetectNoop(true).Do(ctx)
	return err
}

func toESQuery(params map[string]entity.SearchQueryCriteria) (q *elastic.BoolQuery) {
	q = elastic.NewBoolQuery()
	for k, v := range params {
		switch v.CriteriaType {
		case entity.SearchRange:
			if val, ok := v.Val.(entity.RangeFloatType); ok {
				q = q.Filter(elastic.NewRangeQuery(k).Gte(val.From).Lte(val.To))
			}
		case entity.SearchContain:
			if val, ok := v.Val.(string); ok {
				q = q.Should(elastic.NewMatchBoolPrefixQuery(k, val))
				q = q.Should(elastic.NewMatchQuery(k, val))
				q = q.Should(elastic.NewMatchPhraseQuery(k, val))
			}
		}
	}

	return
}

func (r *ElasticResult) GetDataList() (dataList []interface{}) {
	for _, e := range r.Hits.Hits {
		dataList = append(dataList, e.Source)
	}
	return
}
