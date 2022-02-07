package user_serivce

import (
	"context"
	"github.com/alexeyzer/user-api/config"
	"github.com/alexeyzer/user-api/internal/pkg/service"
	desc "github.com/alexeyzer/user-api/pb/api/user/v1"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

type UserApiServiceServer struct {
	userService service.UserService
	desc.UnimplementedUserApiServiceServer
}

func (s *UserApiServiceServer) GetSessionIDFromContext(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		val := md.Get(config.Config.Auth.SessionKey)
		if len(val) > 0 {
			return val[0]
		}
		log.Info("no value with key:", config.Config.Auth.SessionKey)
	}
	log.Info("no metadata")
	return ""
}

func NewUserApiServiceServer(userService service.UserService) *UserApiServiceServer {
	return &UserApiServiceServer{
		userService: userService,
	}
}
