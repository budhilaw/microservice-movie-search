package helper

import (
	"context"
	"microservice-test/Service-1/search/delivery/pb"
)

type HTTPResponse struct {
	Result   []Movie `json:"Search"`
	Total    string  `json:"totalResults"`
	Response string  `json:"Response"`
}

type Movie struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type SearchRequest struct {
	MovieName string
	Page      int
}

type SearchResponse struct {
	Response     []Movie `json:"Search"`
	Total        string  `json:"totalResults"`
	ErrorMessage string  `json:"Error,omitempty"`
}

func DecodeGRPCSearchRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.SearchRequest)
	return SearchRequest{
		MovieName: req.MovieName,
		Page:      int(req.Page),
	}, nil
}

func EncodeGRPCSearchResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(SearchResponse)
	result := make([]*pb.Movie, 0)

	for _, movie := range resp.Response {
		result = append(result, &pb.Movie{
			Title:  movie.Title,
			Year:   movie.Year,
			Type:   movie.Type,
			ImdbID: movie.ImdbID,
			Poster: movie.Poster,
		})
	}

	return &pb.SearchResponse{
		MovieList: result,
		Err:       resp.ErrorMessage,
	}, nil
}
