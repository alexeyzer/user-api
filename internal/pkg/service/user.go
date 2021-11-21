package service

import (
	"context"
	desc "github.com/alexeyzer/user-api/pb/api/user/v1"
)

type UserService interface {
	Echo(ctx context.Context, req *desc.StringMessage) (*desc.StringMessage, error)
}

type userService struct {
}

func (s *userService) Echo(ctx context.Context, request *desc.StringMessage) (*desc.StringMessage, error) {
	resp := &desc.StringMessage{
		Message: request.Message + " beach",
	}
	return resp, nil
}

func NewUserService() UserService {
	return &userService{}
}
