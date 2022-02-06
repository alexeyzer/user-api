package user_serivce

import (
	"context"
	desc "github.com/alexeyzer/user-api/pb/api/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *UserApiServiceServer) GetUserBySession(ctx context.Context, _ *emptypb.Empty) (*desc.CreateUserResponse, error) {
	sessionId := s.GetSessionIDFromContext(ctx)
	if sessionId == "" {
		return nil, status.Errorf(codes.Unauthenticated, "sessionID does`t exists")
	}
	res, err := s.userService.GetBySession(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	return s.serviceCreateUserResponseToProtoCreateUserResponse(res), nil
}
