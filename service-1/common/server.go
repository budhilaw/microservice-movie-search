package common

import (
	"context"
	gRPCTransport "github.com/go-kit/kit/transport/grpc"
	"microservice-test/Service-1/helper"
	"microservice-test/Service-1/search/delivery/http"
	"microservice-test/Service-1/search/delivery/pb"
)

type grpcServer struct {
	movie gRPCTransport.Handler
	pb.UnimplementedMovieSearchServer
}

func (s *grpcServer) SearchMovie(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	_, resp, err := s.movie.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.SearchResponse), nil
}

func NewGRPCServer(_ context.Context, endpoint http.Endpoints) pb.MovieSearchServer {
	return &grpcServer{
		movie: gRPCTransport.NewServer(
			endpoint.SearchEndpoint,
			helper.DecodeGRPCSearchRequest,
			helper.EncodeGRPCSearchResponse,
		),
	}
}
