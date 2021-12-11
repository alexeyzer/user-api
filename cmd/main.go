package main

import (
	"context"
	"github.com/alexeyzer/user-api/internal/pkg/service"
	"github.com/alexeyzer/user-api/internal/user_serivce"
	gw "github.com/alexeyzer/user-api/pb/api/user/v1"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"time"
)

func serveSwagger(mux *http.ServeMux) {
	prefix := "/swagger/"
	sh := http.StripPrefix(prefix, http.FileServer(http.Dir("./swagger/")))
	mux.Handle(prefix, sh)
}

func RunServer(userApiServiceServer *user_serivce.UserApiServiceServer) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	grpcLis, err := net.Listen("tcp", ":8082")
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_logrus.UnaryServerInterceptor(log.WithContext(ctx).WithTime(time.Time{})),
		grpc_validator.UnaryServerInterceptor()),
	))
	gw.RegisterUserApiServiceServer(grpcServer, userApiServiceServer)

	mux := http.NewServeMux()
	gwmux := runtime.NewServeMux()
	mux.Handle("/", gwmux)
	serveSwagger(mux)
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(),
	}
	err = gw.RegisterUserApiServiceHandlerFromEndpoint(ctx, gwmux, ":8082", opts)
	if err != nil {
		return err
	}
	go func() {
		err = grpcServer.Serve(grpcLis)
		log.Fatal(err)
	}()
	log.Println("app started")
	err = http.ListenAndServe(":8080", mux)
	return err
}

func main() {

	userService := service.NewUserService()
	userApiServiceServer := user_serivce.NewUserApiServiceServer(userService)
	if err := RunServer(userApiServiceServer); err != nil {
		log.Fatal(err)
	}
}
