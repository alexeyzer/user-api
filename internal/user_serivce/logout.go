package user_serivce

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *UserApiServiceServer) Logout(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	sessionId := s.GetSessionIDFromContext(ctx)
	if sessionId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "already logged out")
	}
	err := s.userService.DeleteSession(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
