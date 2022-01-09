package user_serivce

import (
	"context"
	desc "github.com/alexeyzer/user-api/pb/api/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *UserApiServiceServer) SessionCheck(ctx context.Context, req *emptypb.Empty) (*desc.SessionCheckResponse, error) {
	sessionId := s.GetSessionIDFromContext(ctx)
	if sessionId == "" {
		return nil, status.Errorf(codes.Unauthenticated, "sessionID does`t exists")
	}
	res, err := s.userService.SessionCheck(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	return &desc.SessionCheckResponse{Username: *res}, nil
}
