package usecase

import (
	"context"
	"encoding/json"
	"log"
	"microservice-test/Service-1/common"
	"microservice-test/Service-1/domain"
	"microservice-test/Service-1/helper"
	"net/http"
)

type searchUsecase struct{}

func NewSearchUsecase() domain.SearchUsecase {
	return &searchUsecase{}
}

func (u *searchUsecase) Search(ctx context.Context, url string, req helper.SearchRequest) ([]helper.Movie, error) {
	var httpResp helper.HTTPResponse

	res, err := http.Get(url)
	if err != nil {
		log.Printf("Error when hit the URL, err: %v", err)
		return httpResp.Result, err
	}

	if res.StatusCode == http.StatusOK {
		decoder := json.NewDecoder(res.Body)
		err = decoder.Decode(&httpResp)
	} else {
		return httpResp.Result, common.ErrNotOk
	}

	return httpResp.Result, nil
}
