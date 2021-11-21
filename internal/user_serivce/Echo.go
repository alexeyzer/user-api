package user_serivce

import (
	"context"
	desc "github.com/alexeyzer/user-api/pb/api/user/v1"
)

func (s *UserApiServiceServer) Echo(ctx context.Context, request *desc.StringMessage) (*desc.StringMessage, error) {
	resp, err := s.UserService.Echo(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
