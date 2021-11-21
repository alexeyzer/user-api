package main

import (
	"context"
	"flag"
	"github.com/alexeyzer/user-api/internal/pkg/service"
	"github.com/alexeyzer/user-api/internal/user_serivce"
	gw "github.com/alexeyzer/user-api/pb/api/user/v1"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
)

func RunGrpc() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		return err
	}

	userService := service.NewUserService()
	userApiServiceServer := user_serivce.NewUserApiServiceServer(userService)

	grpcServer := grpc.NewServer()
	gw.RegisterUserApiServiceServer(grpcServer, userApiServiceServer)
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(),
	}
	go func() {
		gw.RegisterUserApiServiceHandlerFromEndpoint(ctx, mux, ":8082", opts)
	}()
	log.Println("app started")
	grpcServer.Serve(lis)
	return nil
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := RunGrpc(); err != nil {
		glog.Fatal(err)
	}
}
