package user_serivce

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/alexeyzer/user-api/config"
	"github.com/alexeyzer/user-api/internal/pkg/service"
	desc "github.com/alexeyzer/user-api/pb/api/user/v1"
)

type UserApiServiceServer struct {
	userService     service.UserService
	roleService     service.RoleService
	userRoleService service.UserRoleService
	cartService     service.CartService
	orderService    service.OrderService
	favoriteService service.FavoriteService
	desc.UnimplementedUserApiServiceServer
}

const adminRole = "admin"

func (s *UserApiServiceServer) checkUserAdmin(ctx context.Context) (bool, error) {
	res, err := s.SessionCheck(ctx, &emptypb.Empty{})
	if err != nil {
		return false, err
	}
	for _, role := range res.UserRoles {
		if role == adminRole {
			return true, nil
		}
	}
	return false, nil
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

func NewUserApiServiceServer(
	userService service.UserService,
	roleService service.RoleService,
	userRoleService service.UserRoleService,
	cartService service.CartService,
	orderService service.OrderService,
	favoriteService service.FavoriteService,
) *UserApiServiceServer {
	return &UserApiServiceServer{
		userService:     userService,
		roleService:     roleService,
		userRoleService: userRoleService,
		orderService:    orderService,
		cartService:     cartService,
		favoriteService: favoriteService,
	}
}
