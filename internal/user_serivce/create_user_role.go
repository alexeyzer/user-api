package user_serivce

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/alexeyzer/user-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/user-api/pb/api/user/v1"
)

func (s *UserApiServiceServer) CreateUserRole(ctx context.Context, req *desc.CreateUserRoleRequest) (*desc.CreateUserRoleResponse, error) {
	access, err := s.checkUserAdmin(ctx)
	if err != nil {
		return nil, err
	}
	if !access {
		return nil, status.Error(codes.PermissionDenied, "only for admins")
	}

	resp, err := s.userRoleService.CreateUserRole(ctx, s.protoCreateUserRoleRequestToUserRole(req))
	if err != nil {
		return nil, err
	}
	return s.userRoleToProtoCreateUserRoleResponse(resp), nil
}

func (s *UserApiServiceServer) protoCreateUserRoleRequestToUserRole(req *desc.CreateUserRoleRequest) datastruct.UserRole {
	return datastruct.UserRole{
		UserID: req.GetUserId(),
		RoleID: req.GetRoleId(),
	}
}

func (s *UserApiServiceServer) userRoleToProtoCreateUserRoleResponse(resp *datastruct.UserRole) *desc.CreateUserRoleResponse {
	return &desc.CreateUserRoleResponse{
		Id:     resp.ID,
		UserId: resp.UserID,
		RoleId: resp.RoleID,
	}
}
