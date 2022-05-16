package user_serivce

import (
	"context"
	"github.com/alexeyzer/user-api/internal/pkg/datastruct"

	desc "github.com/alexeyzer/user-api/pb/api/user/v1"
)

func (s *UserApiServiceServer) UpdateRole(ctx context.Context, req *desc.UpdateRoleRequest) (*desc.UpdateRoleResponse, error) {
	resp, err := s.roleService.UpdateRole(ctx, datastruct.Role{
		ID:          req.GetId(),
		Name:        req.GetName(),
		Description: req.GetDescription(),
	})
	if err != nil {
		return nil, err
	}
	return &desc.UpdateRoleResponse{
		Id:          resp.ID,
		Name:        resp.Name,
		Description: resp.Description,
	}, nil
}
