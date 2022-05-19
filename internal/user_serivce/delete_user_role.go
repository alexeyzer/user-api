package user_serivce

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/alexeyzer/user-api/pb/api/user/v1"
)

func (s *UserApiServiceServer) DeleteUserRole(ctx context.Context, req *desc.DeleteUserRoleRequest) (*emptypb.Empty, error) {
	access, err := s.checkUserAdmin(ctx)
	if err != nil {
		return nil, err
	}
	if !access {
		return nil, status.Error(codes.PermissionDenied, "only for admins")
	}

	err = s.userRoleService.DeleteUserRole(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
