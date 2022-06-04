package user_serivce

import (
	"context"
	desc "github.com/alexeyzer/user-api/pb/api/user/v1"
	"time"
)

func (s *UserApiServiceServer) Login(ctx context.Context, req *desc.LoginRequest) (*desc.LoginResponse, error) {
	//roles, sessionID, res, err := s.userService.Login(ctx, req)
	//if err != nil {
	//	return nil, err
	//}
	time.Sleep(50 * time.Millisecond)
	//resp := &desc.LoginResponse{
	//	Id:                 res.ID,
	//	Name:               res.Name,
	//	Surname:            res.Surname,
	//	Patronymic:         res.Patronymic,
	//	Phone:              res.Phone,
	//	Email:              res.Email,
	//	Session:            sessionID,
	//	AccessToAdminPanel: len(roles) > 0,
	//	Roles:              make([]*desc.CreateRoleResponse, 0, len(roles)),
	//}
	//for _, role := range roles {
	//	resp.Roles = append(resp.Roles, &desc.CreateRoleResponse{
	//		Id:          role.RoleID,
	//		Name:        role.RoleName,
	//		Description: "",
	//	})
	//}
	resp := &desc.LoginResponse{}
	return resp, nil
}
