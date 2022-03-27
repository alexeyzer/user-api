package user_serivce

import (
	"context"

	"github.com/alexeyzer/user-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/user-api/pb/api/user/v1"
)

func (s *UserApiServiceServer) GetRole(ctx context.Context, req *desc.GetRoleRequest) (*desc.GetRoleResponse, error) {
	res, err := s.roleService.GetRole(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return s.roleToProtoGetRoleResponse(res), nil
}

func (s *UserApiServiceServer) roleToProtoGetRoleResponse(resp *datastruct.Role) *desc.GetRoleResponse {
	return &desc.GetRoleResponse{
		Id:          resp.ID,
		Name:        resp.Name,
		Description: resp.Description,
	}
}
