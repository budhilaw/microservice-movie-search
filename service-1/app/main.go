package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"microservice-test/Service-1/common"
	"microservice-test/Service-1/search/delivery/http"
	"microservice-test/Service-1/search/delivery/pb"
	"microservice-test/Service-1/search/usecase"
	"net"
)

func init() {
	common.LoadEnv()
}

func main() {
	ctx := context.Background()
	service := usecase.NewSearchUsecase()
	endpoints := http.Endpoints{SearchEndpoint: http.MakeSearchEndpoint(service)}
	addr := fmt.Sprintf(":%d", common.Env.AppPort)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("Error when running listener, err: %v", err)
		return
	}

	log.Println("[+] gRPC running on ", addr)
	handler := common.NewGRPCServer(ctx, endpoints)
	gRPCServer := grpc.NewServer()
	pb.RegisterMovieSearchServer(gRPCServer, handler)
	err = gRPCServer.Serve(listener)
	if err != nil {
		log.Printf("grpc application down. please check. err : %s", err.Error())
	}
}
