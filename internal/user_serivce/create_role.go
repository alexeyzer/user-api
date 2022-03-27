package user_serivce

import (
	"context"

	"github.com/alexeyzer/user-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/user-api/pb/api/user/v1"
)

func (s *UserApiServiceServer) CreateRole(ctx context.Context, req *desc.CreateRoleRequest) (*desc.CreateRoleResponse, error) {
	res, err := s.roleService.CreateRole(ctx, s.protoCreateRoleRequestToRole(req))
	if err != nil {
		return nil, err
	}
	return s.roleToProtoCreateRoleResponse(res), nil
}

func (s *UserApiServiceServer) protoCreateRoleRequestToRole(req *desc.CreateRoleRequest) datastruct.Role {
	return datastruct.Role{
		Name:        req.GetName(),
		Description: req.GetDescription(),
	}
}

func (s *UserApiServiceServer) roleToProtoCreateRoleResponse(resp *datastruct.Role) *desc.CreateRoleResponse {
	return &desc.CreateRoleResponse{
		Id:          resp.ID,
		Name:        resp.Name,
		Description: resp.Description,
	}
}
