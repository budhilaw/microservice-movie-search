package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"log"
	"microservice-test/Service-1/constant"
	"microservice-test/Service-1/domain"
	"microservice-test/Service-1/helper"
	"microservice-test/Service-1/search/delivery/pb"
)

type Endpoints struct {
	SearchEndpoint endpoint.Endpoint
}

func MakeSearchEndpoint(srv domain.SearchUsecase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(helper.SearchRequest)
		url := fmt.Sprintf(constant.URL_FORMAT, constant.OMDB_KEY, req.MovieName, req.Page)
		res, err := srv.Search(ctx, url, req)
		if err != nil {
			return helper.SearchResponse{ErrorMessage: err.Error()}, errors.New(err.Error())
		}
		return helper.SearchResponse{Response: res}, nil
	}
}

func (e Endpoints) Search(ctx context.Context) ([]helper.Movie, error) {
	req := pb.SearchRequest{}
	resp, err := e.SearchEndpoint(ctx, req)
	if err != nil {
		log.Printf("Error when running search endpoint, err: %v", err)
		return nil, err
	}
	getResp := resp.(helper.SearchResponse)
	if getResp.ErrorMessage != "" {
		return nil, errors.New(getResp.ErrorMessage)
	}
	return getResp.Response, nil
}
