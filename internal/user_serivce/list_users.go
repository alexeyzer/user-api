package user_serivce

import (
	"context"
	"github.com/alexeyzer/user-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/user-api/pb/api/user/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *UserApiServiceServer) ListUsers(ctx context.Context, _ *emptypb.Empty) (*desc.ListUsersResponse, error) {
	resp, err := s.userService.ListUsers(ctx)
	if err != nil {
		return nil, err
	}
	return s.usersToProtoListUsersResponse(resp), nil
}

func (s *UserApiServiceServer) usersToProtoListUsersResponse(users []*datastruct.User) *desc.ListUsersResponse {
	resp := &desc.ListUsersResponse{
		Users: make([]*desc.CreateUserResponse, 0, len(users)),
	}

	for _, user := range users {
		resp.Users = append(resp.Users, s.serviceCreateUserResponseToProtoCreateUserResponse(user))
	}
	return resp
}
