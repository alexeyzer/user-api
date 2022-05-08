package user_serivce

import (
	"context"
	desc "github.com/alexeyzer/user-api/pb/api/user/v1"
)

func (s *UserApiServiceServer) Login(ctx context.Context, req *desc.LoginRequest) (*desc.LoginResponse, error) {
	sessionID, res, err := s.userService.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	return &desc.LoginResponse{
		Id:         res.ID,
		Name:       res.Name,
		Surname:    res.Surname,
		Patronymic: res.Patronymic,
		Phone:      res.Phone,
		Email:      res.Email,
		Session:    sessionID,
	}, nil
}
