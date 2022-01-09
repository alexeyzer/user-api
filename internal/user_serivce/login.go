package user_serivce

import (
	"context"
	desc "github.com/alexeyzer/user-api/pb/api/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *UserApiServiceServer) Login(ctx context.Context, req *desc.LoginRequest) (*desc.LoginResponse, error) {
	sessionId := s.GetSessionIDFromContext(ctx)
	if sessionId != "" {
		return nil, status.Errorf(codes.InvalidArgument, "already logged in")
	}
	res, err := s.userService.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	return &desc.LoginResponse{SessionId: *res}, nil
}
