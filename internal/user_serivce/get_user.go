package user_serivce

import (
	"context"

	desc "github.com/alexeyzer/user-api/pb/api/user/v1"
)

func (s *UserApiServiceServer) GetUser(ctx context.Context, req *desc.GetUserRequest) (*desc.CreateUserResponse, error) {
	resp, err := s.userService.GetUser(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return s.serviceCreateUserResponseToProtoCreateUserResponse(resp), nil
}
