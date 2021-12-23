package main

import (
	"context"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"microservice-test/Service-1/common"
	"microservice-test/Service-1/search/delivery/http"
	"microservice-test/Service-1/search/delivery/pb"
	"microservice-test/Service-1/search/usecase"
	"net"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	ctx := context.Background()
	service := usecase.NewSearchUsecase()
	endpoints := http.Endpoints{SearchEndpoint: http.MakeSearchEndpoint(service)}
	listener, err := net.Listen("tcp", viper.GetString("server.address"))
	if err != nil {
		log.Printf("Error when running listener, err: %v", err)
		return
	}

	log.Println("[+] gRPC running on ", viper.GetString("server.address"))
	handler := common.NewGRPCServer(ctx, endpoints)
	gRPCServer := grpc.NewServer()
	pb.RegisterMovieSearchServer(gRPCServer, handler)
	err = gRPCServer.Serve(listener)
	if err != nil {
		log.Printf("grpc application down. please check. err : %s", err.Error())
	}
}
