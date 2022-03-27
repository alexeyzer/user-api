package user_serivce

import (
	"context"

	"github.com/alexeyzer/user-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/user-api/pb/api/user/v1"
)

func (s *UserApiServiceServer) GetUserRole(ctx context.Context, req *desc.GetUserRoleRequest) (*desc.GetUserRoleResponse, error) {
	resp, err := s.userRoleService.GetUserRole(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return s.userRoleToProtoGetUserRoleResponse(resp), nil
}

func (s *UserApiServiceServer) userRoleToProtoGetUserRoleResponse(resp *datastruct.UserRole) *desc.GetUserRoleResponse {
	return &desc.GetUserRoleResponse{
		Id:     resp.ID,
		UserId: resp.UserID,
		RoleId: resp.RoleID,
	}
}
