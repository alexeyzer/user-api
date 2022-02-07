package user_serivce

import (
	"context"
	desc "github.com/alexeyzer/user-api/pb/api/user/v1"
)

func (s *UserApiServiceServer) Login(ctx context.Context, req *desc.LoginRequest) (*desc.LoginResponse, error) {
	res, err := s.userService.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	return &desc.LoginResponse{SessionId: *res}, nil
}
