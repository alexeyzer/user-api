package main

import (
	"context"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net"
	"net/http"
	"time"

	"github.com/alexeyzer/user-api/config"
	"github.com/alexeyzer/user-api/internal/client"
	"github.com/alexeyzer/user-api/internal/pkg/repository"
	"github.com/alexeyzer/user-api/internal/pkg/service"
	"github.com/alexeyzer/user-api/internal/user_serivce"
	gw "github.com/alexeyzer/user-api/pb/api/user/v1"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

func serveSwagger(mux *http.ServeMux) {
	prefix := "/swagger/"
	sh := http.StripPrefix(prefix, http.FileServer(http.Dir("./swagger/")))
	mux.Handle(prefix, sh)
}

// look up session and pass sessionId in to context if it exists
func gatewayMetadataAnnotator(_ context.Context, r *http.Request) metadata.MD {
	//SessionID, ok := r.Cookie(config.Config.Auth.SessionKey)
	SessionID := r.Header.Get(config.Config.Auth.SessionKey)
	if SessionID != "" {
		return metadata.Pairs(config.Config.Auth.SessionKey, SessionID)
	}
	return metadata.Pairs()
}

func httpResponseModifier(ctx context.Context, w http.ResponseWriter, _ proto.Message) error {
	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		return nil
	}

	sessionID := md.HeaderMD.Get(config.Config.Auth.SessionKey)
	logout := md.HeaderMD.Get(config.Config.Auth.LogoutKey)
	if len(sessionID) > 0 {
		if len(logout) == 0 {
			//w.Header().Set(config.Config.Auth.SessionKey, sessionID[0])
			//http.SetCookie(w, &http.Cookie{
			//	Name:     config.Config.Auth.SessionKey,
			//	Value:    sessionID[0],
			//	Path:     "/",
			//	HttpOnly: true,
			//	Expires:  time.Now().Add(time.Hour * 24),
			//	SameSite: http.SameSiteNoneMode,
			//	Secure:   true,
			//})
		} else {
			//http.SetCookie(w, &http.Cookie{
			//	Name:     config.Config.Auth.SessionKey,
			//	Value:    sessionID[0],
			//	Path:     "/",
			//	HttpOnly: true,
			//	Expires:  time.Now().Add(time.Duration(-1) * time.Hour * 24),
			//	SameSite: http.SameSiteNoneMode,
			//	Secure:   true,
			//})
		}
	}
	return nil
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("origin"))
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "sessionid, Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		if (*r).Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}

func RunServer(ctx context.Context, userApiServiceServer *user_serivce.UserApiServiceServer) error {

	grpcLis, err := net.Listen("tcp", ":"+config.Config.App.GrpcPort)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_logrus.UnaryServerInterceptor(log.WithContext(ctx).WithTime(time.Time{})),
		grpc_validator.UnaryServerInterceptor(),
		grpc_prometheus.UnaryServerInterceptor)),
	)
	gw.RegisterUserApiServiceServer(grpcServer, userApiServiceServer)

	mux := http.NewServeMux()
	gwmux := runtime.NewServeMux(runtime.WithMetadata(gatewayMetadataAnnotator), runtime.WithForwardResponseOption(httpResponseModifier))
	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("/", cors(gwmux))
	serveSwagger(mux)
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(),
	}
	err = gw.RegisterUserApiServiceHandlerFromEndpoint(ctx, gwmux, ":"+config.Config.App.GrpcPort, opts)
	if err != nil {
		return err
	}

	grpc_prometheus.Register(grpcServer)
	grpc_prometheus.EnableHandlingTimeHistogram()

	go func() {
		err = grpcServer.Serve(grpcLis)
		log.Fatal(err)
	}()
	log.Println("app started")
	err = http.ListenAndServe(":"+config.Config.App.HttpPort, mux)
	log.Fatal(err)
	return err
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	err := config.ReadConf("./config/config.yaml")
	if err != nil {
		log.Fatal("Failed to create config: ", err)
	}

	dao, err := repository.NewDao()
	if err != nil {
		log.Fatal("Failed to connect to db: ", err)
	}

	redis, err := client.NewRedisClient(ctx)
	if err != nil {
		log.Fatal("Failed to connect to redis db: ", err)
	}

	productAPIClient, err := client.NewProductApiClient(config.Config.GRPC.ProductAPI)

	userService := service.NewUserService(dao, redis)
	roleService := service.NewRoleService(dao)
	userRoleService := service.NewUserRoleService(dao)
	cartService := service.NewCartService(dao, productAPIClient)
	orderService := service.NewOrderService(dao, cartService, productAPIClient)
	favoriteService := service.NewFavoriteService(dao, productAPIClient)

	userApiServiceServer := user_serivce.NewUserApiServiceServer(userService, roleService, userRoleService, cartService, orderService, favoriteService)
	if err := RunServer(ctx, userApiServiceServer); err != nil {
		log.Fatal(err)
	}
}
