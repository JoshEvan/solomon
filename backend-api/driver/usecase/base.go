package usecase

import "context"

type Usecase interface {
	Do(ctx context.Context) (interface{}, error)
	Validate(ctx context.Context) error
}
