package user_serivce

import (
	"github.com/alexeyzer/user-api/internal/pkg/service"
	desc "github.com/alexeyzer/user-api/pb/api/user/v1"
)

type UserApiServiceServer struct {
	userService service.UserService
	desc.UnimplementedUserApiServiceServer
}

func NewUserApiServiceServer(userService service.UserService) *UserApiServiceServer {
	return &UserApiServiceServer{
		userService: userService,
	}
}
