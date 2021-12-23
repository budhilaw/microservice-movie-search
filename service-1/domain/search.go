package domain

import (
	"context"
	"microservice-test/Service-1/helper"
)

type SearchUsecase interface {
	Search(ctx context.Context, url string, req helper.SearchRequest) ([]helper.Movie, error)
}
